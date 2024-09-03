package demo

import "github.com/laios-admin/service"

type ApiGroup struct {
	DemoApi
}

var (
	demoService = service.ServiceGroupApp.DemoServiceGroup.DemoService
)
