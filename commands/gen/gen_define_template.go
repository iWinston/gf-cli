package gen

var defineTemplate = `
package define
import "server/app/dto"

type {TplUpperName}CreateParam struct {
	
}

type {TplUpperName}FindParam struct {
	dto.Page
}

type {TplUpperName}FindOneParam struct {
	
}

type {TplUpperName}PatchOneParam struct {
	
}

type {TplUpperName}DeleteOneParam struct {
	dto.BatchIds

}

type {TplUpperName}FindRes struct {

}

type {TplUpperName}FindOneRes struct {

}
`
