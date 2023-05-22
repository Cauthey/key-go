package audit

import (
	"key-go/pkg/constant"
	"key-go/pkg/util/files"
	"path"
	"strings"
	"time"
)

func WriteAuditLog(info string) (err error) {
	logName := "audit_" + time.Now().Format("20060102") + ".log"
	logPath := path.Join(constant.AuditLogPath, logName)
	index := strings.Index(info, "\n")
	if index == -1 || index == len(info)-1 {
		info += "\n"
	}
	err = files.WriteFileByAddContent(logPath, info)
	if err != nil {
		return err
	}
	return nil
}
