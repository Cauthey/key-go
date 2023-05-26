package property

import (
	"context"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/util/structure"

	"gorm.io/gorm"
)

func GetPropertyDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Property))
}

type SchemaProperty schema.Property

func (a SchemaProperty) ToProperty() *Property {
	item := new(Property)
	structure.Copy(a, item)
	return item
}

type Property struct {
	util.Model
	Name  string `gorm:"size:128;index;default:'';not null;"` // 配置名称
	Value string `gorm:"size:256;default:'';"`                // 配置值
}

func (a Property) ToSchemaProperty() *schema.Property {
	item := new(schema.Property)
	structure.Copy(a, item)
	return item
}

type Properties []*Property

func (a Properties) ToSchemaProperties() []*schema.Property {
	list := make([]*schema.Property, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaProperty()
	}
	return list
}
