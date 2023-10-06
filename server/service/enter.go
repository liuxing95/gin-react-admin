package service

import (
	"github.com/liuxing95/gin-react-admin/service/example"
	"github.com/liuxing95/gin-react-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
