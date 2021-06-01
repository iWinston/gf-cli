package gen

var defineTemplate = `
package define
import "server/app/dto"
import "server/app/model"
type {TplUpperName} = model.{TplUpperName}Model

type {TplUpperName}CreateReq struct {
	
}

type {TplUpperName}GetReq struct {
	dto.Page
}

type {TplUpperName}PatchOneReq struct {
	
}

type {TplUpperName}PatchReq struct {
	dto.BatchIds

}

type {TplUpperName}GetOneRes struct {

}

type {TplUpperName}GetRes struct {

}
`
