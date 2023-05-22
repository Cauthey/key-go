package sysxml

import "encoding/xml"

type TAF struct {
	XMLName xml.Name `xml:"TAF"`
	Theme   string   `xml:"theme,attr"`
	Sysctl  Sysctl   `xml:"sysctl"`
	System  System   `xml:"system"`
}

type Sysctl struct {
	XMLName xml.Name `xml:"sysctl"`
	Item    []Item   `xml:"item"`
}

type System struct {
	XMLName                       xml.Name        `xml:"system"`
	Optimization                  string          `xml:"optimization"`
	Hostname                      string          `xml:"hostname"`
	Domain                        string          `xml:"domain"`
	DnsAllowOverRide              string          `xml:"dnsallowoverride"`
	User                          []UserConfigXml `xml:"user"`
	UserGroup                     []UserGroups    `xml:"group"`
	NextUid                       int             `xml:"nextuid"`
	NextGid                       int             `xml:"nextgid"`
	Timezone                      string          `xml:"timezone"`
	Timeservers                   string          `xml:"timeservers"`
	WebGui                        WebGui          `xml:"webgui"`
	DisableNatReflection          string          `xml:"disablenatreflection"`
	UseVirtualTerminal            string          `xml:"usevirtualterminal"`
	DisableConsoleMenu            string          `xml:"disableconsolemenu"`
	DisableVlanHwFilter           string          `xml:"disablevlanhwfilter"`
	DisableChecksumOffloading     string          `xml:"disablechecksumoffloading"`
	DisableSegmentationOffloading string          `xml:"disablesegmentationoffloading"`
	DisableLargeReceiveOffloading string          `xml:"disablelargereceiveoffloading"`
	Ipv6Allow                     string          `xml:"ipv6allow"`
	PowerAcMode                   string          `xml:"powerd_ac_mode"`
	PowerBatteryMode              string          `xml:"powerd_battery_mode"`
	PowerNormalMode               string          `xml:"powerd_normal_mode"`
	Bogon                         Bogon           `xml:"bogons"`
	PfShareForward                string          `xml:"pf_share_forward"`
	LbUseSticky                   int             `xml:"lb_use_sticky"`
	SSH                           SSH             `xml:"ssh"`
	RrdBackup                     string          `xml:"rrdbackup"`
	NetFlowBackup                 string          `xml:"netflowbackup"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Descr   string   `xml:"descr"`
	Tunable string   `xml:"tunable"`
	Value   string   `xml:"value"`
}

type UserConfigXml struct {
	XMLName     xml.Name `xml:"user"`
	Name        string   `xml:"name"`
	Description string   `xml:"descr"`
	GroupName   string   `xml:"groupname"`
	Password    string   `xml:"password"`
	UID         uint64   `xml:"uid"`
}

type UserGroups struct {
	XMLName     xml.Name `xml:"group"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Scope       string   `xml:"scope"`
	GID         int      `xml:"gid"`
	Member      int      `xml:"member"`
	Priv        string   `xml:"priv"`
}

type WebGui struct {
	XMLName  xml.Name `xml:"webgui"`
	Protocol string   `xml:"protocol"`
}

type Bogon struct {
	XMLName  xml.Name `xml:"bogons"`
	Interval string   `xml:"interval"`
}

type SSH struct {
	XMLName xml.Name `xml:"ssh"`
	Group   string   `xml:"group"`
}
