package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenService implements the "gen service" command.
func doGenService() {
	name, description := getNameAndDescription()

	genFileForce(serviceTemplateMap["default"], "./gen/service", name+"_service.go", name, description)
	genFileForce(serviceTemplateMap["core"], "./gen/service", name+".go", name, description)
	genFile(serviceTemplateMap["index"], "./app/system/admin/service", name+"_service.go", name, description)

	mlog.Print("gen service done!")
}
