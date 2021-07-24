package openapi

import "strings"

func doSyncClient(tags []Tag, paths *map[string]map[string]Api) {
	apiFileInfo := make(map[string]*ApiFileInfo)
	defineFileInfo := make(map[string]*DefineFileInfo)

	for _, tag := range tags {
		if len(strings.Split(tag.Name, "/")) != 1 {
			apiFileInfo[tag.Name] = &ApiFileInfo{}
			getApiFileInfo(apiFileInfo[tag.Name], tag.Name)

			defineFileInfo[tag.Name] = &DefineFileInfo{
				System: apiFileInfo[tag.Name].System,
				Name:   apiFileInfo[tag.Name].Name,
			}
		}
	}
	for path, apis := range *paths {
		for method, api := range apis {
			tagName := api.Tags[0]
			apiInfo := getApiInfo(apiFileInfo[tagName].StructName, api, path, method)
			apiFileInfo[tagName].ApiInfos = append(apiFileInfo[tagName].ApiInfos, apiInfo)

			if apiInfo.ParamRef != "" {
				paramDefineInfo := refs[apiInfo.ParamRef]
				defineFileInfo[tagName].DefineInfos = append(defineFileInfo[tagName].DefineInfos, paramDefineInfo)
			}
			if apiInfo.ResRef != "" {
				resDefineInfo := refs[apiInfo.ResRef]
				defineFileInfo[tagName].DefineInfos = append(defineFileInfo[tagName].DefineInfos, resDefineInfo)
			}
		}
	}
	syncApiFiles(&apiFileInfo)
	syncDefineFile(&defineFileInfo)
}
