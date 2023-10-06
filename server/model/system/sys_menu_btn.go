package system

import "github.com/liuxing95/gin-react-admin/global"

type SysBaseMenuBtn struct {
	global.GRA_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
