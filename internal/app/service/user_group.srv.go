package service

import (
	"key-go/internal/app/dao/sysxml"
	"key-go/internal/app/schema"
)

// GetUserGroup 获取所有用户组
func GetUserGroup() ([]schema.Group, error) {
	result, err := sysxml.Get()
	if err != nil {
		return nil, err
	}
	var groups = make([]schema.Group, len(result.System.UserGroup))
	for i, v := range result.System.UserGroup {
		groups[i] = schema.Group{
			GiD:         v.GiD,
			Name:        v.Name,
			Member:      v.Member,
			Description: v.Description,
			Scope:       v.Scope,
			Priv:        v.Priv,
		}
	}
	return groups, nil
}

// GetUserGroupByUserId 获取某用户所在的所有用户组
func GetUserGroupByUserId(uid uint) ([]schema.Group, error) {
	result, err := sysxml.Get()
	if err != nil {
		return nil, err
	}
	var groups = make([]schema.Group, 0)
	for _, v := range result.System.UserGroup {
		for _, v1 := range v.Member {
			if v1 == uid {
				groups = append(groups, schema.Group{
					GiD:         v.GiD,
					Name:        v.Name,
					Member:      v.Member,
					Description: v.Description,
					Scope:       v.Scope,
					Priv:        v.Priv,
				})
				continue
			}
		}
	}
	return groups, nil
}
