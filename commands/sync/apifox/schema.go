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
		}
	}
}

type Field struct {
	Type        string
	Items       FieldItems
	Title       string
	Description string
	Default     string
	MinLength   int
	MaxLength   int
	Pattern     string
	Enum        []string
	Ref         string `json:"$ref"`
}

type FieldItems struct {
	Type string
	Ref  string `json:"$ref"`
}
