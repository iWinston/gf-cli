package apifox

type ApiItem struct {
	Name  string
	Items []ApiItem
	Api   struct {
		Id         string
		Method     string
		Path       string
		Parameters struct {
			Path   []string
			Header []string
			Query  []string
			Cookie []string
		}
		RequestBody RequestBody
		Tags        []string
		Responses   []Response
	}
}

type RequestBody struct {
	Type       string
	Parameters []interface{}
	JsonSchema struct {
		Ref string `json:"$ref"`
	}
}

type Response struct {
	Id         string
	Name       string
	Code       int
	JsonSchema struct {
		Ref string `json:"$ref"`
	}
}
