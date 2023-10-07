package request

import (
	"github.com/liuxing95/gin-react-admin/model/common/request"
	"github.com/liuxing95/gin-react-admin/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
