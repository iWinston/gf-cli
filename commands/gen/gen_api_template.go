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
// @param   entity  body define.{TplUpperName}CreateDto true "新增{TplDescription}"
// @router  /{TplName} [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplUpperName}Api) Post(r *ghttp.Request) {
	req := a.{TplUpperName}.CreateDto(r)
	a.{TplUpperName}.Post(r, req)
}

// @summary {TplDescription}列表接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity query define.{TplUpperName}GetDto true "分页"
// @router  /{TplName} [GET]
// @success 200 {object} response.JsonResponseWithTotal "执行结果"
func (a *{TplUpperName}Api) Get(r *ghttp.Request) {
	req := a.{TplUpperName}.GetDto(r)
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
	id := a.{TplUpperName}.GetOneDto(r)
	a.{TplUpperName}.GetOne(r, id)
}

// @summary {TplDescription}修改接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @param entity body define.{TplUpperName}PatchOneDto true "修改内容"
// @router  /{TplName}/{id} [Patch]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplUpperName}Api) PatchOne(r *ghttp.Request) {
	id, req := a.{TplUpperName}.PatchOneDto(r)
	a.{TplUpperName}.PatchOne(r, id, req)
}

// @summary {TplDescription}删除接口
// @tags    {TplDescription}管理
// @produce  json
// @security ApiKeyAuth
// @router  /{TplName}/{id} [Delete]
// @Param id path int true "{TplUpperName} ID"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *{TplUpperName}Api) DeleteOne(r *ghttp.Request) {
	id := a.{TplUpperName}.DeleteOneDto(r)
	a.{TplUpperName}.DeleteOne(r, id)
}
`

var apiCoreTemplate = `
package api

import (
	"server/app/system/admin/define"
	"server/app/system/admin/service"
	"server/library/response"

	"github.com/gogf/gf/net/ghttp"
)

type {TplUpperName} struct{}

func (a *{TplUpperName}) Post(r *ghttp.Request, req *define.{TplUpperName}CreateDto) {
	if err := service.{TplUpperName}.Create(req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *{TplUpperName}) CreateDto(r *ghttp.Request) (req *define.{TplUpperName}CreateDto) {
	if err := r.ParseForm(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return
}

func (a *{TplUpperName}) Get(r *ghttp.Request, req *define.{TplUpperName}GetDto) {
	if list, total, err := service.{TplUpperName}.Get(req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list, total)
	}
}

func (a *{TplUpperName}) GetDto(r *ghttp.Request) (req *define.{TplUpperName}GetDto) {
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
func (a *{TplUpperName}) GetOneDto(r *ghttp.Request) (id uint) {
	if id = r.GetUint("id"); id == 0 {
		response.JsonExit(r, 1, "未获得id参数")
	}
	return
}

func (a *{TplUpperName}) PatchOne(r *ghttp.Request, id uint, req *define.{TplUpperName}PatchOneDto) {
	if err := service.{TplUpperName}.PatchOne(id, req); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *{TplUpperName}) PatchOneDto(r *ghttp.Request) (id uint, req *define.{TplUpperName}PatchOneDto) {
	if id = r.GetUint("id"); id == 0 {
		response.JsonExit(r, 1, "未获得id参数")
	}
	if err := r.ParseForm(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	return
}

func (a *{TplUpperName}) DeleteOne(r *ghttp.Request, id uint) {
	if err := service.{TplUpperName}.DeleteOne(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *{TplUpperName}) DeleteOneDto(r *ghttp.Request) (id uint) {
	if id = r.GetUint("id"); id == 0 {
		response.JsonExit(r, 1, "未获得id参数")
	}
	return
}
`

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
