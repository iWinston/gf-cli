package sync

var DefineTemplate = `package define
{{range $i := .Data.Imports}}import "{{$i}}"
{{end}}
{{range $d := .Data.DefineInfos}}
type {{$d.Name}} struct {
{{range $f := $d.FieldInfos}}	{{$f.Name}} {{$f.Type}} {{$.Backtick}}{{$f.Tag}}{{$.Backtick}}
{{end}}}
{{end}}
`
