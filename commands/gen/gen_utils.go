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
	fileName = utils.SnakeString(fileName)
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
		"{TplName}":            args["name"],
		"{TplUpperName}":       strings.Title(args["name"]),
		"{TplDescription}":     args["description"],
		"{TplSystemName}":      args["systemName"],
		"{TplSystemUpperName}": strings.Title(args["systemName"]),
		"{TplFuncName}":        args["systemName"],
	}
}

func getArgs() (args g.MapStrStr) {
	name := gcmd.GetArg(3)
	if name == "" {
		mlog.Fatalf("Command arguments are not enough")
	}

	description := gcmd.GetArg(4)
	if description == "" {
		mlog.Fatalf(`Command arguments are not enough, The description argument is needed`)
	}

	parser, err := gcmd.Parse(g.MapStrBool{
		"s,systemName": true,
		"f,funcName":   true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	args = parser.GetOptAll()
	if args["systemName"] == "" {
		args["systemName"] = "Admin"
	}
	args["name"] = name
	args["description"] = description
	return
}
