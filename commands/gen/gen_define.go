package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenDefine implements the "gen define" command.
func doGenDefine() {
	var (
		args       = getArgs()
		name       = args["name"]
		systemName = args["systemName"]
	)

	genFile(defineTemplate, "./app/system/"+systemName+"/define", name+"_define.go", getReplaceMap(args), "", utils.Header)

	mlog.Print("gen define done!")
}
