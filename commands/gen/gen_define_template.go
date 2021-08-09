package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"
import "github.com/iWinston/qk-library/frame/qfield"

// 详情参数
type {{$.CamelName}}GetParam struct {
	qfield.Id
}
// 详情
type {{$.CamelName}}GetRes struct {
	Id *uint
}

// 新增参数
type {{$.CamelName}}PostParam struct {
}

// 修改参数
type {{$.CamelName}}PatchParam struct {
	qfield.Id
}

// 删除参数
type {{$.CamelName}}DeleteParam struct {
	qfield.Id
}

// 列表参数
type {{$.CamelName}}ListParam struct {
	q.Page
}
// 列表
type {{$.CamelName}}ListRes struct {
	Id *uint
}

{{range $method := .Methods}}
{{- if eq $method "get" "post" "patch" "delete" "list"}}
{{- else}}
type {{$.CamelName}}{{Title $method}}Param struct {
	
}
{{- end}}
{{- end}}
`
