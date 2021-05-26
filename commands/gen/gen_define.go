package gen

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenDefine implements the "gen define" command.
func doGenDefine() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"n, name":       true,
		"d,description": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	name := parser.GetOpt("name")
	description := parser.GetOpt("description")

	genFile(defineTemplate, "./app/system/admin/define", name+"_define.go", name, description)

	mlog.Print("gen define done!")
}
