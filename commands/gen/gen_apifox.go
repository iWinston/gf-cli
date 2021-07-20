package gen

import (
	"github.com/gogf/gf/text/gstr"
	"github.com/iWinston/gf-cli/library/mlog"
)

func HelpApifox() {
	mlog.Print(gstr.TrimLeft(`
USAGE 
    gf gen apifox 

`))
}

// doGenApifox implements the "gen service" command.
func doGenApifox() {
	var (
		args       = getArgs()
		name       = args["name"]
		systemName = args["systemName"]
	)

	if args["funcName"] == "" {
		genFile(apifoxTemplate, "./apifox/"+systemName, name+".apifox.json", getReplaceMap(args), "force", "")
	} else {
		genFile(apifoxTemplate, "./apifox/"+systemName, name+"."+args["funcName"]+".apifox.json", getReplaceMap(args), "force", "")
	}

	mlog.Print("gen apifox done!")
}
