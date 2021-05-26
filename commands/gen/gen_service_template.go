package gen

var serviceDefaultTemplate = `
package service

import (
	"server/app/model"
	"server/app/system/admin/define"
	"server/library/unit"

	"github.com/gogf/gf/util/gconv"
)

type {TplUpperName}Service struct {
	*{TplUpperName}
}

func (s *{TplUpperName}Service) Create(req *define.{TplUpperName}CreateReq) error {
	var {TplName} *define.{TplUpperName}
	gconv.Struct(req, &{TplName})
	return s.{TplUpperName}.Create({TplName})
}

func (s *{TplUpperName}Service) GetOne(id uint) (*define.{TplUpperName}GetOneRes, error) {
	return s.{TplUpperName}.GetOne(id, model.DB.Model(&define.{TplUpperName}{}))
}

func (s *{TplUpperName}Service) Get(req *define.{TplUpperName}GetReq) (*[]define.{TplUpperName}GetRes, int64, error) {
	return s.{TplUpperName}.Get(req, model.DB.Model(&define.{TplUpperName}{}))
}

func (s *{TplUpperName}Service) PatchOne(id uint, req *define.{TplUpperName}PatchOneReq) error {
	return s.{TplUpperName}.PatchOne(id, req, unit.StructToMap(req))
}

func (s *{TplUpperName}Service) Patch(req *define.{TplUpperName}PatchReq) error {
	return s.{TplUpperName}.Patch(req, unit.StructToMap(req))
}
`

var serviceCoreTemplate = `
package service

import (
	"server/app/model"
	"server/app/system/admin/define"

	"gorm.io/gorm"
)

type {TplUpperName} struct {
}

func (s *{TplUpperName}) Create({TplName} *define.{TplUpperName}) error {
	if result := model.DB.Create(&{TplName}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *{TplUpperName}) GetOne(id uint, tx *gorm.DB) (*define.{TplUpperName}GetOneRes, error) {
	var {TplName} define.{TplUpperName}GetOneRes
	if result := tx.Take(&{TplName}, id); result.Error != nil {
		return nil, result.Error
	} else {
		return &{TplName}, nil
	}
}

func (s *{TplUpperName}) Get(req *define.{TplUpperName}GetReq, tx *gorm.DB) (*[]define.{TplUpperName}GetRes, int64, error) {
	var {TplName}s []define.{TplUpperName}GetRes
	var total int64
	if result := tx.
		Count(&total).
		Scopes(model.Paginate(req)).
		Find(&{TplName}s); result.Error != nil {
		return nil, 0, result.Error
	} else {
		return &{TplName}s, total, nil
	}
}

func (s *{TplUpperName}) PatchOne(id uint, req *define.{TplUpperName}PatchOneReq, toUpdate interface{}) error {
	if result := model.DB.Model(&define.{TplUpperName}{}).Where("id = ?", id).Updates(toUpdate); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *{TplUpperName}) Patch(req *define.{TplUpperName}PatchReq, toUpdate interface{}) error {
	if result := model.DB.Model(&define.{TplUpperName}{}).Where("id In ?", req.GetIdSlice()).Updates(toUpdate); result.Error != nil {
		return result.Error
	}
	return nil
}
`

var serviceIndexTemplate = `
package service

import (
	"server/gen/service"
)

var {TplUpperName} = &{TplName}Service{&service.{TplUpperName}Service{}}

type {TplName}Service struct {
	*service.{TplUpperName}Service
}
`

var serviceTemplateMap = map[string]string{
	"default": serviceDefaultTemplate,
	"core":    serviceCoreTemplate,
	"index":   serviceIndexTemplate,
}
