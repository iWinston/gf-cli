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
		fileName   = name + ".api.go"
	)

	genFile(apiTemplateMap["index"], "./app/system/"+systemName+"/api", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)
	genFile(apiTemplateMap["internal"], "./app/system/"+systemName+"/api/internal", fileName, getReplaceMap(args), "force", utils.Header)
	mlog.Print("gen api done!")
}
