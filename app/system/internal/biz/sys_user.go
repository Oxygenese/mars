package biz

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/api/system"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	cDto "github.com/mars-projects/mars/lib/dto"
	"github.com/mars-projects/mars/lib/utils"
	"gorm.io/gorm"
)

type SysUser struct {
	orm *gorm.DB
	log *log.Helper
}

func (e *SysUser) FindByUsername(req *system.SysUserInfoReq, reply *system.SysUserReply) error {
	var data models.SysUser
	var err error
	err = e.orm.Model(&data).Debug().
		Where("username = ?", req.Username).
		First(reply).
		Error
	return err
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, list *[]models.SysUser, count *int64) error {
	var err error
	err = e.orm.Debug().Preload("Dept").
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

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, model *models.SysUser) error {
	var data models.SysUser

	err := e.orm.Debug().Model(&data).Debug().
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = e.orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.orm.Create(&data).Error
	if err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq) error {
	var err error
	var model models.SysUser
	db := e.orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := e.orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		e.log.Warnf("db update error")
		return err
	}
	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq) error {
	var err error
	var model models.SysUser
	db := e.orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	err = e.orm.Save(&model).Error
	if err != nil {
		e.log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq) error {
	var err error
	var model models.SysUser
	db := e.orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	err = e.orm.Save(&model).Error
	if err != nil {
		e.log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq) error {
	var err error
	var model models.SysUser
	db := e.orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = e.orm.Save(&model).Error
	if err != nil {
		e.log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *dto.SysUserById) error {
	var err error
	var data models.SysUser

	db := e.orm.Model(&data).
		Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.orm.Model(c).
		Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		e.log.Errorf("db error: %s", err)
		return err
	}
	var ok bool
	ok, err = utils.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		e.log.Errorf("CompareHashAndPassword error, %s", err.Error())
		return err
	}
	if !ok {
		err = errors.New("incorrect Password")
		e.log.Warnf("user[%d] %s", id, err.Error())
		return err
	}
	c.Password = newPassword
	db := e.orm.Model(c).Where("user_id = ?", id).Select("Password", "Salt").Updates(c)
	if err = db.Error; err != nil {
		e.log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		e.log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := e.orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = e.orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}
	err = e.orm.Find(posts, user.PostIds).Error
	if err != nil {
		return err
	}
	return nil
}
