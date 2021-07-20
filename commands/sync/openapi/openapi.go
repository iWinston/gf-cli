package openapi

type OpenApi struct {
	Info struct {
		Title       string
		Description string
	}
	Tags []struct {
		Name string
	}
	Paths      map[string]map[string]Api
	Components struct {
		Schemas interface{}
	}
}
