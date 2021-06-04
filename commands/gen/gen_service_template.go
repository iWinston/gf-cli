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

func (s *{TplUpperName}Service) Create(req *define.{TplUpperName}CreateDto) error {
	var {TplName} *define.{TplUpperName}
	gconv.Struct(req, &{TplName})
	return s.{TplUpperName}.Create({TplName})
}

func (s *{TplUpperName}Service) GetOne(id uint) (*define.{TplUpperName}GetOneRes, error) {
	sql := model.GenSqlByRes("{TplName}", &define.{TplUpperName}GetOneRes{})
	return s.{TplUpperName}.GetOne(id, sql)
}

func (s *{TplUpperName}Service) Get(req *define.{TplUpperName}GetDto) (*[]define.{TplUpperName}GetRes, int64, error) {
	sql := model.GenSqlByRes("{TplName}", &define.{TplUpperName}GetRes{})
	return s.{TplUpperName}.Get(req, model.DB, sql)
}

func (s *{TplUpperName}Service) PatchOne(id uint, req *define.{TplUpperName}PatchOneDto) error {
	var {TplName} *define.{TplUpperName}
	gconv.Struct(req, &{TplName})
	return s.{TplUpperName}.PatchOne(id, {TplName})
}

func (s *{TplUpperName}Service) Patch(req *define.{TplUpperName}PatchDto) error {
	return s.{TplUpperName}.Patch(req, unit.StructToMap(req))
}

func (s *{TplUpperName}Service) DeleteOne(id uint) error {
	return s.{TplUpperName}.DeleteOne(id)
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
	if result := tx.Model(&define.{TplUpperName}{}).Take(&{TplName}, id); result.Error != nil {
		return nil, result.Error
	} else {
		return &{TplName}, nil
	}
}

func (s *{TplUpperName}) Get(req *define.{TplUpperName}GetDto, countTx *gorm.DB, findTx *gorm.DB) (*[]define.{TplUpperName}GetRes, int64, error) {
	var total int64
	if result := countTx.
		Model(&define.{TplUpperName}{}).
		Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	var res []define.{TplUpperName}GetRes
	if result := findTx.
		Model(&define.{TplUpperName}{}).
		Scopes(model.Paginate(req)).
		Find(&res); result.Error != nil {
		return nil, 0, result.Error
	} else {
		return &res, total, nil
	}

}

func (s *{TplUpperName}) PatchOne(id uint, {TplName} *define.{TplUpperName}) error {
	if result := model.DB.Model(&define.{TplUpperName}{}).Where("id = ?", id).Updates({TplName}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *{TplUpperName}) Patch(req *define.{TplUpperName}PatchDto, toUpdate interface{}) error {
	if result := model.DB.Model(&define.{TplUpperName}{}).Where("id In ?", req.GetIdSlice()).Updates(toUpdate); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *{TplUpperName}) DeleteOne(id uint) error {
	if result := model.DB.Delete(&define.{TplUpperName}{},id); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
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
