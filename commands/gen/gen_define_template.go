package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"

type {CamelName}PostParam struct {
	Id *uint {Backtick}where:""{Backtick}
}

type {CamelName}Getaram struct {
	
}

type {CamelName}PatchParam struct {
	Id *uint {Backtick}where:""{Backtick}
}

type {CamelName}DeleteParam struct {
	Id *uint {Backtick}where:""{Backtick}
}

type {CamelName}ListParam struct {
	q.Page
}

type {CamelName}GetRes struct {
	Id *uint {Backtick}json:"id"{Backtick}
}

type {CamelName}ListRes struct {
	Id *uint {Backtick}json:"id"{Backtick}
}
`
