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
{{- if eq $method "post" "patch" "delete"}}
// @summary 【{{Title $method}}】{{if $.CamelPrefix}}{{$.CamelPrefix}}/{{end}}{{$.Description}}
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelName}}{{Title $method}}Param true "{{- if eq $method "post"}}"新增{{$.Description}}{{- else}}{{end}}{{- if eq $method "patch"}} "修改{{$.Description}}{{- else}}{{end}}{{- if eq $method "delete"}} "删除{{$.Description}}{{- else}}{{end}}"
// @router  {{- if $.Prefix}}/{{$.Prefix}}{{- end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [Post]
// @success 200 {object} q.JsonResponse "执行结果"
func (a *{{- if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{- else}}{{$.CamelName}}{{- end}}Api) {{Title $method}}(r *ghttp.Request) {
	var (
		ctx   = qservice.ReqContext.Get(r.Context())
		param = &define.{{$.CamelName}}{{Title $method}}Param{}
		tx    = ctx.TX.Model(&model.{{Title $.SystemName}}{{$.CamelName}}{}).WithContext(r.Context())
	)

	ctx.SetActionHistoryTypeAndDesc("{{$.Description}}管理", {{- if eq $method "post"}} "新增{{$.Description}}{{- else}}{{end}}{{- if eq $method "patch"}} "修改{{$.Description}}{{- else}}{{end}}{{- if eq $method "delete"}}"删除{{$.Description}}{{- else}}{{end}}")
	q.AssignParamFormReq(r, param)
	err := q.{{Title $method}}(tx, param)
	q.Response(r, err)
}

{{- else if eq $method "get"}}
// @summary 【{{$method}}】{{if $.CamelPrefix}}{{$.CamelPrefix}}/{{end}}{{$.Description}}
// @description 获取{{$.Description}}
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelName}}{{$method}}Param true "新增{{$.Description}}"
// @router  {{- if $.Prefix}}/{{$.Prefix}}{{- end}}/{{$.SystemName}}/{{$.Name}}/{{$method}} [{{Title $method}}]
// @success 200 {object} q.JsonResponseWithData{data=define.{{$.CamelName}}{{$method}}Res} "执行结果"
func (a *{{- if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{- else}}{{$.CamelName}}{{- end}}Api) {{Title $method}}(r *ghttp.Request) {
	var (
		ctx   = qservice.ReqContext.{{Title $method}}(r.Context())
		param = &define.{{$.CamelName}}{{Title $method}}Param{}
		res   = &define.{{$.CamelName}}{{Title $method}}Res{}
		tx    = ctx.DB.Model(&model.{{Title $.SystemName }}{{$.CamelName}}{}).WithContext(r.Context())
	)

	q.AssignParamFormReq(r, param)
	err := q.{{Title $method}}(tx, param, res)
	q.ResponseWithData(r, err, res)
}

{{- else if eq $method "list"}}
// @summary 【List】{{$.Description}}
// @description {{$.Description}}列表
// @tags    {{$.SystemName}}/{{$.Description}}管理
// @produce  json
// @security ApiKeyAuth
// @param   entity  body define.{{$.CamelName}}ListParam true "{{$.Description}}列表"
// @router  {{- if $.Prefix}}/{{$.Prefix}}{{- end}}/{{$.SystemName}}/{{$.Name}}/list [Post]
// @success 200 {object} q.JsonResponseWithMeta{data=[]define.{{$.CamelName}}ListRes} "执行结果"
func (a *{{- if $.Prefix}}{{$.Prefix}}{{$.CamelName}}{{- else}}{{$.CamelName}}{{- end}}Api) List(r *ghttp.Request) {
	var (
		ctx         = qservice.ReqContext.{{Title $method}}(r.Context())
		param       = &define.{{$.CamelName}}ListParam{}
		res         = &define.{{$.CamelName}}ListRes{}
		tx          = ctx.DB.Model(&model.{{Title $.SystemName }}{{$.CamelName}}{}).WithContext(r.Context())
		total int64 = 0
	)

	q.AssignParamFormReq(r, param)
	err := q.List(tx, param, res, &total)
	q.ResponseWithMeta(r, err, res, total)
}
{{else}}
{{- end}}
{{- end}}`

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
