package sync

var serviceTemplate = `package service
import "{{.ProjectName}}/app/system/{{.Data.System}}/service/internal"

var {{.Data.StructName}} = &{{.Data.Name}}Service{}

type {{.Data.Name}}Service struct {
	*internal.{{.Data.StructName}}Service
}
`

var serviceInternalTemplate = `package internal

import (
	"{{.ProjectName}}/app/model"
	"{{.ProjectName}}/app/system/{{.Data.System}}/define"

	"github.com/iWinston/qk-library/frame/q"
)

type {{.Data.StructName}}Service struct{}
{{range $api := .Data.ApiInfos}}
// {{$api.FuncName}} {{$api.Name}}
func (s *{{$.Data.StructName}}Service) {{$api.FuncName}}(param *define.{{$api.ParamRef}}) ({{$api.Service.ReturnType}}) {
	{{if eq $api.FuncName "Get"}}
	var (
		res = &define.{{$api.ResRef}}{}
		tx = model.DB.Model(&model.{{$.Data.StructName}}{})
		err = q.Get(tx, param, res)
	)
	{{else if eq $api.FuncName "Post"}}
	var (
		m = &model.{{$.Data.StructName}}{}
		err = q.Post(model.DB, m, param)
	)
	{{else if eq $api.FuncName "Patch"}}
	var (
		m = &model.{{$.Data.StructName}}{}
		tx = model.DB.Model(m)
		err = q.Patch(tx, m, param)
	)
	{{else if eq $api.FuncName "Delete"}}
	var (
		m = &model.{{$.Data.StructName}}{}
		tx = model.DB.Model(m)
		err = q.Delete(tx, m, param)
	)
	{{else if eq $api.FuncName "List"}}
	var (
		total int64
		res   = &[]define.{{$api.ResRef}}{}
		tx    = model.DB.Model(&model.{{$.Data.StructName}}{})
		err   = q.List(tx, param, res, &total)
	)
	{{else if eq $api.RespMode "resp"}}
		var err error
	{{else if eq $api.RespMode "data"}}
	var (
		res = &define.{{$api.ResRef}}{}
		err error
	)
	{{else if eq $api.RespMode "meta"}}
	var (
		total int64
		res   = &[]define.{{$api.ResRef}}{}
		err   error
	)
	{{end}}
	return {{$api.Service.Return}}
}
{{end}}
`
