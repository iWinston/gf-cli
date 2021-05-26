package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenDefine implements the "gen define" command.
func doGenDefine() {
	name, description := getNameAndDescription()

	genFile(defineTemplate, "./app/system/admin/define", name+"_define.go", name, description)

	mlog.Print("gen define done!")
}
