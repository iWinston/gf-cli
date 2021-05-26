package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
)

// doGenModel implements the "gen model" command.
func doGenModel() {
	name, description := getNameAndDescription()

	genFile(modelTemplate, "./app/model", name+"_model.go", name, description)

	mlog.Print("gen model done!")
}
