package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenApi implements the "gen api" command.
func doGenApi(args Args) {
	var (
		name       = args.Name
		systemName = args.SystemName
		fileName   = utils.JoinNotEmptyStr([]string{name, args.Prefix, "api", "go"}, ".")
	)

	genFile(apiTemplateMap["index"], "./app/system/"+systemName+"/api", fileName, args, "", utils.OnlyOnceHeader)
	genFile(apiTemplateMap["internal"], "./app/system/"+systemName+"/api/internal", fileName, args, "force", utils.Header)
	mlog.Print("gen api done!")
}
