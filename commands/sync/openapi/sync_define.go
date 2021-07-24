package openapi

import (
	"strconv"
	"strings"

	"github.com/iWinston/gf-cli/commands/sync"
	"github.com/iWinston/gf-cli/library/utils"
	"github.com/wxnacy/wgo/arrays"
)

type DefineFileInfo struct {
	System      string
	Name        string
	DefineInfos []DefineInfo
	Imports     []string
}

type DefineInfo struct {
	System     string
	Name       string
	FieldInfos []FieldInfo
}

type FieldInfo struct {
	Name string
	Type string
	Tag  string
}

func syncDefineFile(fileInfos *map[string]*DefineFileInfo) {
	for _, fileInfo := range *fileInfos {
		imports := []string{}
		importMap := map[string]string{
			"model":  "server/app/model",
			"q":      "github.com/iWinston/qk-library/frame/q",
			"qfield": "github.com/iWinston/qk-library/frame/qfield",
		}
		for _, defineInfo := range fileInfo.DefineInfos {
			for _, fieldInfo := range defineInfo.FieldInfos {
				arr := strings.Split(fieldInfo.Type, "[]")
				if len(arr) > 1 {
					addImport(&imports, arr[1])
				} else {
					addImport(&imports, fieldInfo.Type)
					// 继承模式
					addImport(&imports, fieldInfo.Name)
				}
			}
		}
		if len(imports) > 0 {
			for i, item := range imports {
				imports[i] = importMap[item]
			}
		}
		fileInfo.Imports = imports

		sync.SyncFileForce("app/system/"+fileInfo.System+"/define", fileInfo.Name+".define.go", sync.DefineTemplate, fileInfo)
	}
}

func addImport(imports *[]string, fieldType string) {
	arr := strings.Split(fieldType, ".")
	if len(arr) > 1 {
		if arrays.ContainsString(*imports, arr[0]) == -1 {
			*imports = append(*imports, arr[0])
		}
	}
}

func getRefsBySchemas(refs *map[string]DefineInfo, schemas *map[string]Schemas) {
	for k, v := range *schemas {
		(*refs)["#/components/schemas/"+k] = getDefineInfo(k, &v)
	}
}

func getDefineInfo(key string, schema *Schemas) (defineInfo DefineInfo) {
	// keyArr := strings.Split(key, ".")
	// defineInfo.System = keyArr[0]
	defineInfo.Name = key
	if schema.AllOf != nil {
		// 特殊处理继承的情况
		for _, v := range schema.AllOf {
			if v.Ref != "" {
				defineInfo.FieldInfos = append(defineInfo.FieldInfos, FieldInfo{Name: getRefName(v.Ref), Tag: v.Description})
			} else {
				for k, v := range v.Properties {
					defineInfo.FieldInfos = append(defineInfo.FieldInfos, getFields(k, &v, schema.Required))
				}
			}
		}
	} else {
		for k, v := range schema.Properties {
			defineInfo.FieldInfos = append(defineInfo.FieldInfos, getFields(k, &v, schema.Required))
		}
	}
	return
}

func getFields(key string, field *Schemas, required []string) (fm FieldInfo) {
	isRequired := arrays.Contains(required, key) != -1
	fm.Name = strings.Title(key)
	fm.Type = getTypeTag(field)
	tagArr := []string{getJsonTag(key), field.Description, getValidateTag(field, isRequired)}
	fm.Tag = utils.JoinNotEmptyStr(tagArr, " ")
	return
}

func getJsonTag(name string) string {
	return `json:"` + name + `"`
}

func getTypeTag(field *Schemas) string {
	if field.Ref != "" {
		return getRefName(field.Ref)
	}
	if fieldType, ok := field.Type.(string); ok {
		return getFieldType(field, fieldType)
	} else {
		return getFieldType(field, field.Type.([]interface{})[0].(string))
	}
}

func getRefName(ref string) string {
	keyArr := strings.Split(ref, "/")
	name := keyArr[len(keyArr)-1]
	if name == "time.Date" {
		return "time.Time"
	}
	return name
}

func getFieldType(field *Schemas, fieldType string) string {
	if fieldType == "any" {
		return "interface{}"
	}
	if fieldType == "array" {
		if field.Items.Ref != "" {
			return "[]" + getRefName(field.Items.Ref)
		}
		return "[]" + getBaseType(field, field.Items.Type)
	}
	return "*" + getBaseType(field, fieldType)
}

func getBaseType(field *Schemas, fieldType string) string {
	switch fieldType {
	case "string":
		if field.Format == "date" || field.Format == "date-time" {
			return "time.Time"
		}
		return "string"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "number":
		return "float"
	case "object":
		//TODO 递归
		return fieldType
	default:
		return ""
	}
}

func getValidateTag(field *Schemas, required bool) string {
	var (
		rules, msgs []string
	)
	if required {
		rules = append(rules, "required")
		msgs = append(msgs, field.Title+"不能为空")
	}
	if field.Pattern != "" {
		rules = append(rules, field.Pattern)
		msgs = append(msgs, field.Title+"不符合规则")
	}
	if field.Format != "" {
		if arrays.Contains([]string{"date", "email", "ipv4", "ipv6"}, field.Format) != -1 {
			rules = append(rules, field.Format)
			msgs = append(msgs, field.Title+"不符合规则")
		}
		if field.Format == "date-time" {
			rules = append(rules, "date-format")
			msgs = append(msgs, field.Title+"不符合规则")
		}
	}
	if len(field.Enum) != 0 {
		enumStr := strings.Join(field.Enum, ",")
		rules = append(rules, "in:"+enumStr)
		msgs = append(msgs, field.Title+"应该是以下值:"+enumStr)
	}
	if field.MinLength != 0 {
		rules = append(rules, "min-length:"+strconv.Itoa(field.MinLength))
		msgs = append(msgs, field.Title+"最小长度为:"+strconv.Itoa(field.MinLength))
	}
	if field.MaxLength != 0 {
		rules = append(rules, "max-length:"+strconv.Itoa(field.MaxLength))
		msgs = append(msgs, field.Title+"最大长度为:"+strconv.Itoa(field.MaxLength))
	}
	if field.Minimum != 0 {
		rules = append(rules, "min:"+strconv.Itoa(field.Minimum))
		msgs = append(msgs, field.Title+"最小为:"+strconv.Itoa(field.Minimum))
	}
	if field.Maximum != 0 {
		rules = append(rules, "max:"+strconv.Itoa(field.Maximum))
		msgs = append(msgs, field.Title+"最大为:"+strconv.Itoa(field.Maximum))
	}
	if field.Type == "number" {
		rules = append(rules, "float")
		msgs = append(msgs, field.Title+"必须为浮点数")
	}
	if field.Type == "integer" {
		rules = append(rules, field.Type.(string))
		msgs = append(msgs, field.Title+"必须为整数")
	}
	if field.Type == "boolean" {
		rules = append(rules, field.Type.(string))
		msgs = append(msgs, field.Title+"必须为布尔值")
	}
	if len(rules) == 0 || len(msgs) == 0 {
		return ""
	}
	return `v:"` + strings.Join(rules, "|") + "#" + strings.Join(msgs, "|") + `"`
}

// func getRefsByPathes(refs *map[string]DefineInfo, pathes *map[string]map[string]Api) {
// 	for path, apis := range *pathes {
// 		for method, api := range apis {
// 			if len(api.Parameters) > 0 {
// 				var defineInfo DefineInfo
// 				tagName := api.Tags[0]
// 				refName := getFuncName(path, method) + tagName + "Param"
// 				defineInfo.Name = refName
// 				ref := getFieldsByParameters(api.Parameters)
// 			}

// 		}
// 	}
// }

// func getFieldsByParameters(parameters []Parameter) (fieldInfos []FieldInfo) {
// 	for _, parameter := range parameters {
// 		var field FieldInfo
// 		field.Name = strings.Title(parameter.Name)
// 	}
// }
