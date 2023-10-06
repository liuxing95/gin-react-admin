package system

import "github.com/liuxing95/gin-react-admin/global"

type JwtBlacklist struct {
	global.GRA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
