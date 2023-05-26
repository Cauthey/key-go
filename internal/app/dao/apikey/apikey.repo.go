package apikey

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/errors"
)

var ApiKeySet = wire.NewSet(wire.Struct(new(ApiKeyRepo), "*"))

type ApiKeyRepo struct {
	DB *gorm.DB
}

func (a *ApiKeyRepo) getQueryOption(ctx context.Context, opts ...schema.ApiKeyQueryOptions) schema.ApiKeyQueryOptions {
	var opt schema.ApiKeyQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *ApiKeyRepo) Query(ctx context.Context, params schema.ApiKeyQueryParam, opts ...schema.ApiKeyQueryOptions) (*schema.ApiKeyQueryResult, error) {
	opt := a.getQueryOption(ctx, opts...)

	db := GetApiKeyDB(ctx, a.DB)
	if v := params.Apikey; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("apikey LIKE ? ", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list ApiKeys
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.ApiKeyQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaApiKeys(),
	}
	return qr, nil
}

func (a *ApiKeyRepo) Get(ctx context.Context, id uint64, opts ...schema.ApiKeyQueryOptions) (*schema.ApiKey, error) {
	db := GetApiKeyDB(ctx, a.DB).Where("id=?", id)
	var item ApiKey
	ok, err := util.FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaApiKey(), nil
}

func (a *ApiKeyRepo) Create(ctx context.Context, item schema.ApiKey) error {
	eitem := SchemaApiKey(item).ToApiKey()
	result := GetApiKeyDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *ApiKeyRepo) Update(ctx context.Context, id uint64, item schema.ApiKey) error {
	eitem := SchemaApiKey(item).ToApiKey()
	result := GetApiKeyDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *ApiKeyRepo) Delete(ctx context.Context, id uint64) error {
	result := GetApiKeyDB(ctx, a.DB).Where("id=?", id).Delete(ApiKey{})
	return errors.WithStack(result.Error)
}
