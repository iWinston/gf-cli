package gen

import (
	"strings"

	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenModel implements the "gen model" command.
func doGenModel(args map[string]string) {
	var (
		name       = args["name"]
		systemName = args["systemName"]
		fileName   = systemName + strings.Title(name) + ".model.go"
	)

	genFile(modelTemplate, "./app/model", fileName, getReplaceMap(args), "", utils.OnlyOnceHeader)

	mlog.Print("gen model done!")
}
