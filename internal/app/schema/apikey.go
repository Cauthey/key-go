package schema

type ApiKey struct {
	ApikeyID     uint64 `json:"apikey_id"`
	Apikey       string `json:"apikey" `
	ApikeySecret string `json:"apikey_secret" `
}

// ApiKeys API密钥列表
type ApiKeys []*ApiKey

// ApiKeyQueryParam 查询条件
type ApiKeyQueryParam struct {
	PaginationParam
	Apikey     string `form:"apikey"`     // API密钥
	QueryValue string `form:"queryValue"` // 模糊查询
}

// ApiKeyQueryOptions 查询可选参数项
type ApiKeyQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// ApiKeyQueryResult 查询结果
type ApiKeyQueryResult struct {
	Data       ApiKeys
	PageResult *PaginationResult
}
