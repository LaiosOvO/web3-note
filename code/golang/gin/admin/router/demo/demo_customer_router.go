package demo

import (
	"github.com/gin-gonic/gin"
)

type DemoCustomerRouter struct{}

func (e *DemoCustomerRouter) InitDemoCustomerRouter(Router *gin.RouterGroup) {

	demoCustomerRouter := Router.Group("demo")
}
