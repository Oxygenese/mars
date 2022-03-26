package biz

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	cDto "github.com/mars-projects/mars/common/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SysRole struct {
	orm   *gorm.DB
	log   *log.Helper
	Error error
}

func NewSysRoleBiz(db *gorm.DB, logger log.Logger) *SysRole {
	return &SysRole{
		orm: db,
		log: log.NewHelper(logger),
	}
}

func (e SysRole) AddError(err error) error {
	e.Error = err
	return err
}

// GetPage 获取SysRole列表
func (e *SysRole) GetPage(c *dto.SysRoleGetPageReq, list *[]models.SysRole, count *int64) error {
	var err error
	var data models.SysRole
	err = e.orm.Model(&data).Preload("SysMenu").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysRole对象
func (e *SysRole) Get(d *dto.SysRoleGetReq, model *models.SysRole) error {
	var err error
	db := e.orm.First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.log.Errorf("db error:%s", err)
		return err
	}
	if err != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		e.log.Errorf("get menuIds error, %s", err.Error())
		return err
	}
	return nil
}

// Insert 创建SysRole对象
func (e *SysRole) Insert(c *dto.SysRoleInsertReq) error {
	var err error
	var data models.SysRole
	var dataMenu []models.SysMenu
	err = e.orm.Where("menu_id in ?", c.MenuIds).Find(&dataMenu).Error
	if err != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	c.SysMenu = dataMenu
	c.Generate(&data)
	tx := e.orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysRole对象
func (e *SysRole) Update(c *dto.SysRoleUpdateReq) error {
	var err error
	tx := e.orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	var mlist = make([]models.SysMenu, 0)
	tx.Preload("SysMenu").First(&model, c.GetId())
	tx.Preload("SysApi").Where("menu_id in ?", c.MenuIds).Find(&mlist)
	err = tx.Model(&model).Association("SysMenu").Delete(model.SysMenu)
	if err != nil {
		e.log.Errorf("delete policy error:%s", err)
		return err
	}
	c.Generate(&model)
	model.SysMenu = &mlist
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)

	if db.Error != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	if err != nil {
		e.log.Errorf("delete policy error:%s", err)
		return err
	}
	return nil
}

// Remove 删除SysRole
func (e *SysRole) Remove(c *dto.SysRoleDeleteReq) error {
	var err error
	e.log.Info("[sys_role] delete ids:%s", c)
	tx := e.orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.Preload("SysMenu").Preload("SysDept").First(&model, c.GetId())
	db := tx.Select(clause.Associations).Delete(&model)
	if db.Error != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// GetRoleMenuId 获取角色对应的菜单ids
func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil
}

func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) *SysRole {
	var err error
	tx := e.orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var dlist = make([]models.SysDept, 0)
	var model = models.SysRole{}
	tx.Preload("SysDept").First(&model, c.RoleId)
	tx.Where("id in ?", c.DeptIds).Find(&dlist)
	err = tx.Model(&model).Association("SysDept").Delete(model.SysDept)
	if err != nil {
		e.log.Errorf("delete SysDept error:%s", err)
		_ = e.AddError(err)
		return e
	}
	c.Generate(&model)
	model.SysDept = dlist
	db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if db.Error != nil {
		e.log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if db.RowsAffected == 0 {
		_ = e.AddError(errors.New("无权更新该数据"))
		return e
	}
	return e
}

// UpdateStatus 修改SysRole对象status
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) error {
	var err error
	tx := e.orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	tx.First(&model, c.GetId())
	c.Generate(&model)
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if db.Error != nil {
		e.log.Errorf("db error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// GetWithName 获取SysRole对象
func (e *SysRole) GetWithName(d *dto.SysRoleByName, model *models.SysRole) *SysRole {
	var err error
	db := e.orm.Where("role_name = ?", d.RoleName).First(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	if err != nil {
		e.log.Errorf("db error:%s", err)
		_ = e.AddError(err)
		return e
	}
	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		e.log.Errorf("get menuIds error, %s", err.Error())
		_ = e.AddError(err)
		return e
	}
	return e
}

// GetById 获取SysRole对象
func (e *SysRole) GetById(roleId int) ([]string, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, nil
}
