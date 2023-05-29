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

// ToShowResult 转换为显示结果
func (a GroupQueryResult) ToShowResult() *GroupShowQueryResult {
	return &GroupShowQueryResult{
		PageResult: a.PageResult,
		Data:       a.Data.ToGroupShows(),
	}
}

// ToGroupShows 转换为用户显示列表
func (a Groups) ToGroupShows() GroupShows {
	list := make(GroupShows, len(a))
	for i, item := range a {
		list[i] = item.ToShow()
	}
	return list
}

// ToMap 转换为map
func (a Groups) ToMap() map[uint64]*Group {
	m := make(map[uint64]*Group)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

// ToIDs 转换为ID列表
func (a Groups) ToIDs() []uint64 {
	list := make([]uint64, len(a))
	for i, item := range a {
		list[i] = item.ID
	}
	return list
}

// ToNameMap 转换为名称映射
func (a Groups) ToNameMap() map[uint64]string {
	m := make(map[uint64]string)
	for _, item := range a {
		m[item.ID] = item.Name
	}
	return m
}

// ToShows 转换为显示项列表
func (a Groups) ToShows() []*GroupShow {
	shows := make([]*GroupShow, len(a))
	for i, item := range a {
		shows[i] = item.ToShow()
	}
	return shows
}

// ToShow 转换为显示项
func (a *Group) ToShow() *GroupShow {
	return &GroupShow{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Member:      a.Member,
	}
}

// ----------------------------------------GroupShow--------------------------------------

// GroupShow 组显示项
type GroupShow struct {
	ID          uint64 `json:"id"`          // 用户ID
	Name        string `json:"name"`        // 用户名
	Description string `json:"description"` // 用户描述
	//Scope       string   `json:"scope"`       // 用户的权限范围
	Member []uint64 `json:"member"` // 用户的成员
	//Priv        string   `json:"priv"`        // 用户的权限
}

// GroupShows 组显示项列表
type GroupShows []*GroupShow

// GroupShowQueryResult 组显示项查询结果
type GroupShowQueryResult struct {
	Data       GroupShows
	PageResult *PaginationResult
}
