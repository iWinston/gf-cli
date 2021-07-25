package gen

var serviceDefaultTemplate = `
package internal

import (
	"server/app/model"
	"server/app/system/{SystemName}/define"

	"github.com/iWinston/qk-library/frame/q"
)

type {CamelName}Service struct{}

// Post 新增{Description}
func (s *{CamelName}Service) Post(ctx *qmodel.QContext, param *define.{CamelName}PostParam) error {
	var (
		m = &model.{CamelName}{}
		err = q.Post(ctx.DB, m, param)
	)
	return err
}

// Get {Description}详情
func (s *{CamelName}Service) Get(ctx *qmodel.QContext, param *define.{CamelName}GetParam) (*define.{CamelName}GetRes, error) {
	var (
		res = &define.{CamelName}GetRes{}
		tx = ctx.DB.Model(&model.{CamelName}{})
		err = q.Get(tx, param, res)
	)
	return res, err
}

// Patch 修改{Description}
func (s *{CamelName}Service) Patch(ctx *qmodel.QContext, param *define.{CamelName}PatchParam) (error) {
	var (
		m = &model.{CamelName}{}
		tx = ctx.DB.Model(m)
		err = q.Patch(tx, m, param)
	)
	return err
}

// Delete 删除{Description}
func (s *{CamelName}Service) Delete(ctx *qmodel.QContext, param *define.{CamelName}DeleteParam) (error) {
	var (
		m = &model.{CamelName}{}
		tx = ctx.DB.Model(m)
		err = q.Delete(tx, m, param)
	)
	return err
}

// List {Description}列表
func (s *{CamelName}Service) List(param *define.{CamelName}ListParam) (ctx *qmodel.QContext, *[]define.{CamelName}ListRes, int64, error) {
	var (
		total int64
		res   = &[]define.{CamelName}ListRes{}
		tx    = ctx.DB.Model(&model.{CamelName}{})
		err   = q.List(tx, param, res, &total)
	)
	return res, total, err
}`

var serviceIndexTemplate = `
package service

import "server/app/system/{SystemName}/service/internal"

var {CamelName} = &{Name}Service{}

type {Name}Service struct {
	*internal.{CamelName}Service
}
`

var serviceTemplateMap = map[string]string{
	"default": serviceDefaultTemplate,
	"index":   serviceIndexTemplate,
}
