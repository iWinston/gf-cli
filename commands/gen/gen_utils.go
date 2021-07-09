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

func genFile(template string, folder string, fileName string, name string, description string, systemName string) {
	fileName = utils.SnakeString(fileName)
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)

	if gcmd.ContainsOpt("f") || !gfile.Exists(path) {
		indexContent := gstr.ReplaceByMap(utils.OnlyOnceHeader+template, g.MapStrStr{
			"{TplName}":        name,
			"{TplUpperName}":   strings.ToUpper(name[:1]) + name[1:],
			"{TplDescription}": description,
			"{TplSystemName}":  systemName,
		})
		if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			mlog.Print("generated:", path)
		}
	}
}

func genFileForce(template string, folder string, fileName string, name string, description string, systemName string) {
	fileName = utils.SnakeString(fileName)
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)
	indexContent := gstr.ReplaceByMap(utils.Header+template, g.MapStrStr{
		"{TplName}":        name,
		"{TplUpperName}":   strings.ToUpper(name[:1]) + name[1:],
		"{TplDescription}": description,
		"{TplSystemName}":  systemName,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		mlog.Print("generated:", path)
	}
}

func getArgs() (name string, description string, systemName string) {
	name = gcmd.GetArg(3)
	if name == "" {
		mlog.Fatalf("Command arguments are not enough")
	}
	description = gcmd.GetArg(4)

	if description == "" {
		mlog.Fatalf(`Command arguments are not enough, The description argument is needed`)
	}

	parser, err := gcmd.Parse(g.MapStrBool{
		"s,systemName": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	systemName = parser.GetOpt("systemName")
	if systemName == "" {
		systemName = "admin"
	}
	return
}
