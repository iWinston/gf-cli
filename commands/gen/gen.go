package gen

import (
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/text/gstr"
	"github.com/iWinston/gf-cli/library/mlog"
)

func Help() {
	switch gcmd.GetArg(2) {

	case "apifox":
		HelpApifox()

	default:
		mlog.Print(gstr.TrimLeft(`
USAGE 
    gf gen TYPE [OPTION]

TYPE
    dao        generate dao and model files.
    model      generate model files, note that these generated model files are different from model files 
               of command "gf gen dao".
    pb         parse proto files and generate protobuf go files.
    pbentity   generate entity message files in protobuf3 format.
    rest       generate restful files.

DESCRIPTION
    The "gen" command is designed for multiple generating purposes. 
    It's currently supporting generating go files for ORM models, protobuf and protobuf entity files.
    Please use "gf gen dao -h" or "gf gen model -h" for specified type help.
`))
	}
}

func Run() {
	genType := gcmd.GetArg(2)
	if genType == "" {
		mlog.Print("generating type cannot be empty")
		return
	}
	if genType == "file" {
		filePath := gcmd.GetArg(3)
		doGenRest(filePath)
	}
	// else {
	// 	args := getArgs()
	// 	switch genType {

	// 	case "model":
	// 		doGenModel(args)

	// 	case "apifox":
	// 		doGenApifox()

	// 	case "api":
	// 		doGenApi(args)

	// 	case "service":
	// 		doGenService(args)

	// 	case "define":
	// 		doGenDefine(args)

	// 	case "router":
	// 		doGenRouter(args)

	// 	case "rest":
	// 		doGenModel(args)
	// 		doGenRouter(args)
	// 		doGenDefine(args)
	// 		doGenService(args)
	// 		doGenApi(args)
	// 	}
	// }
}
