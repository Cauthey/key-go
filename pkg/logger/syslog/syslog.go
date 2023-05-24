package syslog

import (
	"key-go/pkg/constant"
	"log/syslog"
	"time"
)

func SendSyslog(level, content string) error {
	sysLog, err := syslog.Dial("tcp", "localhost:514",
		syslog.LOG_DAEMON, "TAF")
	if err != nil {
		return sysLog
	}
	content = "TAF: " + time.Now().Format("2006-01-02 15:04:05") + " " + content
	switch level {
	case constant.LogLevelFatal:
		sysLog.Crit(content)
		break
	case constant.LogLevelError:
		sysLog.Err(content)
		break
	case constant.LogLevelWarn:
		sysLog.Warning(content)
		break
	case constant.LogLevelInfo:
		sysLog.Info(content)
		break
	case constant.LogLevelDebug:
		sysLog.Debug(content)
		break
	default:
		sysLog.Debug(content)
		break
	}
	return nil
}
