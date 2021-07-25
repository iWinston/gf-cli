package gen

var apiInternalTemplate = `
package api

import (
	"server/app/system/{SystemName}/define"
	"server/app/system/{SystemName}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
)

// {Description}API管理对象
var {CamelName} = &{Name}Api{}

type {Name}Api struct{}

// @summary {Description}新增接口
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{CamelName}PostParam true "新增{Description}"
// @router  /{SystemName}/{Name} [POST]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) Post(r *ghttp.Request) {
	param := &define.{CamelName}PostParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Create(param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary {Description}详情接口
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{SystemName}/{Name} [GET]
// @success 200 {object} q.JsonResponse{data=define.{CamelName}GetRes} "执行结果"
func (a *{Name}Api) Get(r *ghttp.Request) {
	param := &define.{CamelName}GetParam{}
	q.AssignParamFormReq(r, param)
	res, err := service.{CamelName}.Get(param)
	err = q.OptimizeDbErr(err)
	q.ResponseWithData(r, err, res)
}

// @summary {Description}修改接口
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{CamelName}PatchParam true "修改内容"
// @router  /{SystemName}/{Name} [Patch]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) Patch(r *ghttp.Request) {
	param := &define.{CamelName}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Patch(param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary {Description}删除接口
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{SystemName}/{Name} [Delete]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) Delete(r *ghttp.Request) {
	param := &define.{CamelName}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := service.{CamelName}.Delete(param)
	err = q.OptimizeDbErr(err)
	q.Response(r, err)
}

// @summary {Description}列表接口
// @tags    {Description}管理
// @produce  json
// @security ApiKeyAuth
// @param entity query define.{CamelName}ListParam true "分页"
// @router  /{Name} [GET]
// @success 200 {object} q.JsonResponseWithMeta{[]define.{CamelName}ListRes} "执行结果"
func (a *{Name}Api) List(r *ghttp.Request) {
	param := &define.{CamelName}ListParam{}
	q.AssignParamFormReq(r, param)
	res, total, err := service.{CamelName}.List(param)
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
