package initialize

import (
	"os"

	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GRA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.GRA_DB
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
	)
	if err != nil {
		global.GRA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GRA_LOG.Info("register table success")
}
