package service

import "github.com/laios-admin/service/demo"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	DemoServiceGroup demo.ServiceGroup
}
