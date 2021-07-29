package gen

import (
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

func genFile(template string, folder string, fileName string, args g.MapStrStr, mode string, header string) {
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)

	if mode == "force" || gcmd.ContainsOpt("f") || !gfile.Exists(path) {
		indexContent := gstr.ReplaceByMap(header+template, args)
		if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			mlog.Print("generated:", path)
		}
	}
}

func getReplaceMap(args g.MapStrStr) g.MapStrStr {
	return g.MapStrStr{
		"{Backtick}":        "`",
		"{Name}":            args["name"],
		"{CamelName}":       strings.Title(args["name"]),
		"{SnakeName}":       utils.SnakeString(args["name"]),
		"{Description}":     args["description"],
		"{SystemName}":      args["systemName"],
		"{CamelSystemName}": strings.Title(args["systemName"]),
		"{SnakeSystemName}": utils.SnakeString(args["systemName"]),
		"{ModelName}":       strings.Title(args["systemName"]) + strings.Title(args["name"]),
		"{FuncName}":        args["systemName"],
	}
}

func getArgs() (args g.MapStrStr) {
	systemName := gcmd.GetArg(3)
	if systemName == "" {
		mlog.Fatalf("Command arguments are not enough")
	}

	name := gcmd.GetArg(4)
	if name == "" {
		mlog.Fatalf("Command arguments are not enough, The name argument is needed")
	}

	description := gcmd.GetArg(5)
	if description == "" {
		mlog.Fatalf(`Command arguments are not enough, The description argument is needed`)
	}

	parser, err := gcmd.Parse(g.MapStrBool{
		"f,funcName": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	args = parser.GetOptAll()

	args["name"] = name
	args["description"] = description
	args["systemName"] = systemName
	return
}
