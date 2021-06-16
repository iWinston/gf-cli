package gen

var serviceDefaultTemplate = `
package internal

import (
	"server/app/model"
	"server/app/system/admin/define"
	"server/library/restful"

	"github.com/gogf/gf/util/gconv"
)

type {TplUpperName}Service struct{}

func (s *{TplUpperName}Service) Create(param *define.{TplUpperName}CreateParam) error {
	var {TplName} *model.{TplUpperName}
	gconv.Struct(param, &{TplName})
	sql := model.DB.Model(&model.{TplUpperName}{})
	return restful.CreateOne(sql, {TplName})
}

func (s *{TplUpperName}Service) FindOne(id uint, param *define.{TplUpperName}FindOneParam) (*define.{TplUpperName}FindOneRes, error) {
	res := &define.{TplUpperName}FindOneRes{}
	sql := restful.GenSqlByParam(model.DB.Model(&model.{TplUpperName}{}), param)
	sql = restful.GenSqlByRes(sql, res)
	err := restful.TakeOne(sql, id, res)
	return res, err
}

func (s *{TplUpperName}Service) Find(param *define.{TplUpperName}FindParam) (*[]define.{TplUpperName}FindRes, int64, error) {
	countSql := restful.GenSqlByParam(model.DB.Model(&model.{TplUpperName}{}), param)
	var total int64
	if result := countSql.Count(&total); result.Error != nil {
		return nil, 0, result.Error
	}

	res := &[]define.{TplUpperName}FindRes{}
	sql := restful.GenSqlByParam(model.DB.Model(&model.{TplUpperName}{}), param)
	sql = restful.GenSqlByRes(sql, &define.{TplUpperName}FindRes{})
	err := restful.Find(sql, param, res)
	return res, total, err
}

func (s *{TplUpperName}Service) PatchOne(id uint, param *define.{TplUpperName}PatchOneParam) error {
	var {TplName} *model.{TplUpperName}
	gconv.Struct(param, &{TplName})
	sql := restful.GenSqlByParam(model.DB.Debug().Model(&model.{TplUpperName}{}), param)
	return restful.PatchOne(sql, id, {TplName})
}

func (s *{TplUpperName}Service) DeleteOne(id uint, param *define.{TplUpperName}DeleteOneParam) error {
	{TplName} := &model.{TplUpperName}{}
	sql := restful.GenSqlByParam(model.DB.Model(&model.{TplUpperName}{}), param)
	err := restful.TakeOne(sql, id, {TplName})
	if err != nil {
		return err
	}
	return restful.DeleteOne(sql, id, {TplName})
}
`

var serviceIndexTemplate = `
package service

import "server/app/system/admin/service/internal"

var {TplUpperName} = &{TplName}Service{}

type {TplName}Service struct {
	*internal.{TplUpperName}Service
}
`

var serviceTemplateMap = map[string]string{
	"default": serviceDefaultTemplate,
	"index":   serviceIndexTemplate,
}
