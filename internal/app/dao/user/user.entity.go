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
	Name           string    `gorm:"size:64;index;default:'';not null;"` // 用户名
	Password       string    `gorm:"size:128;default:'';not null;"`      // 密码(sha256哈希值)
	Scope          string    `gorm:"size:64;index;default:'';not null;"` // 用户的权限范围
	Description    string    `gorm:"size:256;default:'';not null;"`      // 描述
	ExpirationAt   time.Time `gorm:"index;"`                             // 过期时间
	AuthorizedKeys string    `gorm:"type:text;default:'';not null;"`     // 公钥
	OtpSeed        string    `gorm:"size:64;default:'';not null;"`       // OTP种子
	Email          string    `gorm:"size:128;default:'';not null;"`      // 邮箱
	Comment        string    `gorm:"size:256;default:'';not null;"`      // 备注
	LandingPage    string    `gorm:"size:256;default:'';not null;"`      // 登录页面
	Shell          string    `gorm:"size:256;default:'';not null;"`      // Shell
	Cert           string    `gorm:"size:256;default:'';not null;"`      // 证书
	ApiKeyId       string    `gorm:"size:256;default:'';not null;"`      // API密钥ID
	Status         int       `gorm:"index;default:0;"`                   // 状态(1:启用 2:停用)

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
