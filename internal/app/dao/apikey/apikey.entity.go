package apikey

import (
	"context"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/util/structure"
)

func GetApiKeyDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(ApiKey))
}

type SchemaApiKey schema.ApiKey

func (a SchemaApiKey) ToApiKey() *ApiKey {
	item := new(ApiKey)
	structure.Copy(a, item)
	return item
}

type ApiKey struct {
	util.Model
	Apikey       string `gorm:"size:256;index;default:'';not null;"` // API密钥
	ApiKeySecret string `gorm:"size:512;index;default:'';not null;"` // API密钥
}

func (a ApiKey) ToSchemaApiKey() *schema.ApiKey {
	item := new(schema.ApiKey)
	structure.Copy(a, item)
	return item
}

type ApiKeys []*ApiKey

func (a ApiKeys) ToSchemaApiKeys() []*schema.ApiKey {
	list := make([]*schema.ApiKey, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaApiKey()
	}
	return list
}
