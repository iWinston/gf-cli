package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenDefine implements the "gen define" command.
func doGenDefine() {
	name, description, systemName := getArgs()

	genFile(defineTemplate, "./app/system/"+systemName+"/define", name+"_define.go", name, description, systemName)

	mlog.Print("gen define done!")
}
