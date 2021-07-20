package sync

var modelTemplate = `package model

type {{.Data.StructName}} struct {
	Base
{{range $f := .Data.FieldInfos}}
	{{$f.Name}} {{$f.Type}} {{$.Backtick}}{{$f.Tag}}{{$.Backtick}}
{{end}}}
`
