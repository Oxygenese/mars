package models

import (
	models2 "github.com/mars-projects/mars/common/models"
)

type SysConfig struct {
	models2.Model
	ConfigName  string `json:"configName" gorm:"size:128;comment:ConfigName"`   //
	ConfigKey   string `json:"configKey" gorm:"size:128;comment:ConfigKey"`     //
	ConfigValue string `json:"configValue" gorm:"size:255;comment:ConfigValue"` //
	ConfigType  string `json:"configType" gorm:"size:64;comment:ConfigType"`
	IsFrontend  int    `json:"isFrontend" gorm:"size:64;comment:是否前台"` //
	Remark      string `json:"remark" gorm:"size:128;comment:Remark"`  //
	models2.ControlBy
	models2.ModelTime
}

func (SysConfig) TableName() string {
	return "sys_config"
}

func (e *SysConfig) Generate() models2.ActiveRecord {
	o := *e
	return &o
}

func (e *SysConfig) GetId() interface{} {
	return e.Id
}
