package dto

import (
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/dto"
	common "github.com/mars-projects/mars/common/models"
)

type SysDictDataGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int    `json:"id" search:"type:exact;column:dict_code;table:sys_dict_data" comment:""`
	DictLabel      string `json:"dictLabel" search:"type:contains;column:dict_label;table:sys_dict_data" comment:""`
	DictValue      string `json:"dictValue" search:"type:contains;column:dict_value;table:sys_dict_data" comment:""`
	DictType       string `json:"dictType" search:"type:contains;column:dict_type;table:sys_dict_data" comment:""`
	Status         string `json:"status" search:"type:exact;column:status;table:sys_dict_data" comment:""`
}

func (m *SysDictDataGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysDictDataGetAllResp struct {
	DictLabel string `json:"label"`
	DictValue string `json:"value"`
}

type SysDictDataInsertReq struct {
	Id        int    `json:"-" comment:""`
	DictSort  int    `json:"dictSort" comment:""`
	DictLabel string `json:"dictLabel" comment:""`
	DictValue string `json:"dictValue" comment:""`
	DictType  string `json:"dictType" comment:""`
	CssClass  string `json:"cssClass" comment:""`
	ListClass string `json:"listClass" comment:""`
	IsDefault string `json:"isDefault" comment:""`
	Status    int    `json:"status" comment:""`
	Default   string `json:"default" comment:""`
	Remark    string `json:"remark" comment:""`
	common.ControlBy
}

func (s *SysDictDataInsertReq) Generate(model *models.SysDictData) {
	model.DictCode = s.Id
	model.DictSort = s.DictSort
	model.DictLabel = s.DictLabel
	model.DictValue = s.DictValue
	model.DictType = s.DictType
	model.CssClass = s.CssClass
	model.ListClass = s.ListClass
	model.IsDefault = s.IsDefault
	model.Status = s.Status
	model.Default = s.Default
	model.Remark = s.Remark
}

func (s *SysDictDataInsertReq) GetId() interface{} {
	return s.Id
}

type SysDictDataUpdateReq struct {
	Id        int    `json:"dictCode" comment:""`
	DictSort  int    `json:"dictSort" comment:""`
	DictLabel string `json:"dictLabel" comment:""`
	DictValue string `json:"dictValue" comment:""`
	DictType  string `json:"dictType" comment:""`
	CssClass  string `json:"cssClass" comment:""`
	ListClass string `json:"listClass" comment:""`
	IsDefault string `json:"isDefault" comment:""`
	Status    int    `json:"status" comment:""`
	Default   string `json:"default" comment:""`
	Remark    string `json:"remark" comment:""`
	common.ControlBy
}

func (s *SysDictDataUpdateReq) Generate(model *models.SysDictData) {
	model.DictCode = s.Id
	model.DictSort = s.DictSort
	model.DictLabel = s.DictLabel
	model.DictValue = s.DictValue
	model.DictType = s.DictType
	model.CssClass = s.CssClass
	model.ListClass = s.ListClass
	model.IsDefault = s.IsDefault
	model.Status = s.Status
	model.Default = s.Default
	model.Remark = s.Remark
}

func (s *SysDictDataUpdateReq) GetId() interface{} {
	return s.Id
}

type SysDictDataGetReq struct {
	Id int `json:"id"`
}

func (s *SysDictDataGetReq) GetId() interface{} {
	return s.Id
}

type SysDictDataDeleteReq struct {
	Ids              []int `json:"ids"`
	common.ControlBy `json:"-"`
}

func (s *SysDictDataDeleteReq) GetId() interface{} {
	return s.Ids
}
