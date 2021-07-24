package openapi

import (
	"strings"

	"github.com/iWinston/gf-cli/commands/sync"
	"github.com/iWinston/gf-cli/library/utils"
)

type ApiFileInfo struct {
	System     string
	StructName string
	Name       string
	FileName   string
	ApiInfos   []ApiInfo
}

type ApiInfo struct {
	Summary      string
	FuncName     string
	Param        string
	RespType     string
	Path         string
	Method       string
	Service      ApiService
	RespFunc     string
	RespMode     string
	ParamRef     string
	ParamRefName string
	ResRef       string
	ResRefName   string
}

type ApiService struct {
	Return     string
	ReturnType string
}

func doSyncApi(tags []Tag, paths *map[string]map[string]Api) {
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
			structName := apiFileInfo[tagName].StructName
			apiInfo := getApiInfo(structName, api, path, method)
			apiFileInfo[tagName].ApiInfos = append(apiFileInfo[tagName].ApiInfos, apiInfo)

			if apiInfo.ParamRef != "" && strings.Contains(apiInfo.ParamRef, structName) {
				paramDefineInfo := refs[apiInfo.ParamRef]
				defineFileInfo[tagName].DefineInfos = append(defineFileInfo[tagName].DefineInfos, paramDefineInfo)
			}
			if apiInfo.ResRef != "" && strings.Contains(apiInfo.ResRef, structName) {
				resDefineInfo := refs[apiInfo.ResRef]
				defineFileInfo[tagName].DefineInfos = append(defineFileInfo[tagName].DefineInfos, resDefineInfo)
			}
		}
	}
	syncApiFiles(&apiFileInfo)
	syncDefineFile(&defineFileInfo)
}

func syncApiFiles(fileInfos *map[string]*ApiFileInfo) {
	for _, fileInfo := range *fileInfos {
		sync.SyncFile("app/system/"+fileInfo.System+"/api", fileInfo.FileName+".api.go", sync.ApiTemplate, fileInfo)
		sync.SyncFileForce("app/system/"+fileInfo.System+"/api/internal", fileInfo.FileName+".api.go", sync.ApiInternalTemplate, fileInfo)
		sync.SyncFile("app/system/"+fileInfo.System+"/service", fileInfo.FileName+".service.go", sync.ServiceTemplate, fileInfo)
		sync.SyncFileForce("app/system/"+fileInfo.System+"/service/internal", fileInfo.FileName+".service.go", sync.ServiceInternalTemplate, fileInfo)
	}
}

func getApiFileInfo(info *ApiFileInfo, tagName string) {
	system := strings.Split(tagName, "/")[0]
	info.System = utils.SnakeString(system)
	arr := strings.Split(tagName, "#")
	name := arr[len(arr)-1]
	info.Name = strings.ToLower(name[:1]) + name[1:]
	info.StructName = strings.Title(name)
	info.FileName = utils.SnakeString(name)
}

func getApiInfo(name string, api Api, path string, method string) (apiInfo ApiInfo) {
	apiInfo.Summary = api.Summary
	apiInfo.Path = path
	apiInfo.Method = method
	apiInfo.FuncName = getFuncName(path, method)
	apiInfo.ResRef, apiInfo.ResRefName = getResRef(api.Responses)
	apiInfo.ParamRef = getParamRef(api.RequestBody, name, apiInfo.FuncName)
	apiInfo.ParamRefName = getRefName(apiInfo.ParamRef)
	apiInfo.RespMode = getRespMode(apiInfo.ResRefName, apiInfo.ParamRef)
	apiInfo.RespFunc = getRespFunc(apiInfo.RespMode)
	apiInfo.Service = getService(apiInfo.FuncName, apiInfo.RespMode, apiInfo.ResRefName)
	return
}

func getFuncName(path string, method string) string {
	// openapi格式默认有斜杠
	arr := strings.Split(path, "/")
	if len(arr) == 3 {
		return strings.Title(method)
	}
	var str string
	for i, v := range arr {
		if i > 2 {
			str += strings.Title(v)
		}
	}
	return str
}
func getParamRef(body RequestBody, name string, funcName string) string {
	ref := body.Content.ApplicationJSON.Schema.Ref
	if ref != "" {
		return ref
	} else {
		return "#/components/schemas/" + name + funcName + "Param"
	}
}

func getResRef(resps map[int]Response) (resRef string, resRefName string) {
	resp := resps[200]
	resRef = resp.Content.ApplicationJSON.Schema.Ref
	resRefName = getRefName(resp.Content.ApplicationJSON.Schema.Ref)
	return
}

func getRespMode(resRefName string, paramRef string) (mode string) {
	if resRefName == "q.Resp" {
		mode = "resp"
		return
	}
	for _, field := range refs[paramRef].FieldInfos {
		if field.Name == "q.Page" {
			mode = "meta"
			return
		}
	}
	mode = "data"
	return
}

func getRespFunc(respMode string) string {
	switch respMode {
	case "resp":
		return "q.Response(r, err)"
	case "meta":
		return "q.ResponseWithMeta(r, err, res, total)"
	case "data":
		return "q.ResponseWithData(r, err, res)"
	default:
		return ""
	}
}

func getService(funcName string, respMode string, res string) (apiService ApiService) {
	switch respMode {
	case "resp":
		apiService.Return = "err"
		apiService.ReturnType = "error"
		return
	case "meta":
		apiService.Return = "res, total, err"
		apiService.ReturnType = "*[]define." + res + ", int64, error"
		return
	default:
		apiService.Return = "res, err"
		apiService.ReturnType = "*define." + res + ", error"
		return
	}
}
