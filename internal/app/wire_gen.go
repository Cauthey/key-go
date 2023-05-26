// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"key-go/internal/app/api"
	"key-go/internal/app/dao/menu"
	"key-go/internal/app/dao/role"
	"key-go/internal/app/dao/user"

	//"key-go/internal/app/dao/user"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/module/adapter"
	"key-go/internal/app/router"
	"key-go/internal/app/service"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	roleRepo := &role.RoleRepo{
		DB: db,
	}
	roleMenuRepo := &role.RoleMenuRepo{
		DB: db,
	}
	menuActionResourceRepo := &menu.MenuActionResourceRepo{
		DB: db,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	//userRoleRepo := &user.UserRoleRepo{
	//	DB: db,
	//}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleRepo:         roleRepo,
		RoleMenuRepo:     roleMenuRepo,
		MenuResourceRepo: menuActionResourceRepo,
		UserRepo:         userRepo,
		//UserRoleRepo:     userRoleRepo,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	menuRepo := &menu.MenuRepo{
		DB: db,
	}
	menuActionRepo := &menu.MenuActionRepo{
		DB: db,
	}
	loginSrv := &service.LoginSrv{
		Auth:     auther,
		UserRepo: userRepo,
		//UserRoleRepo:   userRoleRepo,
		RoleRepo:       roleRepo,
		RoleMenuRepo:   roleMenuRepo,
		MenuRepo:       menuRepo,
		MenuActionRepo: menuActionRepo,
	}
	loginAPI := &api.LoginAPI{
		LoginSrv: loginSrv,
	}
	trans := &util.Trans{
		DB: db,
	}
	menuSrv := &service.MenuSrv{
		TransRepo:              trans,
		MenuRepo:               menuRepo,
		MenuActionRepo:         menuActionRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
	}
	menuAPI := &api.MenuAPI{
		MenuSrv: menuSrv,
	}
	roleSrv := &service.RoleSrv{
		Enforcer:               syncedEnforcer,
		TransRepo:              trans,
		RoleRepo:               roleRepo,
		RoleMenuRepo:           roleMenuRepo,
		UserRepo:               userRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
	}
	roleAPI := &api.RoleAPI{
		RoleSrv: roleSrv,
	}
	userSrv := &service.UserSrv{
		Enforcer:  syncedEnforcer,
		TransRepo: trans,
		UserRepo:  userRepo,
		//UserRoleRepo: userRoleRepo,
		//RoleRepo:     roleRepo,
	}
	userAPI := &api.UserAPI{
		UserSrv: userSrv,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		LoginAPI:       loginAPI,
		MenuAPI:        menuAPI,
		RoleAPI:        roleAPI,
		UserAPI:        userAPI,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        menuSrv,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
