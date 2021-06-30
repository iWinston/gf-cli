package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenModel implements the "gen model" command.
func doGenModel() {
	name, description, systemName := getArgs()

	genFile(modelTemplate, "./app/model", name+"_model.go", name, description, systemName)

	mlog.Print("gen model done!")
}
