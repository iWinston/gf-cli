package gen

var modelTemplate = `
package model

type {CamelName} struct {
	Base

}

func ({CamelName}) TableName() string {
	return "{SnakeSystemName}_{SnakeName}"
}
`
