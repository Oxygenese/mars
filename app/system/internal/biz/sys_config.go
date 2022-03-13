package biz

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	cDto "github.com/mars-projects/mars/lib/dto"
	"gorm.io/gorm"
)

type SysConfig struct {
	orm *gorm.DB
	log *log.Helper
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(c *dto.SysConfigGetPageReq, list *[]models.SysConfig, count *int64) error {
	err := e.orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysConfig对象
func (e *SysConfig) Get(d *dto.SysConfigGetReq, model *models.SysConfig) error {
	err := e.orm.First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	if err != nil {
		e.log.Errorf("Service GetSysConfig error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *dto.SysConfigControl) error {
	var err error
	var data models.SysConfig
	c.Generate(&data)
	err = e.orm.Create(&data).Error
	if err != nil {
		e.log.Errorf("Service InsertSysConfig error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *dto.SysConfigControl) error {
	var err error
	var model = models.SysConfig{}
	e.orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.orm.Save(&model)
	err = db.Error
	if err != nil {
		e.log.Errorf("Service UpdateSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	return nil
}

// SetSysConfig 修改SysConfig对象
func (e *SysConfig) SetSysConfig(c *[]dto.GetSetSysConfigReq) error {
	var err error
	for _, req := range *c {
		var model = models.SysConfig{}
		e.orm.Where("config_key = ?", req.ConfigKey).First(&model)
		if model.Id != 0 {
			req.Generate(&model)
			db := e.orm.Save(&model)
			err = db.Error
			if err != nil {
				e.log.Errorf("Service SetSysConfig error:%s", err)
				return err
			}
			if db.RowsAffected == 0 {
				return errors.New("无权更新该数据")
			}
		}
	}
	return nil
}

func (e *SysConfig) GetForSet(c *[]dto.GetSetSysConfigReq) error {
	var err error
	var data models.SysConfig

	err = e.orm.Model(&data).
		Find(c).Error
	if err != nil {
		e.log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

func (e *SysConfig) UpdateForSet(c *[]dto.GetSetSysConfigReq) error {
	m := *c
	for _, req := range m {
		var data models.SysConfig
		if err := e.orm.Where("config_key = ?", req.ConfigKey).
			First(&data).Error; err != nil {
			e.log.Errorf("Service GetSysConfigPage error:%s", err)
			return err
		}
		if data.ConfigValue != req.ConfigValue {
			data.ConfigValue = req.ConfigValue

			if err := e.orm.Save(&data).Error; err != nil {
				e.log.Errorf("Service GetSysConfigPage error:%s", err)
				return err
			}
		}

	}

	return nil
}

// Remove 删除SysConfig
func (e *SysConfig) Remove(d *dto.SysConfigDeleteReq) error {
	var err error
	var data models.SysConfig

	db := e.orm.Delete(&data, d.Ids)
	if db.Error != nil {
		err = db.Error
		e.log.Errorf("Service RemoveSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetWithKey 根据Key获取SysConfig
func (e *SysConfig) GetWithKey(c *dto.SysConfigByKeyReq, resp *dto.GetSysConfigByKEYForServiceResp) error {
	var err error
	var data models.SysConfig
	err = e.orm.Debug().Model(&data).Where("config_key = ?", c.ConfigKey).First(resp).Error
	if err != nil {
		e.log.Errorf("At Service GetSysConfigByKEY Error:%s", err)
		return err
	}

	return nil
}

func (e *SysConfig) GetWithKeyList(c *dto.SysConfigGetToSysAppReq, list *[]models.SysConfig) error {
	var err error
	err = e.orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.log.Errorf("Service GetSysConfigByKey error:%s", err)
		return err
	}
	return nil
}
