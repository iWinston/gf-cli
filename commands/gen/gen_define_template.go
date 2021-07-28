package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"

type {CamelName}PostParam struct {
}

type {CamelName}GetParam struct {
	Id *uint {Backtick}where:"=" example:"1" v:"required|integer#id必填|id必须为整数"{Backtick}
}

type {CamelName}GetRes struct {
	Id *uint {Backtick}json:"id"{Backtick}
}

type {CamelName}PatchParam struct {
	Id *uint {Backtick}where:"=" example:"1" v:"required|integer#id必填|id必须为整数"{Backtick}
}

type {CamelName}DeleteParam struct {
	Id *uint {Backtick}where:"=" example:"1" v:"required|integer#id必填|id必须为整数"{Backtick}
}

type {CamelName}ListParam struct {
	q.Page
}

type {CamelName}ListRes struct {
	Id *uint {Backtick}json:"id"{Backtick}
}
`
