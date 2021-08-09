package gen

var apiInternalTemplate = `
package internal

import (
	"server/app/model"
	"server/app/system/{{$.SystemName}}/define"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
	"github.com/iWinston/qk-library/frame/qservice"
)

// {{$.Description}}API
type {{$.CamelPrefix}}{{$.CamelName}}Api struct{}

{{range $method := .Methods}}
{{if eq $method "get" "post" "patch" "delete" "list"}}
// @summary 【{{Title $method }}】{{if $.CamelPrefix}}{{$.CamelPrefix}}/{{end}}{{$.Description}}
// @description 新增{{$.Description}}
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelPrefix}}{{$.CamelName}}{{Title $method}}Param true "新增{{$.Description}}"
// @router  {{if $.Prefix}}/{{$.Prefix}}{{end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [Post]
{{- if eq $method "post"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Post(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "新增{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}PostParam{}
	q.AssignParamFormReq(r, param)
	err := q.Post(ctx.TX.Model(&model.{{$.ModelName}}{}).WithContext(r.Context()), param)
	q.Response(r, err)
}
{{- else if eq $method "get"}}
// @success 200 {object} q.JsonResponseWithData{data=define.{{$.CamelPrefix}}{{$.CamelName}}GetRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Get(r *ghttp.Request) {
	var (
		ctx = qservice.ReqContext.Get(r.Context())
		param = &define.{{$.CamelPrefix}}{{$.CamelName}}GetParam{}
		res = &define.{{$.CamelPrefix}}{{$.CamelName}}GetRes{}
	)
	q.AssignParamFormReq(r, param)
	err := q.Get(ctx.DB.Model(&model.{{$.ModelName}}{}), param, res)
	q.ResponseWithData(r, err, res)
}
{{- else if eq $method "patch"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Patch(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "修改{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}PatchParam{}
	q.AssignParamFormReq(r, param)
	err := q.Patch(ctx.TX.Model(&model.{{$.ModelName}}{}).WithContext(r.Context()), param)
	q.Response(r, err)
}
{{- else if eq $method "delete"}}
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) Delete(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "删除{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}DeleteParam{}
	q.AssignParamFormReq(r, param)
	err := q.Delete(ctx.TX.Model(&model.{{$.ModelName}}{}).WithContext(r.Context()), param)
	q.Response(r, err)
}
{{- else if eq $method "list"}}
// @success 200 {object} q.JsonResponseWithMeta{data=[]define.{{$.CamelName}}ListRes} "执行结果"
func (a *{{$.CamelPrefix}}{{$.CamelName}}Api) List(r *ghttp.Request) {
	var (
		ctx = qservice.ReqContext.Get(r.Context())
		param = &define.{{$.CamelPrefix}}{{$.CamelName}}ListParam{}
		res = &[]define.{{$.CamelPrefix}}{{$.CamelName}}ListRes{}
		total int64
	)
	q.AssignParamFormReq(r, param)
	err := q.List(ctx.TX.Model(&model.{{$.ModelName}}{}), param, res, &total)
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
// @param   entity  body define.{{$.CamelPrefix}}{{$.CamelName}}{{Title $method }}Param true ""
// @router  {{- if $.Prefix}}/{{$.Prefix}}{{- end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [Post]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{- if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{- else}}{{$.Name}}{{- end}}Api) {{Title $method}}(r *ghttp.Request) {
	ctx := qservice.ReqContext.Get(r.Context())
	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", "新增{{$.Description}}")
	param := &define.{{$.CamelPrefix}}{{$.CamelName}}{{Title $method}}Param{}
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
