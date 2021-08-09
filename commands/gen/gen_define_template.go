package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"
import "github.com/iWinston/qk-library/frame/qfield"

{{- range $method := .Methods}}
{{- if eq $method "post"}}
// 新增参数
type {{$.CamelPrefix}}{{$.CamelName}}PostParam struct {
}
{{- else if eq $method "get"}}
// 详情参数
type {{$.CamelPrefix}}{{$.CamelName}}GetParam struct {
	qfield.Id
}
// 详情
type {{$.CamelPrefix}}{{$.CamelName}}GetRes struct {
	Id *uint
}
{{- else if eq $method "patch"}}
// 修改参数
type {{$.CamelPrefix}}{{$.CamelName}}PatchParam struct {
	qfield.Id
}
{{- else if eq $method "delete"}}
// 删除参数
type {{$.CamelPrefix}}{{$.CamelName}}DeleteParam struct {
	qfield.Id
}
{{- else if eq $method "list"}}
// 列表参数
type {{$.CamelPrefix}}{{$.CamelName}}ListParam struct {
	q.Page
}
// 列表
type {{$.CamelPrefix}}{{$.CamelName}}ListRes struct {
	Id *uint
}
{{- else}}
type {{$.CamelPrefix}}{{$.CamelName}}{{Title $method}}Param struct {
	
}
{{- end}}
{{- end}}
`
