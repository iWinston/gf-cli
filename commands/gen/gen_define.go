package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenDefine(args Args) {
	var (
		systemName = args.SystemName
		fileName   = utils.JoinNotEmptyStr([]string{args.Name, "define", "go"}, ".")
	)

	genFile(defineTemplate, "./app/system/"+systemName+"/define", fileName, args, "", utils.OnlyOnceHeader)

	mlog.Print("gen define done!")
}
