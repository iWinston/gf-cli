package gen

var apiDefaultTemplate = `
package api
import (
	"github.com/gogf/gf/net/ghttp"
)

type {TplUpperName}Api struct {
	*{TplUpperName}
}


// @summary {TplDescription}新增接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{TplUpperName}CreateReq true "新增{TplDescription}"
// @router  /{TplName} [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplUpperName}Api) Post(r *ghttp.Request) {
	req := a.{TplUpperName}.CreateReq(r)
	a.{TplUpperName}.Post(r, req)
}

// @summary {TplDescription}列表接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity query define.{TplUpperName}GetReq true "分页"
// @router  /{TplName} [GET]
// @success 200 {object} response.JsonResponseWithTotal "执行结果"
func (a *{TplUpperName}Api) Get(r *ghttp.Request) {
	req := a.{TplUpperName}.CetReq(r)
	a.{TplUpperName}.Get(r, req)
}

// @summary {TplDescription}详情接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{TplName}/{id} [GET]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} define.{TplUpperName}GetOneRes "执行结果"
func (a *{TplUpperName}Api) GetOne(r *ghttp.Request) {
	id := a.{TplUpperName}.CetOneReq(r)
	a.{TplUpperName}.GetOne(r, id)
}

// @summary {TplDescription}修改接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{TplUpperName}PatchOneReq true "修改内容"
// @router  /{TplName}/{id} [Patch]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplUpperName}Api) PatchOne(r *ghttp.Request) {
	id, req := a.{TplUpperName}.PatchOneReq(r)
	a.{TplUpperName}.PatchOne(r, id, req)
}`

var apiCoreTemplate = `
package api

import (
	"server/app/system/admin/define"
	"server/app/system/admin/service"
	"server/library/response"

	"github.com/gogf/gf/net/ghttp"
)

type {TplUpperName} struct{}

func (a *{TplUpperName}) Post(r *ghttp.Request, req *define.{TplUpperName}CreateReq) {
	if err := service.{TplUpperName}.Create(req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *{TplUpperName}) CreateReq(r *ghttp.Request) (req *define.{TplUpperName}CreateReq) {
	if err := r.ParseForm(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return
}

func (a *{TplUpperName}) Get(r *ghttp.Request, req *define.{TplUpperName}GetReq) {
	if list, total, err := service.{TplUpperName}.Get(req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list, total)
	}
}

func (a *{TplUpperName}) CetReq(r *ghttp.Request) (req *define.{TplUpperName}GetReq) {
	if err := r.ParseQuery(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return
}

func (a *{TplUpperName}) GetOne(r *ghttp.Request, id uint) {
	if {TplName}, err := service.{TplUpperName}.GetOne(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", {TplName})
	}
}
func (a *{TplUpperName}) CetOneReq(r *ghttp.Request) (id uint) {
	if id = r.GetUint("id"); id == 0 {
		response.JsonExit(r, 1, "未获得id参数")
	}
	return
}

func (a *{TplUpperName}) PatchOne(r *ghttp.Request, id uint, req *define.{TplUpperName}PatchOneReq) {
	if err := service.{TplUpperName}.PatchOne(id, req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *{TplUpperName}) PatchOneReq(r *ghttp.Request) (id uint, req *define.{TplUpperName}PatchOneReq) {
	if id = r.GetUint("id"); id == 0 {
		response.JsonExit(r, 1, "未获得id参数")
	}
	if err := r.ParseForm(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return
}`

var apiIndexTemplate = `
package api

import (
	"server/gen/api"
)

// {TplDescription}API管理对象
var {TplUpperName} = &{TplName}Api{&api.{TplUpperName}Api{}}

type {TplName}Api struct {
	*api.{TplUpperName}Api
}

`

var apiTemplateMap = map[string]string{
	"default": apiDefaultTemplate,
	"core":    apiCoreTemplate,
	"index":   apiIndexTemplate,
}
