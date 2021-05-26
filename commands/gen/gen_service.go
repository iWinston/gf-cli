package gen

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenService implements the "gen service" command.
func doGenService() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"n, name":       true,
		"d,description": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	name := parser.GetOpt("name")
	description := parser.GetOpt("description")
	genFileForce(serviceTemplateMap["default"], "./gen/service", name+"_service.go", name, description)
	genFileForce(serviceTemplateMap["core"], "./gen/service", name+".go", name, description)
	genFile(serviceTemplateMap["index"], "./app/system/admin/service", name+"_service.go", name, description)

	mlog.Print("gen service done!")
}
