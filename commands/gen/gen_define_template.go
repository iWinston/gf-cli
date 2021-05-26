package gen

var defineTemplate = `
package define
import "server/app/model"
type {TplUpperName} = model.{TplUpperName}MODEL

type {TplUpperName}CreateReq struct {
	
}

type {TplUpperName}GetReq struct {
	
}

type {TplUpperName}PatchOneReq struct {
	
}

type {TplUpperName}PatchReq struct {
	
}

type {TplUpperName}GetOneRes struct {
	
}

type {TplUpperName}GetRes struct {
	
}
`
