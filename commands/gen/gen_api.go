package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenApi implements the "gen api" command.
func doGenApi() {
	var (
		args       = getArgs()
		name       = args["name"]
		systemName = args["systemName"]
	)

	genFile(apiTemplateMap["index"], "./app/system/"+systemName+"/api", name+"_api.go", getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen api done!")
}
