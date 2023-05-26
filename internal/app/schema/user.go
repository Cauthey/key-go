package schema

import (
	"context"
	"time"

	"key-go/internal/app/config"
	"key-go/pkg/util/hash"
	"key-go/pkg/util/json"
	"key-go/pkg/util/structure"
)

// GetRootUser 获取root用户
func GetRootUser() *User {
	user := config.C.Root
	return &User{
		ID:   user.UserID,
		Name: user.UserName,
		//RealName: user.RealName,
		Password: hash.MD5String(user.Password),
	}
}

// CheckIsRootUser 检查是否是root用户
func CheckIsRootUser(ctx context.Context, userID uint64) bool {
	return GetRootUser().ID == userID
}

// User 用户对象
type User struct {
	ID             uint64    `json:"id"`               // 用户ID
	Name           string    `json:"user_name"`        // 用户名
	Password       string    `json:"password"`         // 密码
	Scope          string    `json:"scope"`            // 用户的权限范围
	CreatedAt      time.Time `json:"created_at"`       // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`       // 更新时间
	ExpirationAt   time.Time `json:"expiration_at"`    // 过期时间
	Status         int       `gorm:"index;default:0;"` // 状态(1:启用 2:停用)
	Description    string    `json:"description"`      // 描述
	AuthorizedKeys string    `json:"authorized_keys"`  // 授权密钥
	OtpSeed        string    `json:"otp_seed"`         // OTP种子
	Email          string    `json:"email"`            // 邮箱
	Comment        string    `json:"comment"`          // 备注
	LandingPage    string    `json:"landing_page"`     // 登录页面
	Shell          string    `json:"shell"`            // Shell
	Cert           string    `json:"cert"`             // 证书
	ApiKeyId       string    `json:"api_key_id"`       // API密钥ID
}

type Apikey struct {
	ID     uint64 `json:"id,string"` // API密钥ID
	Key    string `json:"key"`       // API密钥
	Secret string `json:"secret"`    // API密钥密钥
}

func (a *User) String() string {
	return json.MarshalToString(a)
}

// CleanSecure 清理安全数据
func (a *User) CleanSecure() *User {
	a.Password = ""
	return a
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	PaginationParam
	Name       string `form:"name"`       // 用户名
	QueryValue string `form:"queryValue"` // 模糊查询
	Status     int    `form:"status"`     // 用户状态(1:启用 2:停用)
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

// UserQueryResult 查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}

// ToShowResult 转换为显示结果
func (a UserQueryResult) ToShowResult(mUserRoles map[uint64]UserRoles, mRoles map[uint64]*Role) *UserShowQueryResult {
	return &UserShowQueryResult{
		PageResult: a.PageResult,
		Data:       a.Data.ToUserShows(mUserRoles, mRoles),
	}
}

// Users 用户对象列表
type Users []*User

// ToIDs 转换为唯一标识列表
func (a Users) ToIDs() []uint64 {
	idList := make([]uint64, len(a))
	for i, item := range a {
		idList[i] = item.ID
	}
	return idList
}

// ToUserShows 转换为用户显示列表
func (a Users) ToUserShows(mUserRoles map[uint64]UserRoles, mRoles map[uint64]*Role) UserShows {
	list := make(UserShows, len(a))
	for i, item := range a {
		showItem := new(UserShow)
		structure.Copy(item, showItem)
		//for _, roleID := range mUserRoles[item.ID].ToRoleIDs() {
		//	if v, ok := mRoles[roleID]; ok {
		//		showItem.Roles = append(showItem.Roles, v)
		//	}
		//}
		list[i] = showItem
	}

	return list
}

// ----------------------------------------UserRole--------------------------------------

// UserRole 用户角色
type UserRole struct {
	ID     uint64 `json:"id,string"`      // 唯一标识
	UserID uint64 `json:"user_id,string"` // 用户ID
	RoleID uint64 `json:"role_id,string"` // 角色ID
}

// UserRoleQueryParam 查询条件
type UserRoleQueryParam struct {
	PaginationParam
	UserID  uint64   // 用户ID
	UserIDs []uint64 // 用户ID列表
}

// UserRoleQueryOptions 查询可选参数项
type UserRoleQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// UserRoleQueryResult 查询结果
type UserRoleQueryResult struct {
	Data       UserRoles
	PageResult *PaginationResult
}

// UserRoles 角色菜单列表
type UserRoles []*UserRole

// ToMap 转换为map
func (a UserRoles) ToMap() map[uint64]*UserRole {
	m := make(map[uint64]*UserRole)
	for _, item := range a {
		m[item.RoleID] = item
	}
	return m
}

// ToRoleIDs 转换为角色ID列表
func (a UserRoles) ToRoleIDs() []uint64 {
	list := make([]uint64, len(a))
	for i, item := range a {
		list[i] = item.RoleID
	}
	return list
}

// ToUserIDMap 转换为用户ID映射
func (a UserRoles) ToUserIDMap() map[uint64]UserRoles {
	m := make(map[uint64]UserRoles)
	for _, item := range a {
		m[item.UserID] = append(m[item.UserID], item)
	}
	return m
}

// ----------------------------------------UserShow--------------------------------------

// UserShow 用户显示项
type UserShow struct {
	//ID        uint64    `json:"id,string"`  // 唯一标识
	//UserName  string    `json:"user_name"`  // 用户名
	//RealName  string    `json:"real_name"`  // 真实姓名
	//Phone     string    `json:"phone"`      // 手机号
	//Email     string    `json:"email"`      // 邮箱
	//Status    int       `json:"status"`     // 用户状态(1:启用 2:停用)
	//CreatedAt time.Time `json:"created_at"` // 创建时间
	//Roles     []*Role   `json:"roles"`      // 授权角色列表

	ID             uint64    `json:"id"`               // 用户ID
	Name           string    `json:"user_name"`        // 用户名
	Scope          string    `json:"scope"`            // 用户的权限范围
	CreatedAt      time.Time `json:"created_at"`       // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`       // 更新时间
	ExpirationAt   time.Time `json:"expiration_at"`    // 过期时间
	Status         int       `gorm:"index;default:0;"` // 状态(1:启用 2:停用)
	Description    string    `json:"description"`      // 描述
	AuthorizedKeys string    `json:"authorized_keys"`  // 授权密钥
	OtpSeed        string    `json:"otp_seed"`         // OTP种子
	Email          string    `json:"email"`            // 邮箱
	Comment        string    `json:"comment"`          // 备注
	LandingPage    string    `json:"landing_page"`     // 登录页面
	Shell          string    `json:"shell"`            // Shell
	Cert           string    `json:"cert"`             // 证书
	ApiKeyId       string    `json:"api_key_id"`       // API密钥ID
}

// UserShows 用户显示项列表
type UserShows []*UserShow

// UserShowQueryResult 用户显示项查询结果
type UserShowQueryResult struct {
	Data       UserShows
	PageResult *PaginationResult
}

// ----------------------------------------UserApiKey--------------------------------------

// UserApiKey 用户API密钥
type UserApiKey struct {
	ID       uint64 `json:"id"`         // 唯一标识
	UserID   uint64 `json:"user_id"`    // 用户ID
	ApiKeyID uint64 `json:"api_key_id"` // API密钥ID
}

// UserApiKeyQueryParam 查询条件
type UserApiKeyQueryParam struct {
	PaginationParam
	UserID    uint64   // 用户ID
	UserIDs   []uint64 // 用户ID列表
	ApikeyID  uint64   // API密钥ID
	ApikeyIDs []uint64 // API密钥ID列表
}

// UserApiKeyQueryOptions 查询可选参数项
type UserApiKeyQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

type UserApiKeys []*UserApiKey

// UserApiKeyQueryResult 查询结果
type UserApiKeyQueryResult struct {
	Data       UserApiKeys
	PageResult *PaginationResult
}
