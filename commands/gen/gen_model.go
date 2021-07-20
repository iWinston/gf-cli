package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenModel implements the "gen model" command.
func doGenModel() {
	var (
		args = getArgs()
		name = args["name"]
	)

	genFile(modelTemplate, "./app/model", name+"_model.go", getReplaceMap(args), "", utils.Header)

	mlog.Print("gen model done!")
}
