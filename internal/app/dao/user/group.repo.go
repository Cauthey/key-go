package user

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/errors"
)

var GroupSet = wire.NewSet(wire.Struct(new(GroupRepo), "*"))

type GroupRepo struct {
	DB *gorm.DB
}

func (a *GroupRepo) getQueryOption(ctx context.Context, opts ...schema.GroupQueryOptions) schema.GroupQueryOptions {
	var opt schema.GroupQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *GroupRepo) Query(ctx context.Context, params schema.GroupQueryParam, opts ...schema.GroupQueryOptions) (*schema.GroupQueryResult, error) {
	opt := a.getQueryOption(ctx, opts...)

	db := GetGroupDB(ctx, a.DB)
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	//if v := params.Status; v > 0 {
	//	db = db.Where("status=?", v)
	//}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name LIKE ? OR description LIKE ?", v, v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(opt.OrderFields)
	}

	var list Groups
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.GroupQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaGroups(),
	}
	return qr, nil

}

func (a *GroupRepo) Get(ctx context.Context, id uint64) (*schema.Group, error) {
	var item Group
	ok, err := util.FindOne(ctx, GetGroupDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaGroup(), nil
}

func (a *GroupRepo) Create(ctx context.Context, item schema.Group) error {
	eitem := SchemaGroup(item)
	result := GetGroupDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *GroupRepo) Update(ctx context.Context, id uint64, item schema.Group) error {
	eitem := SchemaGroup(item)
	result := GetGroupDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *GroupRepo) Delete(ctx context.Context, id uint64) error {
	result := GetGroupDB(ctx, a.DB).Where("id=?", id).Delete(Group{})
	return errors.WithStack(result.Error)
}
