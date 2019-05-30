package iland

import "time"

type Domain struct {
	ID   string `json:"domain_id"`
	Name string `json:"domain_name"`
	Type string `json:"domain_type"`
}

type DHCP struct {
	Enabled          bool    `json:"enabled"`
	IPRange          IPRange `json:"ip_range"`
	DefaultLeaseTime int     `json:"default_lease_time"`
	MaxLeaseTime     int     `json:"max_lease_time"`
}

type SubnetParticipation struct {
	Gateway   string    `json:"gateway"`
	Netmask   string    `json:"netmask"`
	IPAddress string    `json:"ip_address"`
	IPRanges  []IPRange `json:"ip_ranges"`
}

type IPScope struct {
	Inherited    bool      `json:"inherited"`
	Enabled      bool      `json:"enabled"`
	Gateway      string    `json:"gateway"`
	Netmask      string    `json:"netmask"`
	PrimaryDNS   string    `json:"primary_dns"`
	SecondaryDNS string    `json:"secondary_dns"`
	DNSSuffix    string    `json:"dns_suffix"`
	IPRanges     []IPRange `json:"ip_ranges"`
}

type IPRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Snapshot struct {
	Size         int64 `json:"size"`
	IsPoweredOn  bool  `json:"is_powered_on"`
	CreationDate int64 `json:"creation_date"`
}

type Metadata struct {
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
	Type   string      `json:"type"`
	Access string      `json:"access"`
}

func getUnixMilliseconds(datetime time.Time) int {
	return int(datetime.UnixNano() / int64(time.Millisecond))
}
