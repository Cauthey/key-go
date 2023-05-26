package schema

type Property struct {
	Name  string `gorm:"type:varchar(64);primary_key;not null;comment:配置名称"  json:"name"  validate:"required,max=64" `
	Value string `gorm:"type:varchar(2048);default:'';comment:配置值"  json:"value"            validate:"max=256"  `
}
