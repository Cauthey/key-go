package api

import (
	"fmt"
	"github.com/LyricTian/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"key-go/internal/app/contextx"
	"key-go/internal/app/ginx"
	"key-go/internal/app/schema"
	"key-go/internal/app/service"
	"key-go/pkg/errors"
	"key-go/pkg/logger"

	"key-go/internal/app/config"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginAPI), "*"))

type LoginAPI struct {
	LoginSrv *service.LoginSrv
}

func (a *LoginAPI) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.LoginSrv.GetCaptcha(ctx, config.C.Captcha.Length)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *LoginAPI) ResCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	captchaID := c.Query("id")
	if captchaID == "" {
		ginx.ResError(c, errors.New400Response("captcha id not empty"))
		return
	}

	if c.Query("reload") != "" {
		if !captcha.Reload(captchaID) {
			ginx.ResError(c, errors.New400Response("not found captcha id"))
			return
		}
	}

	cfg := config.C.Captcha
	err := a.LoginSrv.ResCaptcha(ctx, c.Writer, captchaID, cfg.Width, cfg.Height)
	if err != nil {
		ginx.ResError(c, err)
	}
}

func (a *LoginAPI) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	//if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
	//	ginx.ResError(c, errors.New400Response("无效的验证码"))
	//	return
	//}

	// 验证用户是否存在/获取用户
	ok, user := service.GetUserByUsername(item.UserName)
	if !ok {
		ginx.ResError(c, errors.New400Response("not found user_name"))
		return
	}

	// 验证用户登录密码
	err := a.LoginSrv.Verify(item.Password, user.Password)
	if err != nil {
		logger.WriteAuditLog(fmt.Sprintf("用户[%s]登录失败，密码错误", user.Name))
		ginx.ResError(c, err)
		return
	}
	// 验证用户是否在有效期内
	ok = a.LoginSrv.VerifyUserValid(user)
	if !ok {
		logger.WriteAuditLog(fmt.Sprintf("用户[%s]登录失败，用户未在有效期内", user.Name))
		ginx.ResError(c, err)
		return
	}

	tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(uint64(user.UID), user.Name))
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, tokenInfo)
}

func (a *LoginAPI) formatTokenUserID(userID uint64, userName string) string {
	return fmt.Sprintf("%d-%s", userID, userName)
}

func (a *LoginAPI) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	userID := contextx.FromUserID(ctx)

	fmt.Println("userID ========================>>>>", userID)

	fmt.Println("ginx.GetToken(c)========================>>>>", ginx.GetToken(c))

	if userID != 0 {
		ctx = logger.NewTagContext(ctx, "__logout__")
		err := a.LoginSrv.DestroyToken(ctx, ginx.GetToken(c))
		if err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
		logger.WithContext(ctx).Infof("logout")
	}
	ginx.ResOK(c)
}

func (a *LoginAPI) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(contextx.FromUserID(ctx), contextx.FromUserName(ctx)))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, tokenInfo)
}

func (a *LoginAPI) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Println("contextx.FromUserID(ctx)========================>>>>", contextx.FromUserID(ctx))
	info, err := a.LoginSrv.GetLoginInfo(ctx, contextx.FromUserID(ctx))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, info)
}

func (a *LoginAPI) QueryUserMenuTree(c *gin.Context) {
	//ctx := c.Request.Context()
	//menus, err := a.LoginSrv.QueryUserMenuTree(ctx, contextx.FromUserID(ctx))
	//if err != nil {
	//	ginx.ResError(c, err)
	//	return
	//}
	//ginx.ResList(c, menus)
}

func (a *LoginAPI) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.UpdatePasswordParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.LoginSrv.UpdatePassword(ctx, contextx.FromUserID(ctx), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
