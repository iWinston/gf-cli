package sync

import (
	"strconv"
	"strings"

	"github.com/iWinston/gf-cli/commands/sync/apifox"
	"github.com/iWinston/gf-cli/library/utils"
	"github.com/wxnacy/wgo/arrays"
)

func HelpDefine() {

}

func doSyncDefine(schemaCollection *[]apifox.SchemaItem) {
	for _, schemaSystem := range *schemaCollection {
		if schemaSystem.Name == "Model" || schemaSystem.Name == "Schemas" || strings.HasPrefix(schemaSystem.Name, "q") {
			continue
		}
		genDefineFiles(&schemaSystem)
	}
}

func genDefineFiles(schemaSystem *apifox.SchemaItem) {
	for _, schemaFolder := range schemaSystem.Items {
		genDefineFile(&schemaFolder, utils.SnakeString(schemaSystem.Name))
	}
}

func genDefineFile(schemaFolder *apifox.SchemaItem, systemName string) (defineFileInfo DefineFileInfo) {
	defineFileInfo.Name = utils.GetFileName(schemaFolder.Name) + ".define.go"
	for _, schemaItem := range schemaFolder.Items {
		defineFileInfo.DefineInfos = append(defineFileInfo.DefineInfos, getDefine(&schemaItem, systemName))
	}
	syncFileForce("app/system/"+systemName+"/define", defineFileInfo.Name, defineTemplate, defineFileInfo.DefineInfos)
	return
}

func getDefine(schemaItem *apifox.SchemaItem, systemName string) (defineInfo DefineInfo) {
	defineInfo.System = systemName
	defineInfo.Name = schemaItem.Name
	for k, v := range schemaItem.Schema.JSONSchema.Properties {
		defineInfo.FieldInfos = append(defineInfo.FieldInfos, getFields(k, &v, schemaItem.Schema.JSONSchema.Required))
	}
	return
}

func getFields(key string, field *apifox.Field, required []string) (fm FieldInfo) {
	fm.Name = strings.Title(key)
	fm.Type = getTypeTag(field.Type, field.Items, field.Ref)
	isRequired := arrays.Contains(required, fm.Name) != -1
	fm.Tag = getJsonTag(key) + " " + field.Description + " " + getValidateTag(field, isRequired)
	return
}

func getJsonTag(name string) string {
	return `json:"` + name + `"`
}

func getTypeTag(fieldType string, items apifox.FieldItems, ref string) string {
	if ref != "" {
		return getRef(ref)
	}

	switch fieldType {
	case "string":
		return "*string"
	case "integer":
		return "*int"
	case "boolean":
		return "*bool"
	case "array":
		if items.Ref != "" {
			return "[]" + getRef(items.Ref)
		}
		return "[]" + items.Type
	case "number":
		return "*float"
	case "any":
		return "interface{}"
	case "object":
		//TODO 递归
		return "*" + fieldType
	default:
		return "interface{}"
	}
}

func getValidateTag(field *apifox.Field, required bool) string {
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
	if len(rules) == 0 || len(msgs) == 0 {
		return ""
	}
	return strings.Join(rules, "|") + "#" + strings.Join(msgs, "|")
}

type DefineFileInfo struct {
	System      string
	Name        string
	DefineInfos []DefineInfo
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
