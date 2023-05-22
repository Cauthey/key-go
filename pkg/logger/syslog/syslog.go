package syslog

import (
	"key-go/pkg/constant"
	"log/syslog"
)

func SendSyslog(level, content string) error {
	sysLog, err := syslog.Dial("udp", "localhost:514",
		syslog.LOG_DAEMON, "TAF")
	if err != nil {
		return err
	}
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
