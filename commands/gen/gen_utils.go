package gen

import (
	"strings"
	"text/template"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

func genFile(temp string, folder string, fileName string, args Args, mode string, header string) {
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)

	if mode == "force" || gcmd.ContainsOpt("f") || !gfile.Exists(path) {
		// indexContent := gstr.ReplaceByMap(header+template, args)
		// if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
		// 	mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		// } else {
		// 	mlog.Print("generated:", path)
		// }

		t, err := template.New(fileName).Parse(header + temp)
		if err != nil {
			mlog.Fatalf("template to '%s' failed: %v", path, err)
		}
		f, _ := gfile.Create(path)
		if err := t.Execute(f, args); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			mlog.Print("generated:", path)
		}

	}
}

type Args struct {
	Backtick        string
	Name            string
	CamelName       string
	SnakeName       string
	Description     string
	SystemName      string
	CamelSystemName string
	SnakeSystemName string
	ModelName       string
	FuncName        string
	Prefix          string
	CamelPrefix     string
	Methods         []string
}

func getReplaceMap(args g.MapStrStr, methods []string) Args {
	if len(methods) == 0 {
		methods = []string{"post", "get", "patch", "delete", "list"}
	}
	return Args{
		Backtick:        "`",
		Name:            args["name"],
		CamelName:       strings.Title(args["name"]),
		SnakeName:       utils.SnakeString(args["name"]),
		Description:     args["description"],
		SystemName:      args["systemName"],
		CamelSystemName: strings.Title(args["systemName"]),
		SnakeSystemName: utils.SnakeString(args["systemName"]),
		ModelName:       strings.Title(args["systemName"]) + strings.Title(args["name"]),
		FuncName:        args["systemName"],
		Prefix:          args["prefix"],
		CamelPrefix:     strings.Title(args["prefix"]),
		Methods:         methods,
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
