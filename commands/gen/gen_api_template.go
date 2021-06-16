package gen

var apiIndexTemplate = `
package api

import (
	"server/app/system/admin/define"
	"server/app/system/admin/service"
	"server/library/restful"

	"github.com/gogf/gf/net/ghttp"
)

// {TplDescription}API管理对象
var {TplUpperName} = &{TplName}Api{}

type {TplName}Api struct{}

// @summary {TplDescription}新增接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{TplUpperName}CreateParam true "新增{TplDescription}"
// @router  /{TplName} [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplName}Api) Post(r *ghttp.Request) {
	param := &define.{TplUpperName}CreateParam{}
	restful.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.Create(param)
	restful.Response(r, err)
}

// @summary {TplDescription}详情接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{TplName}/{id} [GET]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} define.{TplUpperName}FindOneRes "执行结果"
func (a *{TplName}Api) GetOne(r *ghttp.Request) {
	id := restful.GetIdFormReq(r)
	param := &define.{TplUpperName}FindOneParam{}
	restful.AssignParamFormReq(r, param)
	res, err := service.{TplUpperName}.FindOne(id, param)
	restful.ResponseWithData(r, err, res)
}

// @summary {TplDescription}列表接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity query define.{TplUpperName}FindParam true "分页"
// @router  /{TplName} [GET]
// @success 200 {object} response.JsonResponseWithTotal "执行结果"
func (a *{TplName}Api) Get(r *ghttp.Request) {
	param := &define.{TplUpperName}FindParam{}
	restful.AssignParamFormReq(r, param)
	res, total, err := service.{TplUpperName}.Find(param)
	restful.ResponseWithTotal(r, err, res, total)
}

// @summary {TplDescription}修改接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{TplUpperName}PatchOneParam true "修改内容"
// @router  /{TplName}/{id} [Patch]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplName}Api) PatchOne(r *ghttp.Request) {
	id := restful.GetIdFormReq(r)
	param := &define.{TplUpperName}PatchOneParam{}
	restful.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.PatchOne(id, param)
	restful.Response(r, err)
}

// @summary {TplDescription}删除接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{TplName}/{id} [Delete]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplName}Api) DeleteOne(r *ghttp.Request) {
	id := restful.GetIdFormReq(r)
	param := &define.{TplUpperName}DeleteOneParam{}
	restful.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.DeleteOne(id, param)
	restful.Response(r, err)
}
`

var apiTemplateMap = map[string]string{
	"index": apiIndexTemplate,
}
