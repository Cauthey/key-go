package sysxml

import (
	"encoding/xml"
)

type TAF struct {
	XMLName xml.Name `xml:"TAF"`
	Theme   string   `xml:"theme,attr"`
	Sysctl  Sysctl   `xml:"sysctl"`
	System  System   `xml:"system"`
}

type Sysctl struct {
	XMLName xml.Name `xml:"sysctl"`
	//Item    []Item   `xml:"item"`
}

type System struct {
	XMLName                       xml.Name `xml:"system"`
	Optimization                  string   `xml:"optimization"`
	Hostname                      string   `xml:"hostname"`
	Domain                        string   `xml:"domain"`
	DnsAllowOverRide              string   `xml:"dnsallowoverride"`
	User                          []User   `xml:"user"`
	UserGroup                     []Group  `xml:"group"`
	NextUid                       uint64   `xml:"nextuid"`
	NextGid                       uint64   `xml:"nextgid"`
	Timezone                      string   `xml:"timezone"`
	Timeservers                   string   `xml:"timeservers"`
	WebGui                        WebGui   `xml:"webgui"`
	DisableNatReflection          string   `xml:"disablenatreflection"`
	UseVirtualTerminal            string   `xml:"usevirtualterminal"`
	DisableConsoleMenu            string   `xml:"disableconsolemenu"`
	DisableVlanHwFilter           string   `xml:"disablevlanhwfilter"`
	DisableChecksumOffloading     string   `xml:"disablechecksumoffloading"`
	DisableSegmentationOffloading string   `xml:"disablesegmentationoffloading"`
	DisableLargeReceiveOffloading string   `xml:"disablelargereceiveoffloading"`
	Ipv6Allow                     string   `xml:"ipv6allow"`
	PowerAcMode                   string   `xml:"powerd_ac_mode"`
	PowerBatteryMode              string   `xml:"powerd_battery_mode"`
	PowerNormalMode               string   `xml:"powerd_normal_mode"`
	Bogon                         Bogon    `xml:"bogons"`
	PfShareForward                string   `xml:"pf_share_forward"`
	LbUseSticky                   int      `xml:"lb_use_sticky"`
	SSH                           SSH      `xml:"ssh"`
	RrdBackup                     string   `xml:"rrdbackup"`
	NetFlowBackup                 string   `xml:"netflowbackup"`
}

type User struct {
	XMLName        xml.Name `xml:"user"`
	Name           string   `xml:"name"`
	Description    string   `xml:"descr"`
	GroupName      string   `xml:"groupname"`
	Password       string   `xml:"password"`
	UID            uint64   `xml:"uid"`
	Expires        string   `xml:"expires"`
	AuthorizedKeys string   `xml:"authorizedkeys"`
	OtpSeed        string   `xml:"otp_seed"`
	Comment        string   `xml:"comment"`
	LandingPage    string   `xml:"landing_page"`
	Shell          string   `xml:"shell"`
	Cert           string   `xml:"cert"`
}

type Group struct {
	XMLName     xml.Name `xml:"group"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Scope       string   `xml:"scope"`
	GiD         uint64   `xml:"gid"`
	Member      []uint64 `xml:"member"`
	Priv        string   `xml:"priv"`
}
type Item struct {
	XMLName xml.Name `xml:"item"`
	Descr   string   `xml:"descr"`
	Tunable string   `xml:"tunable"`
	Value   string   `xml:"value"`
}

type WebGui struct {
	XMLName     xml.Name `xml:"webgui"`
	Protocol    string   `xml:"protocol"`
	SslCertRef  string   `xml:"ssl-certref"`
	NoAuto      string   `xml:"noauto"`
	Compression string   `xml:"compression"`
	Interfaces  string   `xml:"interfaces"`
}

type Bogon struct {
	XMLName  xml.Name `xml:"bogons"`
	Interval string   `xml:"interval"`
}

type SSH struct {
	XMLName xml.Name `xml:"ssh"`
	Group   string   `xml:"group"`
}
