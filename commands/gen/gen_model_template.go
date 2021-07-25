package gen

var modelTemplate = `
package model

type {TplUpperName} struct {
	Base

}

func (TplUpperName) TableName() string {
	return "srm_label_suppliers"
}
`
