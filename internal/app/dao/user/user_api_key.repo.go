package user

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
)

var UserApiKeySet = wire.NewSet(wire.Struct(new(UserApiKeyRepo), "*"))

type UserApiKeyRepo struct {
	DB *gorm.DB
}

func (a *UserApiKeyRepo) getQueryOption(opts ...schema.UserApiKeyQueryOptions) schema.UserApiKeyQueryOptions {
	var opt schema.UserApiKeyQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *UserApiKeyRepo) Query(ctx context.Context, params schema.UserApiKeyQueryParam, opts ...schema.UserApiKeyQueryOptions) (*schema.UserApiKeyQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetUserApiKeyDB(ctx, a.DB)
	if v := params.UserID; v > 0 {
		db = db.Where("user_id=?", v)
	}
	if v := params.UserIDs; len(v) > 0 {
		db = db.Where("user_id IN (?)", v)
	}
	if v := params.ApikeyID; v > 0 {
		db = db.Where("id=?", v)
	}
	if v := params.ApikeyIDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list UserApiKeys
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.UserApiKeyQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUserApiKeys(),
	}

	return qr, nil
}

func (a *UserApiKeyRepo) Get(ctx context.Context, id uint64, opts ...schema.UserApiKeyQueryOptions) (*schema.UserApiKey, error) {
	db := GetUserApiKeyDB(ctx, a.DB).Where("id=?", id)
	var item UserApiKey
	ok, err := util.FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaUserApiKey(), nil
}

func (a *UserApiKeyRepo) Create(ctx context.Context, item schema.UserApiKey) error {
	eitem := SchemaUserApiKey(item).ToUserApiKey()
	result := GetUserApiKeyDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *UserApiKeyRepo) Update(ctx context.Context, id uint64, item schema.UserApiKey) error {
	eitem := SchemaUserApiKey(item).ToUserApiKey()
	result := GetUserApiKeyDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *UserApiKeyRepo) Delete(ctx context.Context, id uint64) error {
	result := GetUserApiKeyDB(ctx, a.DB).Where("id=?", id).Delete(UserApiKey{})
	return errors.WithStack(result.Error)
}

func (a *UserApiKeyRepo) DeleteByUserID(ctx context.Context, userID uint64) error {
	result := GetUserApiKeyDB(ctx, a.DB).Where("user_id=?", userID).Delete(UserApiKey{})
	return errors.WithStack(result.Error)
}

func (a *UserApiKeyRepo) DeleteByApikeyID(ctx context.Context, apikeyID uint64) error {
	result := GetUserApiKeyDB(ctx, a.DB).Where("apikey_id=?", apikeyID).Delete(UserApiKey{})
	return errors.WithStack(result.Error)
}
