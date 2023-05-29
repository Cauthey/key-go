package app

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"key-go/internal/app/service"
	"key-go/pkg/auth"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	UserSrv        *service.UserSrv
	MenuSrv        *service.MenuSrv
	PropertySrv    *service.PropertySrv
}
