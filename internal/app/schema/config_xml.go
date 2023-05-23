package schema

type TAF struct {
	Theme  string `json:"theme"`
	Sysctl Sysctl `json:"sysctl"`
	System System `json:"system"`
}

type Sysctl struct {
	//Item    []Item  `json:"item"`
}

type System struct {
	Optimization                  string  `json:"optimization"`
	Hostname                      string  `json:"hostname"`
	Domain                        string  `json:"domain"`
	DnsAllowOverRide              string  `json:"dnsAllowOverRide"`
	User                          []User  `json:"user"`
	UserGroup                     []Group `json:"group"`
	NextUid                       uint64  `json:"nextUid"`
	NextGid                       uint64  `json:"nextGid"`
	TimeZone                      string  `json:"timeZone"`
	TimeServers                   string  `json:"timeServers"`
	WebGui                        WebGui  `json:"webGui"`
	DisableNatReflection          string  `json:"disableNatReflection"`
	UseVirtualTerminal            string  `json:"useVirtualTerminal"`
	DisableConsoleMenu            string  `json:"disableConsoleMenu"`
	DisableVlanHwFilter           string  `json:"disableVlanHwFilter"`
	DisableChecksumOffloading     string  `json:"disableChecksumOffloading"`
	DisableSegmentationOffloading string  `json:"disableSegmentationOffloading"`
	DisableLargeReceiveOffloading string  `json:"disableLargeReceiveOffloading"`
	Ipv6Allow                     string  `json:"ipv6Allow"`
	PowerAcMode                   string  `json:"powerAcMode"`
	PowerBatteryMode              string  `json:"powerBatteryMode"`
	PowerNormalMode               string  `json:"powerNormalMode"`
	Bogon                         Bogon   `json:"bogon"`
	PfShareForward                string  `json:"pfShareForward"`
	LbUseSticky                   int     `json:"lbUseSticky"`
	SSH                           SSH     `json:"ssh"`
	RrdBackup                     string  `json:"rrdBackup"`
	NetFlowBackup                 string  `json:"netFlowBackup"`
}

type Item struct {
	Descr   string `json:"descr"`
	Tunable string `json:"tunable"`
	Value   string `json:"value"`
}

type WebGui struct {
	Protocol    string `json:"protocol"`
	SslCertRef  string `json:"sslCertRef"`
	NoAuto      string `json:"noAuto"`
	Compression string `json:"compression"`
	Interfaces  string `json:"interfaces"`
}

type Bogon struct {
	Interval string `json:"interval"`
}

type SSH struct {
	Group string `json:"group"`
}
