package logger

import (
	"key-go/pkg/constant"
	"key-go/pkg/util/files"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func AuditLog(info string) {
	logName := "audit.log" + "." + time.Now().Format("2006-01-02")
	logPath := path.Join(filepath.Dir(constant.AuditLogPath), logName)
	if !files.FileExists(logPath) {
		// 创建文件
		_, err := os.Create(logPath)
		if err != nil {
			Errorf("创建文件失败：%s", err.Error())
		}
		// 更新软连接audit.log的链接地址
		_ = os.Remove(constant.AuditLogPath)
		err = os.Symlink(logPath, constant.AuditLogPath)
		if err != nil {
			Errorf("创建软链接失败：%s", err.Error())
		}
	}
	index := strings.Index(info, "\n")
	if index == -1 || index != len(info)-1 {
		info += "\n"
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	err := files.WriteFileByAddContent(logPath, nowStr+" "+info)
	if err != nil {
		Errorf("写入文件失败：%s", err.Error())
	}
	return
}
