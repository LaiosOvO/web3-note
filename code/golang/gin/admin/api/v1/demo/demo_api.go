package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/laios-admin/model/demo"
)

type DemoApi struct{}

func (e *DemoApi) CreateDemo(c *gin.Context) {
	var demoEntity demo.DemoEntity

	err := c.ShouldBindJSON(demoEntity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = demoService.CreateDemo(demoEntity)
	if err != nil {

		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (d *DemoApi) DeleteDemo(c *gin.Context) {

	var demoEntity demo.DemoEntity
	err := c.ShouldBindJSON(demoEntity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = demoService.DeleteDemo(demoEntity)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (d *DemoApi) PageList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dataList, total, err := demoService.Page(1, pageInfo)

	response.OkWithDetailed(response.PageResult{
		List: dataList,
	})
}
