package openapi

type DtoFileInfo struct {
	System      string
	Name        string
	DefineInfos []DefineInfo
}

type DtoInfo struct {
	System       string
	Name         string
	DtoFieldInfo []DtoFieldInfo
}

type DtoFieldInfo struct {
	Name     string
	Type     string
	Validate string
}
