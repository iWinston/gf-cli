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
		if schemaSystem.Name == "Model" || schemaSystem.Name == "Schemas" || schemaSystem.Name == "Q" {
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
	if schemaItem.Schema.JSONSchema.AllOf != nil {
		// 特殊处理继承的情况
		for _, v := range schemaItem.Schema.JSONSchema.AllOf {
			if v.Ref != "" {
				defineInfo.FieldInfos = append(defineInfo.FieldInfos, FieldInfo{Name: getRef(v.Ref), Tag: v.Description})
			} else {
				for k, v := range v.Properties {
					defineInfo.FieldInfos = append(defineInfo.FieldInfos, getFields(k, &v, schemaItem.Schema.JSONSchema.Required))
				}
			}
		}
	} else {
		for k, v := range schemaItem.Schema.JSONSchema.Properties {
			defineInfo.FieldInfos = append(defineInfo.FieldInfos, getFields(k, &v, schemaItem.Schema.JSONSchema.Required))
		}
	}
	return
}

func getFields(key string, field *apifox.Field, required []string) (fm FieldInfo) {
	fm.Name = strings.Title(key)
	fm.Type = getTypeTag(field)
	isRequired := arrays.Contains(required, fm.Name) != -1
	tagArr := []string{getJsonTag(key), field.Description, getValidateTag(field, isRequired)}
	fm.Tag = joinNotEmpty(tagArr, " ")
	return
}

func getJsonTag(name string) string {
	return `json:"` + name + `"`
}

func getTypeTag(field *apifox.Field) string {
	if field.Ref != "" {
		return getRef(field.Ref)
	}
	if fieldType, ok := field.Type.(string); ok {
		return getFieldType(field, fieldType)
	} else {
		return getFieldType(field, field.Type.([]interface{})[0].(string))
	}
}

func getFieldType(field *apifox.Field, fieldType string) string {
	switch fieldType {
	case "string":
		if field.Format == "date" || field.Format == "date-time" {
			return "*time.Time"
		}
		return "*string"
	case "integer":
		return "*int"
	case "boolean":
		return "*bool"
	case "array":
		if field.Items.Ref != "" {
			return "[]" + getRef(field.Items.Ref)
		}
		return "[]" + field.Items.Type
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
	if arrays.Contains([]string{"integer", "string", "boolean"}, field.Type) != -1 {
		rules = append(rules, field.Type.(string))
		msgs = append(msgs, field.Title+"不符合类型规则")
	}
	if len(rules) == 0 || len(msgs) == 0 {
		return ""
	}
	return `v:"` + strings.Join(rules, "|") + "#" + strings.Join(msgs, "|") + `"`
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
