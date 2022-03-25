package models

import (
	models2 "github.com/mars-projects/mars/common/models"
)

type SysDictType struct {
	ID       int    `json:"id" gorm:"primaryKey;column:dict_id;autoIncrement;comment:主键编码"`
	DictName string `json:"dictName" gorm:"size:128;comment:DictName"`
	DictType string `json:"dictType" gorm:"size:128;comment:DictType"`
	Status   int    `json:"status" gorm:"size:4;comment:Status"`
	Remark   string `json:"remark" gorm:"size:255;comment:Remark"`
	models2.ControlBy
	models2.ModelTime
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}

func (e *SysDictType) Generate() models2.ActiveRecord {
	o := *e
	return &o
}

func (e *SysDictType) GetId() interface{} {
	return e.ID
}
