package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/laios-admin/global"
	"github.com/laios-admin/model/demo"
)

type DemoService struct{}

var DemoServiceApp = new(DemoService)

func (dms *DemoService) CreateDemo(d demo.DemoEntity) (err error) {
	err = global.GVA_DB.Create(&d).Error
	return err
}

func (dms *DemoService) DeleteDemo(d demo.DemoEntity) (err error) {
	err = global.GVA_DB.Delete(d).Error
	return err
}

func (dms *DemoService) UpdateDemo(d demo.DemoEntity) (err error) {
	err = global.GVA_DB.Save(d).Error
	return err
}

func (dms *DemoService) GetDemo(id int) (demo demo.DemoEntity, err error) {
	err = global.GVA_DB.Where("id =?", id).First(&demo).Error
	return demo, err
}

func (dms *DemoService) Page(id int, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&demo.DemoEntity{})

	var DemoList []demo.DemoEntity
	err = db.Where("sys_user_authority_id in ?", []int{1, 2, 3, 4}).Count(&total).Error
	if err != nil {
		return DemoList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", []int{1, 2, 3, 4}).Find(&DemoList).Error
	}

	return DemoList, total, err
}
