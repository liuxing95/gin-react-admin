package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(ctx.Writer, ctx.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		ctx.Next()
	}
}
