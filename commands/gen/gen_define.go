package gen

import (
	"strings"

	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenDefine() {
	var (
		args       = getArgs()
		name       = args["name"]
		systemName = args["systemName"]
		fileName   = systemName + strings.Title(name) + ".define.go"
	)

	genFile(defineTemplate, "./app/system/"+systemName+"/define", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen define done!")
}
