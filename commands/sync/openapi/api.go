package openapi

type Api struct {
	Summary     string
	Description string
	Tags        []string
	Parameters  []interface{}
	RequestBody struct {
		Content struct {
			ApplicationJSON struct {
				Schema struct {
					Ref string `json:"$ref"`
				}
			}
		}
	}
	Responses map[int]Response
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
		}
	}
}
