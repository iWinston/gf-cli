package openapi

type Apifox struct {
	Info struct {
		Title       string
		Description string
	}
	Tags       interface{}
	Paths      interface{}
	Components struct {
		Schemas interface{}
	}
}
