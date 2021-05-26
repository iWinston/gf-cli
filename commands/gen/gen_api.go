package gen

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenApi implements the "gen api" command.
func doGenApi() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"n, name":       true,
		"d,description": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	name := parser.GetOpt("name")
	description := parser.GetOpt("description")
	genFileForce(apiTemplateMap["default"], "./gen/api", name+"_api.go", name, description)
	genFileForce(apiTemplateMap["core"], "./gen/api", name+".go", name, description)
	genFile(apiTemplateMap["index"], "./app/system/admin/api", name+"_api.go", name, description)

	mlog.Print("gen api done!")
}
