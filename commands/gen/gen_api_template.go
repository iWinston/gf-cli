package gen

var apiIndexTemplate = `
package api

import (
	"server/app/system/{TplSystemName}/define"
	"server/app/system/{TplSystemName}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
)

// {TplDescription}API管理对象
var {TplUpperName} = &{Name}Api{}

type {Name}Api struct{}

// @summary {TplDescription}新增接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{TplUpperName}PostParam true "新增{TplDescription}"
// @router  /{Name} [POST]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) Post(r *ghttp.Request) {
	param := &define.{TplUpperName}PostParam{}
	q.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.Create(param)
	q.Response(r, err)
}

// @summary {TplDescription}详情接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{Name} [GET]
// @success 200 {object} define.{TplUpperName}FindOneRes "执行结果"
func (a *{Name}Api) Get(r *ghttp.Request) {
	param := &define.{TplUpperName}GetParam{}
	q.AssignParamFormReq(r, param)
	res, err := service.{TplUpperName}.Get(param)
	q.ResponseWithData(r, err, res)
}

// @summary {TplDescription}列表接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity query define.{TplUpperName}FindParam true "分页"
// @router  /{Name} [GET]
// @success 200 {object} []define.{TplUpperName}FindRes "执行结果"
func (a *{Name}Api) List(r *ghttp.Request) {
	param := &define.{TplUpperName}FindParam{}
	q.AssignParamFormReq(r, param)
	res, total, err := service.{TplUpperName}.List(param)
	q.ResponseWithMeta(r, err, res, total)
}

// @summary {TplDescription}修改接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{TplUpperName}PatchOneParam true "修改内容"
// @router  /{Name} [Patch]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) PatchOne(r *ghttp.Request) {
	id := q.GetIdFormReq(r)
	param := &define.{TplUpperName}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.Patch(param)
	q.Response(r, err)
}

// @summary {TplDescription}删除接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{Name} [Delete]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{Name}Api) Delete(r *ghttp.Request) {
	id := q.GetIdFormReq(r)
	param := &define.{TplUpperName}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := service.{TplUpperName}.Delete(param)
	q.Response(r, err)
}
`

var apiTemplateMap = map[string]string{
	"index": apiIndexTemplate,
}
