package gen

var defineTemplate = `
package define
import "server/app/dto"
import "server/app/model"
type {TplUpperName} = model.{TplUpperName}Model

type {TplUpperName}CreateDto struct {
	
}

type {TplUpperName}GetDto struct {
	dto.Page
}

type {TplUpperName}PatchOneDto struct {
	
}

type {TplUpperName}PatchDto struct {
	dto.BatchIds

}

type {TplUpperName}GetOneRes struct {

}

type {TplUpperName}GetRes struct {

}
`
