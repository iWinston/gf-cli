package gen

var apiInternalTemplate = `
package internal

import (
	"server/app/system/{{$.SystemName}}/define"
	"server/app/system/{{$.SystemName}}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
	"github.com/iWinston/qk-library/frame/qservice"
)

// {{$.Description}}API
type {{$.CamelPrefix}}{{$.CamelName}}Api struct{}

{{range $method := .Methods}}
// @summary 【Post】{{$.Description}}
// @description 新增{{$.Description}}
// @tags    /{{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelPrefix}}{{$.CamelName}}PostParam true ""
// @router  /{{$.SystemName}}{{if $.Prefix}}/{{$.Prefix}}{{end}}/{{$.Name}}/{{$method}} [{{$method}}]
{{- if eq $method "post"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Post(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "新增{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}PostParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Post(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "get"}}
// @success 200 {object} q.JsonResponseWithData{data=define.{{$.CamelPrefix}}{{$.CamelName}}GetRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Get(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}GetParam{}
	q.AssignParamFormReq(r, param)
	res, err := service.{{$.CamelName}}.Get(ctx, param)
	q.ResponseWithData(r, err, res)
}
{{- else if eq $method "patch"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Patch(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "修改{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Patch(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "delete"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Delete(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "删除{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Delete(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "list"}}
// @success 200 {object} q.JsonResponseWithMeta{data=[]define.{{$.CamelName}}ListRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) List(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}ListParam{}
	q.AssignParamFormReq(r, param)
	res, total, err := service.{{$.CamelName}}.List(ctx, param)
	q.ResponseWithMeta(r, err, res, total)
}
{{end}}
{{end}}`

var apiIndexTemplate = `package api

import "server/app/system/{{$.SystemName}}/api/internal"

var {{$.CamelName}} = &{{$.Name}}Api{}

type {{$.Name}}Api struct {
	*internal.{{$.CamelName}}Api
}`

var apiTemplateMap = map[string]string{
	"index":    apiIndexTemplate,
	"internal": apiInternalTemplate,
}
