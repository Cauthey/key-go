package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/util/structure"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}

type SchemaUser schema.User

func (a SchemaUser) ToUser() *User {
	item := new(User)
	structure.Copy(a, item)
	return item
}

type User struct {
	util.Model
	Name           string    `json:"user_name"`        // 用户名
	Password       string    `json:"password"`         // 密码
	Scope          string    `json:"scope"`            // 用户的权限范围
	Description    string    `json:"description"`      // 描述
	ExpirationAt   time.Time `json:"expiration_at"`    // 过期时间
	AuthorizedKeys string    `json:"authorized_keys"`  // 授权密钥
	OtpSeed        string    `json:"otp_seed"`         // OTP种子
	Email          string    `json:"email"`            // 邮箱
	Comment        string    `json:"comment"`          // 备注
	LandingPage    string    `json:"landing_page"`     // 登录页面
	Shell          string    `json:"shell"`            // Shell
	Cert           string    `json:"cert"`             // 证书
	ApiKeyId       string    `json:"api_key_id"`       // API密钥ID
	Status         int       `gorm:"index;default:0;"` // 状态(1:启用 2:停用)

}

func (a User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	structure.Copy(a, item)
	return item
}

type Users []*User

func (a Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}
