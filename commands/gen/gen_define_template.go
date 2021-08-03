package gen

var defineTemplate = `
package define
import "github.com/iWinston/qk-library/frame/q"

type {CamelName}PostParam struct {
}

type {CamelName}GetParam struct {
	qfield.Id
}

type {CamelName}GetRes struct {
	Id *uint
}

type {CamelName}PatchParam struct {
	qfield.Id
}

type {CamelName}DeleteParam struct {
	qfield.Id
}

type {CamelName}ListParam struct {
	q.Page
}

type {CamelName}ListRes struct {
	Id *uint
}
`
