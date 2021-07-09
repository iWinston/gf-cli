package sync

import (
	"strings"

	"github.com/iWinston/gf-cli/commands/sync/apifox"
	"github.com/iWinston/gf-cli/library/utils"
	"github.com/wxnacy/wgo/arrays"
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
	modelInfo.Name = utils.FolderName(schemaItem.Name)
	for k, v := range schemaItem.Schema.JSONSchema.Properties {
		modelInfo.FieldInfos = append(modelInfo.FieldInfos, getModelFields(k, &v, schemaItem.Schema.JSONSchema.Required))
	}
	syncFileForce("app/model", utils.SnakeString(modelInfo.Name)+".model.go", modelTemplate, modelInfo)
	return
}

func getModelFields(key string, field *apifox.Field, required []string) (fm FieldInfo) {
	fm.Name = strings.Title(key)
	fm.Type = getTypeTag(field.Type, field.Items, field.Ref)
	isRequired := arrays.Contains(required, fm.Name) != -1
	fm.Tag = getJsonTag(key) + " " + getOrmTag(field, isRequired) + field.Description
	return
}

func getOrmTag(field *apifox.Field, required bool) string {
	if required {
		return "gorm:not null;"
	}
	return "gorm:"
}

type ModelFolderInfo struct {
	Name string
}

type ModelInfo struct {
	Name       string
	FieldInfos []FieldInfo
}
