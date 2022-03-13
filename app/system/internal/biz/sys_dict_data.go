package biz

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	cDto "github.com/mars-projects/mars/lib/dto"

	"gorm.io/gorm"
)

type SysDictData struct {
	orm *gorm.DB
	log *log.Helper
}

// GetPage 获取列表
func (e *SysDictData) GetPage(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData, count *int64) error {
	var err error
	var data models.SysDictData

	err = e.orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取对象
func (e *SysDictData) Get(d *dto.SysDictDataGetReq, model *models.SysDictData) error {
	var err error
	var data models.SysDictData

	db := e.orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.log.Errorf("db error: %s", err)
		return err
	}
	if db.Error != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建对象
func (e *SysDictData) Insert(c *dto.SysDictDataInsertReq) error {
	var err error
	var data = new(models.SysDictData)
	c.Generate(data)
	err = e.orm.Create(data).Error
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改对象
func (e *SysDictData) Update(c *dto.SysDictDataUpdateReq) error {
	var err error
	var model = models.SysDictData{}
	e.orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.orm.Save(model)
	if err = db.Error; err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// Remove 删除
func (e *SysDictData) Remove(c *dto.SysDictDataDeleteReq) error {
	var err error
	var data models.SysDictData

	db := e.orm.Delete(&data, c.GetId())
	if db.Error != nil {
		err = db.Error
		e.log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetAll 获取所有
func (e *SysDictData) GetAll(c *dto.SysDictDataGetPageReq, list *[]models.SysDictData) error {
	var err error
	var data models.SysDictData

	err = e.orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
