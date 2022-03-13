package biz

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/models"

	"github.com/mars-projects/mars/conf"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

// ProviderBizSet is biz providers.
var ProviderBizSet = wire.NewSet(NewBizsOption)

type BizsOption struct {
	SysUser     *SysUser
	SysRole     *SysRole
	SysMenu     *SysMenu
	SysConfig   *SysConfig
	SysDept     *SysDept
	SysDictType *SysDictType
	SysDictData *SysDictData
	SysPost     *SysPost
}

func NewBizsOption(data *conf.Data, db *gorm.DB, logger log.Logger) *BizsOption {
	if data.Database.Migrate {
		err := migrate(db)
		if err != nil {
			panic(err)
		}
	}
	return &BizsOption{
		SysUser:     NewSysUserBiz(db, logger),
		SysRole:     NewSysRoleBiz(db, logger),
		SysMenu:     NewSysMenuBiz(db, logger),
		SysConfig:   NewSysConfigBiz(db, logger),
		SysDept:     NewSysDeptBiz(db, logger),
		SysDictType: NewSysDictTypeBiz(db, logger),
		SysDictData: NewSysDictDataBiz(db, logger),
		SysPost:     NewSysPostBiz(db, logger),
	}
}

func NewSysUserBiz(db *gorm.DB, logger log.Logger) *SysUser {
	return &SysUser{orm: db, log: log.NewHelper(logger)}
}
func NewSysConfigBiz(db *gorm.DB, logger log.Logger) *SysConfig {
	return &SysConfig{orm: db, log: log.NewHelper(logger)}
}

func NewSysDictTypeBiz(db *gorm.DB, logger log.Logger) *SysDictType {
	return &SysDictType{orm: db, log: log.NewHelper(logger)}
}

func NewSysDictDataBiz(db *gorm.DB, logger log.Logger) *SysDictData {
	return &SysDictData{orm: db, log: log.NewHelper(logger)}
}

func NewSysDeptBiz(db *gorm.DB, logger log.Logger) *SysDept {
	return &SysDept{orm: db, log: log.NewHelper(logger)}
}

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(
		new(models.SysConfig),
		new(models.SysDept),
		new(models.SysUser),
		new(models.SysDictData),
		new(models.SysDictType),
		new(models.SysRole),
		new(models.SysPost),
	)
	if err != nil {
		return err
	}
	filePath := "../../db.sql"
	err = ExecSql(db, filePath)
	if err != nil {
		return err
	}
	return nil
}

func ExecSql(db *gorm.DB, filePath string) error {
	sql, err := Ioutil(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error())
		return err
	}
	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList)-1; i++ {
		if strings.Contains(sqlList[i], "--") {
			fmt.Println(sqlList[i])
			continue
		}
		sql := strings.Replace(sqlList[i]+";", "\n", "", 0)
		sql = strings.TrimSpace(sql)
		if err = db.Exec(sql).Error; err != nil {
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}
	return nil
}

func Ioutil(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:", result)
		return result, nil
	} else {
		return "", err
	}
}
