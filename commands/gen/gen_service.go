package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenService implements the "gen service" command.
func doGenService() {
	var (
		args       = getArgs()
		name       = args["name"]
		systemName = args["systemName"]
	)

	genFile(serviceTemplateMap["default"], "./app/system/"+systemName+"/service/internal", name+"_service.go", getReplaceMap(args), "force", utils.Header)
	genFile(serviceTemplateMap["index"], "./app/system/"+systemName+"/service", name+"_service.go", getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen service done!")
}
