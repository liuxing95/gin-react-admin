package request

import (
	"github.com/liuxing95/gin-react-admin/model/common/request"
	"github.com/liuxing95/gin-react-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
