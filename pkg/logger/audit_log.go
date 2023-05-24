package logger

import (
	"key-go/pkg/constant"
	"key-go/pkg/util/files"
	"path"
	"strings"
	"time"
)

func WriteAuditLog(info string) {
	logName := "audit_" + time.Now().Format("20060102") + ".log"
	logPath := path.Join(constant.AuditLogPath, logName)
	index := strings.Index(info, "\n")
	if index == -1 || index == len(info)-1 {
		info += "\n"
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	err := files.WriteFileByAddContent(logPath, nowStr+" "+info)
	if err != nil {
		panic(err)
	}
	return
}
