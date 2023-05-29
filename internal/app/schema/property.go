package schema

type Property struct {
	Name  string `json:"name" `
	Value string `json:"value" `
}

// Properties 用户对象列表
type Properties []*Property

// PropertyQueryParam 查询条件
type PropertyQueryParam struct {
	PaginationParam
	Name       string `form:"name"`       // 用户名
	QueryValue string `form:"queryValue"` // 模糊查询
}

// PropertyQueryOptions 查询可选参数项
type PropertyQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// PropertyQueryResult 查询结果
type PropertyQueryResult struct {
	Data       Properties
	PageResult *PaginationResult
}
