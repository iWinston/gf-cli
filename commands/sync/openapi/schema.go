package openapi

type Schemas struct {
	Properties  map[string]Schemas
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
	AllOf       []Schemas
	Required    []string
}

type Field struct {
	Name        string
	Required    []string
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
