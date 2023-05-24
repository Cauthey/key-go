package sysxml

import (
	"encoding/xml"
	"key-go/pkg/constant"
	"os"
)

func Get() (taf TAF, err error) {
	// 读取/conf/config.xml 中的所有用户信息
	file, err := os.Open(constant.SystemConfigXmlPath)
	if err != nil {
		return TAF{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	var configXml TAF
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&configXml)
	if err != nil {
		return TAF{}, err
	}
	return configXml, nil
}

func UserAdd(user User) (err error) {
	// 读取/conf/config.xml 中的所有用户信息
	file, err := os.Open(constant.SystemConfigXmlPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	var configXml TAF
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&configXml)
	if err != nil {
		return err
	}
	// 添加用户
	configXml.System.User = append(configXml.System.User, user)
	// 保存
	file, err = os.OpenFile(constant.SystemConfigXmlPath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	encoder := xml.NewEncoder(file)
	err = encoder.Encode(configXml)
	if err != nil {
		return err
	}
	return nil
}

func UserDel(user User) (err error) {
	// 读取/conf/config.xml 中的所有用户信息
	file, err := os.Open(constant.SystemConfigXmlPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	var configXml TAF
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&configXml)
	if err != nil {
		return err
	}
	// 删除用户
	for i, v := range configXml.System.User {
		if v.Name == user.Name {
			configXml.System.User = append(configXml.System.User[:i], configXml.System.User[i+1:]...)
			break
		}
	}
	// 保存
	file, err = os.OpenFile(constant.SystemConfigXmlPath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	encoder := xml.NewEncoder(file)
	err = encoder.Encode(configXml)
	if err != nil {
		return err
	}
	return nil
}

func UserGroupAdd(userGroup Group) (err error) {
	// 读取/conf/config.xml 中的所有用户信息
	file, err := os.Open(constant.SystemConfigXmlPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	var configXml TAF
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&configXml)
	if err != nil {
		return err
	}
	// 添加用户
	configXml.System.UserGroup = append(configXml.System.UserGroup, userGroup)
	// 保存
	file, err = os.OpenFile(constant.SystemConfigXmlPath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(configXml)
	if err != nil {
		return err
	}
	return nil
}
