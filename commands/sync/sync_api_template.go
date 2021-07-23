package sync

var ApiTemplate = `package api
import "{{.ProjectName}}/app/system/{{.Data.System}}/api/internal"

var {{.Data.StructName}} = &{{.Data.Name}}Api{}

type {{.Data.Name}}Api struct {
	*internal.{{.Data.StructName}}Api
}`

var ApiInternalTemplate = `package internal
import (
	"{{.ProjectName}}/app/system/{{.Data.System}}/define"
	"{{.ProjectName}}/app/system/{{.Data.System}}/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/q"
)

type {{.Data.StructName}}Api struct{}
{{range $api := .Data.ApiInfos}}
// {{$api.FuncName}} {{$api.Summary}}
func (a *{{$.Data.StructName}}Api) {{$api.FuncName}}(r *ghttp.Request) {
	param := &define.{{$api.ParamRefName}}{}
	q.AssignParamFormReq(r, param)
	{{$api.Service.Return}} := service.{{$.Data.StructName}}.{{$api.FuncName}}(param)
	{{$api.RespFunc}}
}
{{end}}
`
