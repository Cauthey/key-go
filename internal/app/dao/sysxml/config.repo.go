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
