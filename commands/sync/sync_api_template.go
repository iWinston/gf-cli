package sync

var apiTemplate = `package api
import "{{.ProjectName}}/app/system/{{.Data.System}}/api/internal"

var {{.Data.StructName}} = &{{.Data.Name}}Api{}

type {{.Data.Name}}Api struct {
	*internal.{{.Data.StructName}}Api
}`

var apiInternalTemplate = `package internal
import (
	"{{.ProjectName}}/app/system/{{.Data.System}}/define"
	"{{.ProjectName}}/app/system/{{.Data.System}}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
)

var {{.Data.StructName}} = &{{.Data.Name}}Api{}

type {{.Data.Name}}Api struct{}
{{range $api := .Data.ApiInfos}}
// {{$api.FuncName}} {{$api.Name}}
func (a *{{$.Data.Name}}Api) {{$api.FuncName}}(r *ghttp.Request) {
	param := &define.{{$api.ParamRef}}{}
	q.AssignParamFormReq(r, param)
	{{$api.Service.Return}} := service.{{$.Data.StructName}}.{{$api.FuncName}}(param)
	{{$api.RespFunc}}
}
{{end}}
`
