package gen

var serviceDefaultTemplate = `
package internal

import (
	"server/app/model"
	"server/app/system/{{$.SystemName}}/define"

	"github.com/iWinston/qk-library/frame/q"
	"github.com/iWinston/qk-library/frame/qmodel"
)

type {{$.CamelName}}Service struct{}

// Post 新增{{$.Description}}
func (s *{{$.CamelName}}Service) Post(ctx *qmodel.ReqContext, param interface{}) error {
	var (
		m = &model.{{$.ModelName}}{}
		err = q.Post(ctx.DB.WithContext(ctx.RCtx), m, param)
	)
	return err
}

// Get {{$.Description}}详情
func (s *{{$.CamelName}}Service) Get(ctx *qmodel.ReqContext, param interface{}) (*define.{{$.CamelName}}GetRes, error) {
	var (
		res = &define.{{$.CamelName}}GetRes{}
		tx = ctx.DB.Model(&model.{{$.ModelName}}{})
		err = q.Get(tx, param, res)
	)
	return res, err
}

// Patch 修改{{$.Description}}
func (s *{{$.CamelName}}Service) Patch(ctx *qmodel.ReqContext, param interface{}) (error) {
	var (
		m = &model.{{$.ModelName}}{}
		tx = ctx.DB.WithContext(ctx.RCtx).Model(m)
		err = q.Patch(tx, m, param)
	)
	return err
}

// Delete 删除{{$.Description}}
func (s *{{$.CamelName}}Service) Delete(ctx *qmodel.ReqContext, param interface{}) (error) {
	var (
		m = &model.{{$.ModelName}}{}
		tx = ctx.DB.WithContext(ctx.RCtx).Model(m)
		err = q.Delete(tx, m, param)
	)
	return err
}

// List {{$.Description}}列表
func (s *{{$.CamelName}}Service) List(ctx *qmodel.ReqContext, param interface{}) (*[]define.{{$.CamelName}}ListRes, int64, error) {
	var (
		total int64
		res   = &[]define.{{$.CamelName}}ListRes{}
		tx    = ctx.DB.Model(&model.{{$.ModelName}}{})
		err   = q.List(tx, param, res, &total)
	)
	return res, total, err
}`

var serviceIndexTemplate = `
package service

var {{$.CamelName}} = &{{$.Name}}Service{}

type {{$.Name}}Service struct {}
`

var serviceTemplateMap = map[string]string{
	"default": serviceDefaultTemplate,
	"index":   serviceIndexTemplate,
}
