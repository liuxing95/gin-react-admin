package main

import (
	"github.com/liuxing95/gin-react-admin/global"
	"github.com/liuxing95/gin-react-admin/initialize"
)

func main() {
	global.GRA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GRA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GRA_DB.DB()
		defer db.Close()
	}
}
