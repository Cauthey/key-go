package constant

// 日志存放路径

const (
	// AuditLogPath 审计日志：记录所有用户的登录、退出、权限切换等信息
	AuditLogPath = "/var/log/audit/audit.log"

	// SystemLogPath 系统日志：记录所有与系统安全有关的事件
	SystemLogPath = "/var/log/system/system.log"

	//// AccessLogPath 访问日志：记录所有进出防火墙的数据包，包括源 IP、目标 IP、协议、端口等信息
	//AccessLogPath = "/var/log/access.log"
	//
	//// IntrusionLogPath 入侵检测日志：记录所有与已知攻击行为相匹配的事件，包括攻击类型、攻击IP、攻击时间等信息
	//IntrusionLogPath = "/var/log/intrusion.log"
	//
	//// UserBehaviorLogPath 用户行为日志：记录所有用户对防火墙的行为，包括登录、权限切换、命令执行等信息
	//UserBehaviorLogPath = "/var/log/user_behavior.log"
	//
	//// SystemEventLogPath 系统事件日志：记录与系统安全有关的事件，包括服务启动、关机、崩溃等信息
	//SystemEventLogPath = "/var/log/system_event.log"
	//
	//// SecurityPolicyLogPath 安全策略日志：记录防火墙的安全策略修改、增加或删除等变化，包括策略名称、修改者、修改时间等信息
	//SecurityPolicyLogPath = "/var/log/security_policy.log"
	//
	//// FirewallTrafficAnalysisLogPath 流量分析日志：记录防火墙内部网络流量的分析结果，包括端口扫描、连接失败等信息
	//FirewallTrafficAnalysisLogPath = "/var/log/firewall_traffic_analysis.log"
	//
	//// NotificationLogPath 通知日志：记录防火墙的警报和通知信息，包括邮件通知、短信通知等信息
	//NotificationLogPath = "/var/log/notification.log"
	//
	//// SystemPerformanceLogPath 系统性能日志：记录系统的性能信息，包括CPU、内存、磁盘空间等指标
	//SystemPerformanceLogPath = "/var/log/system_performance.log"
)

// 配置文件路径

const (
	SystemConfigXmlPath = "/conf/config.xml"
)

// 系统配置属性

var SystemConfigPropertiesInitData = map[string]string{
	"system.optimization":                  "normal",
	"system.hostname":                      "TAF 1.0",
	"system.domain":                        "localdomain",
	"system.dnsallowoverride":              "1",
	"system.nextuid":                       "2000",
	"system.nextgid":                       "2000",
	"system.timezone":                      "Etc/UTC",
	"system.timeservers":                   "ntp1.aliyun.com,ntp2.aliyun.com,ntp3.aliyun.com",
	"system.disablenatreflection":          "yes",
	"system.usevirtualterminal":            "1",
	"system.disableconsolemenu":            "",
	"system.disablevlanhwfilter":           "1",
	"system.disablechecksumoffloading":     "1",
	"system.disablesegmentationoffloading": "1",
	"system.disablelargereceiveoffloading": "1",
	"system.ipv6allow":                     "",
	"system.powerd_ac_mode":                "hadp",
	"system.powerd_battery_mode":           "hadp",
	"system.powerd_normal_mode":            "hadp",
	"system.bogons.interval":               "monthly",
	"system.pf_share_forward":              "1",
	"system.lb_use_sticky":                 "1",
	"system.ssh.group":                     "1",
	"system.rrdbackup":                     "-1",
	"system.netflowbackup":                 "-1",
}
