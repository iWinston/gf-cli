package sync

import (
	"errors"
	"strings"

	"github.com/iWinston/gf-cli/commands/sync/apifox"
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
	Name     string
	FuncName string
	Param    string
	RespType string
	Path     string
	Method   string
	Service  ApiService
	RespFunc string
	RespMode string
	ParamRef string
	ResRef   string
}

type ApiService struct {
	Return     string
	ReturnType string
}

func HelpApi() {

}

func doSyncApi(apiCollection *[]apifox.ApiItem) {
	for _, apiSystem := range *apiCollection {
		genApiFiles(&apiSystem)
	}
}

func genApiFiles(apiSystem *apifox.ApiItem) {
	for _, apiFolder := range apiSystem.Items {
		genApiFile(apiSystem.Name, &apiFolder)
	}
}

func genApiFile(systemName string, apiFolder *apifox.ApiItem) {
	var fileInfo ApiFileInfo
	fileInfo.System = utils.SnakeString(systemName)
	fileInfo.StructName = utils.GetStructName(apiFolder.Name)
	fileInfo.FileName = utils.GetFileName(apiFolder.Name)
	fileInfo.Name = strings.ToLower(fileInfo.StructName)
	for _, apiItem := range apiFolder.Items {
		fileInfo.ApiInfos = append(fileInfo.ApiInfos, getApi(&apiItem))
	}
	syncFile("app/system/"+fileInfo.System+"/api", fileInfo.FileName+".api.go", apiTemplate, fileInfo)
	syncFileForce("app/system/"+fileInfo.System+"/api/internal", fileInfo.FileName+".api.go", apiInternalTemplate, fileInfo)
	syncFile("app/system/"+fileInfo.System+"/service", fileInfo.FileName+".service.go", serviceTemplate, fileInfo)
	syncFileForce("app/system/"+fileInfo.System+"/service/internal", fileInfo.FileName+".service.go", serviceInternalTemplate, fileInfo)
}

func getApi(apiItem *apifox.ApiItem) (apiInfo ApiInfo) {
	apiInfo.Name = apiItem.Name
	apiInfo.Path = apiItem.Api.Path
	apiInfo.Method = apiItem.Api.Method
	apiInfo.FuncName = getFuncName(apiInfo.Path, apiInfo.Method)
	apiInfo.ResRef, _ = getResRef(apiItem.Api.Responses)
	apiInfo.ParamRef = getParamRef(apiItem.Api.RequestBody)
	apiInfo.RespMode = getRespMode(apiInfo.ResRef, apiItem.Api.RequestBody.JsonSchema.Ref)
	apiInfo.RespFunc = getRespFunc(apiInfo.RespMode)
	apiInfo.Service = getService(apiInfo.FuncName, apiInfo.RespMode, apiInfo.ResRef)
	return
}

func getFuncName(path string, method string) string {
	arr := strings.Split(path, "/")
	if len(arr) == 2 {
		return strings.Title(method)
	}
	var str string
	for i, v := range arr {
		if i > 1 {
			str += strings.Title(v)
		}
	}
	return str
}

func getParamRef(requestBody apifox.RequestBody) string {
	return getRef(requestBody.JsonSchema.Ref)
}

func getResRef(resps []apifox.Response) (resRef string, err error) {
	if len(resps) != 1 {
		err = errors.New("缺少Resp")
		return
	}
	resp := resps[0]
	resRef = getRef(resp.JsonSchema.Ref)
	return

}

func getRespMode(resRef string, paramRefId string) (mode string) {
	if resRef == "q.Resp" {
		mode = "resp"
		return
	}
	for _, field := range refs[paramRefId].FieldInfos {
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
