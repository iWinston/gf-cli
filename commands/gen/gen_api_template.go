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

{{- range $method := .Methods}}
{{- if eq $method "get" "post" "patch" "delete" "list"}}
// @summary 【{{Title $method }}】{{if $.CamelPrefix}}{{$.CamelPrefix}}/{{end}}{{$.Description}}
// @description 新增{{$.Description}}
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelName}}{{Title $method}}Param true "新增{{$.Description}}"
// @router  {{if $.Prefix}}/{{$.Prefix}}{{end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [Post]
{{- if eq $method "post"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Post(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "新增{{$.Description}}")
	param := &define.{{$.CamelName}}PostParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Post(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "get"}}
// @success 200 {object} q.JsonResponseWithData{data=define.{{$.CamelName}}GetRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Get(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{{$.CamelName}}GetParam{}
	q.AssignParamFormReq(r, param)
	res, err := service.{{$.CamelName}}.Get(ctx, param)
	q.ResponseWithData(r, err, res)
}
{{- else if eq $method "patch"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Patch(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "修改{{$.Description}}")
	param := &define.{{$.CamelName}}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Patch(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "delete"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Delete(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "删除{{$.Description}}")
	param := &define.{{$.CamelName}}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.Delete(ctx, param)
	q.Response(r, err)
}
{{- else if eq $method "list"}}
// @success 200 {object} q.JsonResponseWithMeta{data=[]define.{{$.CamelName}}ListRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) List(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	param := &define.{{$.CamelName}}ListParam{}
	q.AssignParamFormReq(r, param)
	res, total, err := service.{{$.CamelName}}.List(ctx, param)
	q.ResponseWithMeta(r, err, res, total)
}
{{end}}
{{end}}
{{end}}`

var apiIndexTemplate = `package api

import "server/app/system/{{$.SystemName}}/api/internal"

var {{$.CamelPrefix}}{{$.CamelName}} = &{{if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{else}}{{$.Name}}{{end}}Api{}

type {{if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{else}}{{$.Name}}{{end}}Api struct {
	*internal.{{if $.Prefix}}{{$.CamelPrefix}}{{$.CamelName}}{{else}}{{$.CamelName}}{{end}}Api
}
{{- range $method := .Methods}}
{{- if eq $method "get" "post" "patch" "delete" "list"}}
{{- else}}
// @summary 【{{Title $method}}】{{if $.CamelPrefix}}{{$.CamelPrefix}}/{{end}}{{$.Description}}
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelName}}{{Title $method }}Param true ""
// @router  {{- if $.Prefix}}/{{$.Prefix}}{{- end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [Post]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{- if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{- else}}{{$.Name}}{{- end}}Api) {{Title $method}}(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "新增{{$.Description}}")
	param := &define.{{$.CamelName}}{{Title $method}}Param{}
	q.AssignParamFormReq(r, param)
	err := service.{{$.CamelName}}.{{Title $method}}(ctx, param)
	q.Response(r, err)
}
{{- end}}
{{- end}}
`

var apiTemplateMap = map[string]string{
	"index":    apiIndexTemplate,
	"internal": apiInternalTemplate,
}
