package sync

import (
	"strings"

	"github.com/iWinston/gf-cli/commands/sync/apifox"
	"github.com/iWinston/gf-cli/library/utils"
)

func HelpModel() {

}

func doSyncModel(schemaCollection *[]apifox.SchemaItem) {
	for _, schemaFolder := range *schemaCollection {
		if schemaFolder.Name == "Model" {
			genModels(&schemaFolder)
		}
	}
}

func genModels(schemaFolder *apifox.SchemaItem) (modelFolderInfo ModelFolderInfo) {
	for _, schemaItem := range schemaFolder.Items {
		genModel(&schemaItem)
	}
	return
}

func genModel(schemaItem *apifox.SchemaItem) (modelInfo ModelInfo) {
	modelInfo.Name = utils.GetFileName(schemaItem.Name)
	modelInfo.StructName = strings.Title(modelInfo.Name)
	for k, v := range schemaItem.Schema.JSONSchema.Properties {
		modelInfo.FieldInfos = append(modelInfo.FieldInfos, getModelFields(k, &v, schemaItem.Schema.JSONSchema.Required))
	}
	syncFileForce("app/model", utils.SnakeString(modelInfo.Name)+".model.go", modelTemplate, modelInfo)
	return
}

func getModelFields(key string, field *apifox.Field, required []string) (fm FieldInfo) {
	fm.Name = strings.Title(key)
	fm.Type = getTypeTag(field)
	tagArr := []string{getJsonTag(key), getOrmTag(field)}
	fm.Tag = strings.Join(tagArr, " ")
	return
}

func getOrmTag(field *apifox.Field) string {
	var arr []string
	if _, ok := field.Type.(string); ok {
		arr = append(arr, "not null")
	}
	if field.Default != "" {
		arr = append(arr, "default:"+field.Default)
	}
	if field.Title != "" {
		arr = append(arr, "comment:"+field.Title)
	}
	if field.MaxLength != 0 {
		arr = append(arr, "length:"+field.Title)
	}
	if field.Description != "" {
		arr = append(arr, field.Description)
	}
	if len(arr) == 0 {
		return ""
	}
	return `gorm:"` + strings.Join(arr, ";") + `"`
}

type ModelFolderInfo struct {
	Name string
}

type ModelInfo struct {
	Name       string
	StructName string
	FieldInfos []FieldInfo
}
