package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/model/system"
	"github.com/liuxing95/gin-react-admin/service"
	"github.com/liuxing95/gin-react-admin/utils"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body []byte
		var userId int
		if ctx.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(ctx.Request.Body)
			if err != nil {
				global.GRA_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := ctx.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.GetClaims(ctx)
		if claims.BaseClaims.ID != 0 {
			userId = int(claims.BaseClaims.ID)
		} else {
			id, err := strconv.Atoi(ctx.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:     ctx.ClientIP(),
			Method: ctx.Request.Method,
			Path:   ctx.Request.URL.Path,
			Agent:  ctx.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(ctx.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}

		ctx.Writer = writer
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = ctx.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Contains(ctx.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(ctx.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(ctx.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(ctx.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Resp = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.GRA_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
