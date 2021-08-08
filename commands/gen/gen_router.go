package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenRouter(args Args) {
	var (
		systemName = args.SystemName
		fileName   = systemName + ".go"
	)

	genFile(routerTemplate, "./app/system/"+systemName+"", fileName, args, "", utils.OnlyOnceHeader)

	mlog.Print("gen define done!")
}
