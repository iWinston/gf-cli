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
	name, description, systemName := getArgs()

	genFileWithoutHeader(apifoxTemplate, "./apifox/", name+".apifox.json", name, description, systemName)

	mlog.Print("gen apifox done!")
}
