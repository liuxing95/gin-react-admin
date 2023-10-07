package request

import "github.com/liuxing95/gin-react-admin/model/common/request"

type SysAutoHistory struct {
	request.PageInfo
}

// GetById Find by id structure
type RollBack struct {
	ID          int  `json:"id" form:"id"`                   // 主键ID
	DeleteTable bool `json:"deleteTable" form:"deleteTable"` // 是否删除表
}
