package openapi

type OpenApi struct {
	Info struct {
		Title       string
		Description string
	}
	Tags       []Tag
	Paths      map[string]map[string]Api
	Components struct {
		Schemas map[string]Schemas
	}
}

type Tag struct {
	Name string
}
