package apifox

type Apifox struct {
	Info struct {
		Name        string
		Description string
	}
	SchemaCollection []SchemaItem
	ApiCollection    []ApiItem
}
