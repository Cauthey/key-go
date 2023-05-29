package service

import (
	"context"
	"github.com/google/wire"
	"key-go/internal/app/dao"
	"key-go/internal/app/schema"
	"key-go/pkg/auth"
	"key-go/pkg/constant"
	"key-go/pkg/errors"
)

var PropertySet = wire.NewSet(wire.Struct(new(PropertySrv), "*"))

type PropertySrv struct {
	PropertyRepo   *dao.PropertyRepo
	Auth           auth.Auther
	UserRepo       *dao.UserRepo
	RoleRepo       *dao.RoleRepo
	RoleMenuRepo   *dao.RoleMenuRepo
	MenuRepo       *dao.MenuRepo
	MenuActionRepo *dao.MenuActionRepo
	//UserRoleRepo   *dao.UserRoleRepo
}

func (a *PropertySrv) Query(ctx context.Context, params schema.PropertyQueryParam, opts ...schema.PropertyQueryOptions) (*schema.PropertyQueryResult, error) {
	return a.PropertyRepo.Query(ctx, params, opts...)
}

func (a *PropertySrv) Get(ctx context.Context, name string, opts ...schema.PropertyQueryOptions) (*schema.Property, error) {
	item, err := a.PropertyRepo.Get(ctx, name, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.New400Response("configuration's property does not exist")
	}
	return item, nil
}

func (a *PropertySrv) Count(ctx context.Context) (int64, error) {
	return a.PropertyRepo.Count(ctx)
}

func (a *PropertySrv) Create(ctx context.Context, item schema.Property) error {
	if err := a.checkName(ctx, item); err != nil {
		return err
	}
	return a.PropertyRepo.Create(ctx, item)
}

func (a *PropertySrv) CreateByMap(ctx context.Context, item map[string]string) error {
	return a.PropertyRepo.CreateByMap(ctx, item)
}

func (a *PropertySrv) Update(ctx context.Context, name string, item schema.Property) error {
	oldItem, err := a.PropertyRepo.Get(ctx, name)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.New400Response("configuration's property does not exist")
	}

	return a.PropertyRepo.Update(ctx, name, item)
}

func (a *PropertySrv) checkName(ctx context.Context, item schema.Property) error {
	result, err := a.PropertyRepo.Query(ctx, schema.PropertyQueryParam{
		Name: item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("configuration's property has been exists")
	}
	return nil
}

func (a *PropertySrv) InitData(ctx context.Context) error {
	count, err := a.Count(ctx)
	if err != nil {
		return err
	}
	if count != 0 {
		return nil
	}
	return a.CreateByMap(ctx, constant.SystemConfigPropertiesInitData)
}
