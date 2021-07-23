package openapi

import (
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/text/gstr"
	"github.com/iWinston/gf-cli/commands/sync"
	"github.com/iWinston/gf-cli/library/mlog"
)

func Help() {
	switch gcmd.GetArg(2) {
	case "model":
	case "api":
	default:
		mlog.Print(gstr.TrimLeft(`
USAGE 
    gf sync TYPE [OPTION]

TYPE
    model      sync model files
    api        sync api, service and define files
    all        sync model, api, service and define files

DESCRIPTION
    The "sync" command is designed for multiple generating purposes. 
    Please use "gf gen model -h" or "gf gen api -h" for specified type help.
`))
	}
}

var refs map[string]DefineInfo
var projectName string

func Run() {
	syncType := gcmd.GetArg(2)
	if syncType == "" {
		mlog.Print("generating type cannot be empty")
		return
	}
	openApi := &OpenApi{}
	sync.Load(openApi)
	projectName = openApi.Info.Description
	refs = make(map[string]DefineInfo)
	getRefsBySchemas(&refs, &(openApi.Components.Schemas))
	switch syncType {

	// case "model":
	// doSyncModel(&(apifox.SchemaCollection))

	//case "router":
	// doSyncRouter(&(apifox.ApiCollection))

	case "api":
		doSyncApi(openApi.Tags, &(openApi.Paths))

	case "all":
		// doSyncModel(&(apifox.SchemaCollection))
		doSyncApi(openApi.Tags, &(openApi.Paths))
	}
}
