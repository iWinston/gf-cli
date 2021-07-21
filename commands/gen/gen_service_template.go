package gen

var serviceDefaultTemplate = `
package internal

import (
	"server/app/model"
	"server/app/system/{TplSystemName}/define"

	"github.com/iWinston/qk-library/frame/q"
)

type {TplUpperName}Service struct{}

func (s *{TplUpperName}Service) Create(param *define.{TplUpperName}PostParam) error {
	{TplName} := &model.{TplUpperName}{}
	return q.CreateOne(model.DB, {TplName}, param)
}

func (s *{TplUpperName}Service) FindOne(id uint, param *define.{TplUpperName}FindOneParam) (*define.{TplUpperName}FindOneRes, error) {
	res := &define.{TplUpperName}FindOneRes{}
	tx := model.DB.Model(&model.{TplUpperName}{}).Where(id)
	err := q.FindOne(tx, param, res)
	return res, err
}

func (s *{TplUpperName}Service) Find(param *define.{TplUpperName}FindParam) (*[]define.{TplUpperName}FindRes, int64, error) {
	res := &[]define.{TplUpperName}FindRes{}
	tx := model.DB.Model(&model.{TplUpperName}{})
	var total int64
	err := q.Find(tx, param, res, &total)
	return res, total, err
}

func (s *{TplUpperName}Service) PatchOne(id uint, param *define.{TplUpperName}PatchOneParam) error {
	{TplName} := &model.{TplUpperName}{}
	tx := model.DB.Model({TplName}).Where(id)
	return q.PatchOne(tx, {TplName}, param)
}

func (s *{TplUpperName}Service) DeleteOne(id uint, param *define.{TplUpperName}DeleteOneParam) error {
	{TplName} := &model.{TplUpperName}{}
	tx := model.DB.Model({TplName}).Where(id)
	return q.DeleteOne(tx, {TplName}, param)
}
`

var serviceIndexTemplate = `
package service

import "server/app/system/{TplSystemName}/service/internal"

var {TplUpperName} = &{TplName}Service{}

type {TplName}Service struct {
	*internal.{TplUpperName}Service
}
`

var serviceTemplateMap = map[string]string{
	"default": serviceDefaultTemplate,
	"index":   serviceIndexTemplate,
}
