package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenModel implements the "gen model" command.
func doGenModel(args Args) {
	var (
		fileName = args.SnakeSystemName + args.SnakeName + ".model.go"
	)

	genFile(modelTemplate, "./app/model", fileName, args, "", utils.OnlyOnceHeader)

	mlog.Print("gen model done!")
}
