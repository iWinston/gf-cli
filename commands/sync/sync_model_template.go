package sync

var modelTemplate = `package model
import "github.com/iWinston/qk-library/frame/q"

type {{.Data.Name}} struct {
	q.Base
{{range $f := .Data.FieldInfos}}
	{{$f.Name}} {{$f.Type}} {{$.Backtick}}{{$f.Tag}}{{$.Backtick}}
{{end}}}
`
