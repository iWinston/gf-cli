package gen

import (
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/iWinston/gf-cli/library/mlog"
)

const header = `
// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================
`

const onlyOnceHeader = `
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================
`

func genFile(template string, folder string, fileName string, name string, description string) {
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)
	if !gfile.Exists(path) {
		indexContent := gstr.ReplaceByMap(onlyOnceHeader+template, g.MapStrStr{
			"{TplName}":        name,
			"{TplUpperName}":   strings.ToUpper(name[:1]) + name[1:],
			"{TplDescription}": description,
		})
		if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			mlog.Print("generated:", path)
		}
	}
}

func genFileForce(template string, folder string, fileName string, name string, description string) {
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)
	indexContent := gstr.ReplaceByMap(header+template, g.MapStrStr{
		"{TplName}":        name,
		"{TplUpperName}":   strings.ToTitle(name),
		"{TplDescription}": description,
	})
	if err := gfile.PutContents(path, strings.TrimSpace(indexContent)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		mlog.Print("generated:", path)
	}

}

func getNameAndDescription() (name string, description string) {
	name = gcmd.GetArg(3)
	if name == "" {
		mlog.Fatalf("Command arguments are not enough")
	}

	parser, err := gcmd.Parse(g.MapStrBool{
		"d,description": true,
	})
	if err != nil {
		mlog.Fatal(err)
	}
	//name := parser.GetOpt("name")
	description = parser.GetOpt("description")

	if description == "" {
		mlog.Fatalf(`The description option is needed, please use "-d" for description`)
	}
	return
}