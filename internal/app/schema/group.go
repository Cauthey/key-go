package schema

type Group struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scope       string   `json:"scope"`
	ID          uint64   `json:"id"`
	Member      []uint64 `json:"member"`
	Priv        string   `json:"priv"`
}

// Groups 用户对象列表
type Groups []*Group

// GroupQueryParam 查询条件
type GroupQueryParam struct {
	PaginationParam
	Name       string `form:"name"`       // 用户名
	QueryValue string `form:"queryValue"` // 模糊查询

}

// GroupQueryOptions 查询可选参数项
type GroupQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// GroupQueryResult 查询结果
type GroupQueryResult struct {
	Data       Groups
	PageResult *PaginationResult
}
