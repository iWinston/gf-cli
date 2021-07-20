package apifox

type SchemaItem struct {
	Name   string
	Items  []SchemaItem
	Id     string
	Schema struct {
		JSONSchema struct {
			Properties map[string]Field
			Required   []string
			Type       string
			AllOf      []Field
		}
	}
}
type Field struct {
	Properties  map[string]Field
	Type        interface{}
	Items       FieldItems
	Title       string
	Description string
	Default     string
	Minimum     int
	Maximum     int
	MinLength   int
	MaxLength   int
	Pattern     string
	Format      string
	Enum        []string
	Ref         string `json:"$ref"`
}

type FieldItems struct {
	Type string
	Ref  string `json:"$ref"`
}
