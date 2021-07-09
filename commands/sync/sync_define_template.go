package sync

var defineTemplate = `package define
import "github.com/iWinston/qk-library/frame/q"

{{range $d := .Data}}
type {{$d.Name}} struct {
{{range $f := $d.FieldInfos}}
	{{$f.Name}} {{$f.Type}} {{$.Backtick}}{{$f.Tag}}{{$.Backtick}}
{{end}}}
{{end}}
`
