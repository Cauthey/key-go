package user

import (
	"context"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/util/structure"
)

func GetGroupDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(Group))
}

type SchemaGroup schema.Group

type Group struct {
	util.Model
	Name        string   `gorm:"size:64;index;default:'';not null;"` // 组名
	Description string   `gorm:"size:256;default:'';not null;"`      // 描述
	Scope       string   `gorm:"size:64;index;default:'';not null;"` // 组的权限范围
	Member      []uint64 `gorm:"type:text;default:'';not null;"`     // 成员
	Priv        string   `gorm:"size:64;index;default:'';not null;"` // 权限
}

func (a Group) ToSchemaGroup() *schema.Group {
	item := new(schema.Group)
	structure.Copy(a, item)
	return item
}

type Groups []*Group

func (a Groups) ToSchemaGroups() []*schema.Group {
	list := make([]*schema.Group, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaGroup()
	}
	return list
}
