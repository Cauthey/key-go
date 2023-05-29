package schema

// 原配置文件中的sysctl项,暂未明确使用在何处
//
//TODO 此处仅建立该结构体，尚未进行数据库的写入操作及其他任何操作

type Sysctl struct {
	Tunable     string `json:"tunable"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
