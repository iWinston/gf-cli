package gen

var apiInternalTemplate = `
package internal

import (
	"server/app/system/{SystemName}/define"
	"server/app/system/{SystemName}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
	"github.com/iWinston/qk-library/frame/qservice"
)

// {Description}API
type {CamelName}Api struct{}

// @summary 新增{Description}
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{CamelName}PostParam true "新增{Description}"
// @router  /{SystemName}/{Name}/post [POST]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{CamelName}Api) Post(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{CamelName}PostParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Post(ctx, param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary {Description}详情
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{CamelName}GetParam true "{Description}详情"
// @router  /{SystemName}/{Name}/get [POST]
// @success 200 {object} q.JsonResponseWithData{data=define.{CamelName}GetRes} "执行结果"
func (a *{CamelName}Api) Get(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{CamelName}GetParam{}
	q.AssignParamFormReq(r, param)
	res, err := service.{CamelName}.Get(ctx, param)
	err = q.OptimizeDbErr(err)
	q.ResponseWithData(r, err, res)
}

// @summary 修改{Description}
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{CamelName}PatchParam true "修改内容"
// @router  /{SystemName}/{Name}/patch [POST]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{CamelName}Api) Patch(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{CamelName}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Patch(ctx, param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary 删除{Description}
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{CamelName}DeleteParam true "删除{Description}"
// @router  /{SystemName}/{Name}.delete [POST]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{CamelName}Api) Delete(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{CamelName}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Delete(ctx, param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary {Description}列表
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{CamelName}ListParam true "分页"
// @router  /{SystemName}/{Name}/list [POST]
// @success 200 {object} q.JsonResponseWithMeta{data=[]define.{CamelName}ListRes} "执行结果"
func (a *{CamelName}Api) List(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{CamelName}ListParam{}
	q.AssignParamFormReq(r, param)
	res, total, err := service.{CamelName}.List(ctx, param)
	err = q.OptimizeDbErr(err)
	q.ResponseWithMeta(r, err, res, total)
}`

var apiIndexTemplate = `package api

import "server/app/system/{SystemName}/api/internal"

var {CamelName} = &{Name}Api{}

type {Name}Api struct {
	*internal.{CamelName}Api
}`

var apiTemplateMap = map[string]string{
	"index":    apiIndexTemplate,
	"internal": apiInternalTemplate,
}
