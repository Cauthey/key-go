package middleware

import (
	"key-go/internal/app/contextx"
	"key-go/internal/app/ginx"
	"key-go/pkg/auth"
	"key-go/pkg/errors"
	"key-go/pkg/logger"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func wrapUserAuthContext(c *gin.Context, userID uint64, userName string) {
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = contextx.NewUserName(ctx, userName)
	ctx = logger.NewUserIDContext(ctx, userID)
	ctx = logger.NewUserNameContext(ctx, userName)
	c.Request = c.Request.WithContext(ctx)
}

// Valid user token (jwt)
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	//if !config.C.JWTAuth.Enable {
	//	return func(c *gin.Context) {
	//		wrapUserAuthContext(c, config.C.Root.UserID, config.C.Root.UserName)
	//		c.Next()
	//	}
	//}
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			return
		}
		tokenUserID, err := a.ParseUserID(c.Request.Context(), ginx.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				//if config.C.IsDebugMode() {
				//	//wrapUserAuthContext(c, config.C.Root.UserID, config.C.Root.UserName)
				//	//c.Next()
				//	//return
				//	user, err := service.GetRootUser()
				//	if err == nil {
				//		wrapUserAuthContext(c, user.UID, user.Name)
				//		c.Next()
				//		return
				//	}
				//}
				ginx.ResError(c, errors.ErrInvalidToken)
				return
			}
			ginx.ResError(c, errors.WithStack(err))
			return
		}

		idx := strings.Index(tokenUserID, "-")
		if idx == -1 {
			ginx.ResError(c, errors.ErrInvalidToken)
			return
		}

		userID, _ := strconv.ParseUint(tokenUserID[:idx], 10, 64)
		wrapUserAuthContext(c, userID, tokenUserID[idx+1:])
		c.Next()
	}
}
