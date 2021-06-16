package gen

var defineTemplate = `
package define
import "server/library/restful"

type {TplUpperName}CreateParam struct {
	
}

type {TplUpperName}FindParam struct {
	restful.Page
}

type {TplUpperName}FindOneParam struct {
	
}

type {TplUpperName}PatchOneParam struct {
	
}

type {TplUpperName}DeleteOneParam struct {

}

type {TplUpperName}FindRes struct {

}

type {TplUpperName}FindOneRes struct {

}
`
