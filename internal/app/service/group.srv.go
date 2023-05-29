package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"key-go/internal/app/dao"
	"key-go/internal/app/schema"
	"key-go/pkg/errors"
	"key-go/pkg/util/snowflake"
)

var GroupSet = wire.NewSet(wire.Struct(new(GroupSrv), "*"))

type GroupSrv struct {
	Enforcer     *casbin.SyncedEnforcer
	TransRepo    *dao.TransRepo
	UserRepo     *dao.UserRepo
	UserRoleRepo *dao.UserRoleRepo
	RoleRepo     *dao.RoleRepo
	GroupRepo    *dao.GroupRepo
}

func (a *GroupSrv) Query(ctx context.Context, params schema.GroupQueryParam, opts ...schema.GroupQueryOptions) (*schema.GroupQueryResult, error) {
	return a.GroupRepo.Query(ctx, params, opts...)
}

func (a *GroupSrv) QueryShow(ctx context.Context, params schema.GroupQueryParam, opts ...schema.GroupQueryOptions) (*schema.GroupShowQueryResult, error) {
	result, err := a.GroupRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, nil
	}

	return result.ToShowResult(), nil
}

func (a *GroupSrv) Get(ctx context.Context, id uint64, opts ...schema.GroupQueryOptions) (*schema.Group, error) {
	item, err := a.GroupRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (a *GroupSrv) Create(ctx context.Context, item schema.Group) (*schema.IDResult, error) {
	err := a.checkName(ctx, item)
	if err != nil {
		return nil, err
	}
	item.ID = snowflake.MustID()
	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.GroupRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(item.ID), nil
}

func (a *GroupSrv) checkName(ctx context.Context, item schema.Group) error {
	result, err := a.GroupRepo.Query(ctx, schema.GroupQueryParam{
		Name: item.Name,
	})
	if err != nil {
		return err
	} else if len(result.Data) > 0 {
		return errors.New400Response("group name already exists")
	}
	return nil
}

func (a *GroupSrv) Update(ctx context.Context, id uint64, item schema.Group) error {
	oldItem, err := a.GroupRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.New400Response("group does not exist")
	}

	if oldItem.Name != item.Name {
		err = a.checkName(ctx, item)
		if err != nil {
			return err
		}
	}
	item.ID = oldItem.ID
	return a.GroupRepo.Update(ctx, id, item)
}

func (a *GroupSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.GroupRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return nil
	}

	return a.GroupRepo.Delete(ctx, id)
}
