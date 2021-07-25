package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"

type {TplUpperName}PostParam struct {
	Id *uint {}json:"id" where:""{}
}

type {TplUpperName}Getaram struct {
	
}

type {TplUpperName}PatchParam struct {
	Id *uint {}json:"id" where:""{}
}

type {TplUpperName}DeleteParam struct {
	Id *uint {}json:"id" where:""{}
}

type {TplUpperName}ListParam struct {
	q.Page
}

type {TplUpperName}GetRes struct {
	Id *uint {}json:"id"{}
}

type {TplUpperName}ListRes struct {
	Id *uint {}json:"id"{}
}
`
