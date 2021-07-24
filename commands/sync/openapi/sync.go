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
    gf sync TYPE projectId [OPTION]

TYPE
    model      sync model files
	server	   sync api, service and define files
    client     sync client files
    all        sync server and client

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

	switch syncType {

	case "client":
		doSyncClient(openApi.Tags, &(openApi.Paths))

	case "server":
		refs = make(map[string]DefineInfo)
		getRefsBySchemas(&refs, &(openApi.Components.Schemas))
		doSyncApi(openApi.Tags, &(openApi.Paths))

	case "all":
		refs = make(map[string]DefineInfo)
		getRefsBySchemas(&refs, &(openApi.Components.Schemas))
		doSyncClient(openApi.Tags, &(openApi.Paths))
		doSyncApi(openApi.Tags, &(openApi.Paths))
	}
}
