package user

import (
	"context"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/util/structure"
)

func GetUserApiKeyDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(UserApiKey))
}

type SchemaUserApiKey schema.UserApiKey

func (a SchemaUserApiKey) ToUserApiKey() *UserApiKey {
	item := new(UserApiKey)
	structure.Copy(a, item)
	return item
}

type UserApiKey struct {
	util.Model
	UserID   uint64 `gorm:"index;default:0;"` // 用户内码
	ApikeyID uint64 `gorm:"index;default:0;"` // API密钥内码
}

func (a UserApiKey) ToSchemaUserApiKey() *schema.UserApiKey {
	item := new(schema.UserApiKey)
	structure.Copy(a, item)
	return item
}

type UserApiKeys []UserApiKey

func (a UserApiKeys) ToSchemaUserApiKeys() []*schema.UserApiKey {
	list := make([]*schema.UserApiKey, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUserApiKey()
	}
	return list
}
