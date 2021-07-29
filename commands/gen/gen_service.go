package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenService implements the "gen service" command.
func doGenService(args map[string]string) {
	var (
		name       = args["name"]
		systemName = args["systemName"]
		fileName   = name + ".service.go"
	)

	genFile(serviceTemplateMap["default"], "./app/system/"+systemName+"/service/internal", fileName, getReplaceMap(args), "force", utils.Header)
	genFile(serviceTemplateMap["index"], "./app/system/"+systemName+"/service", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen service done!")
}
