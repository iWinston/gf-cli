package openapi

type Api struct {
	Summary     string
	Description string
	Tags        []string
	Parameters  []Parameter
	RequestBody RequestBody
	Responses   map[int]Response
}

type Parameter struct {
	Name        string
	In          string
	Description string
	Required    bool
	Example     string
	Schema      struct {
		Type string
	}
}

type RequestBody struct {
	Content struct {
		ApplicationJSON struct {
			Schema struct {
				Ref string `json:"$ref"`
			}
		} `json:"application/json"`
	}
}

type Response struct {
	Description string
	Content     struct {
		ApplicationJSON struct {
			Schema struct {
				Ref string `json:"$ref"`
			}
			Examples struct {
			}
		} `json:"application/json"`
	}
}
