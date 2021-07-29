package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenDefine(args map[string]string) {
	var (
		name       = args["name"]
		systemName = args["systemName"]
		fileName   = name + ".define.go"
	)

	genFile(defineTemplate, "./app/system/"+systemName+"/define", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen define done!")
}
