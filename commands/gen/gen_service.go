package gen

import (
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

// doGenService implements the "gen service" command.
func doGenService(args Args) {
	var (
		name       = args.Name
		systemName = args.SystemName
		fileName   = name + ".service.go"
	)

	// genFile(serviceTemplateMap["default"], "./app/system/"+systemName+"/service/internal", fileName, args, "force", utils.Header)
	genFile(serviceTemplateMap["index"], "./app/system/"+systemName+"/service", fileName, args, "", utils.OnlyOnceHeader)

	mlog.Print("gen service done!")
}
