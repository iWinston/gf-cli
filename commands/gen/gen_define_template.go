package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"

type {TplUpperName}PostParam struct {
	
}

type {TplUpperName}FindParam struct {
	q.Page
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
