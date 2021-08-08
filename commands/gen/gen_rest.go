package gen

import (
	"github.com/iWinston/gf-cli/commands/sync"
)

type ApiJson struct {
	Api map[string]ApiInfos
}

type ApiInfos []ApiInfo

type ApiInfo struct {
	Prefix      string
	Name        string
	Description string
	Methods     []string
}

func doGenRest(filePath string) {
	apiJson := &ApiJson{}
	jsonParse := sync.NewJsonStruct()
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	jsonParse.Load(filePath, apiJson)
	for k, apis := range apiJson.Api {
		for _, api := range apis {
			apiInfo := map[string]string{
				"systemName":  k,
				"name":        api.Name,
				"description": api.Description,
				"prefix":      api.Prefix,
			}
			args := getReplaceMap(apiInfo, api.Methods)
			doGenModel(args)
			doGenRouter(args)
			doGenDefine(args)
			doGenService(args)
			doGenApi(args)
		}
	}
}
