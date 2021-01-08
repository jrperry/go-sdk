package iland

import (
	"encoding/json"
	"fmt"
)

type Edge struct {
	ID                        string          `json:"uuid"`
	Name                      string          `json:"name"`
	Description               string          `json:"description"`
	Status                    int             `json:"status"`
	Interfaces                []EdgeInterface `json:"interfaces"`
	Size                      string          `json:"gateway_backing_config"`
	HighAvailabilityEnabled   bool            `json:"high_availability_enabled"`
	DefaultDNSRelay           bool            `json:"default_dns_relay_route"`
	BackwardCompatibilityMode bool            `json:"backward_compatibility_mode"`
	VdcID                     string          `json:"vdc_uuid"`
	OrgID                     string          `json:"org_uuid"`
	CompanyID                 string          `json:"company_id"`
	LocationID                string          `json:"location_id"`
	UpdatedDate               int             `json:"updated_date"`
}

type EdgeInterface struct {
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	NetworkID      string                `json:"network_uuid"`
	NetworkName    string                `json:"network"`
	DefaultRoute   bool                  `json:"default_route"`
	ApplyRateLimit bool                  `json:"apply_rate_limit"`
	InRateLimit    int                   `json:"in_rate_limit"`
	OutRateLimit   int                   `json:"out_rate_limit"`
	Subnets        []SubnetParticipation `json:"subnet_participation"`
}

type edgeService struct {
	client *client
}

func (s *edgeService) Get(edgeID string) (Edge, error) {
	edge := Edge{}
	err := s.client.getObject(fmt.Sprintf("/v1/edges/%s", edgeID), &edge)
	if err != nil {
		return Edge{}, err
	}
	return edge, nil
}

// type EdgeFirewall struct {
// 	Log           bool               `json:"log"`
// 	Enabled       bool               `json:"enabled"`
// 	DefaultAction string             `json:"default_action"`
// 	Rules         []EdgeFirewallRule `json:"rules"`
// }

// type EdgeFirewallRule struct {
// 	ID                   string   `json:"id"`
// 	Order                int      `json:"idx"`
// 	Description          string   `json:"description"`
// 	Enabled              bool     `json:"enabled"`
// 	Policy               string   `json:"policy"`
// 	SourceIP             string   `json:"source_ip,omitempty"`
// 	SourcePort           int      `json:"source_port,omitempty"`
// 	SourcePortRange      string   `json:"source_port_range,omitempty"`
// 	DestinationIP        string   `json:"destination_ip,omitempty"`
// 	DestinationPort      int      `json:"port,omitempty"`
// 	DestinationPortRange string   `json:"destination_port_range,omitempty"`
// 	Direction            string   `json:"direction,omitempty"`
// 	Protocols            []string `json:"protocols"`
// 	MatchOnTranslate     bool     `json:"match_on_translate,omitempty"`
// }

// func (s *edgeService) GetFirewall(edgeID string) (EdgeFirewall, error) {
// 	firewall := EdgeFirewall{}
// 	err := s.client.getObject(fmt.Sprintf("/v1/edges/%s/firewall", edgeID), &firewall)
// 	if err != nil {
// 		return EdgeFirewall{}, err
// 	}
// 	return firewall, nil
// }

// func (s *edgeService) UpdateFirewallRules(edgeID string, rules []EdgeFirewallRule) (Task, error) {
// 	firewall, err := s.GetFirewall(edgeID)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	firewall.Rules = rules
// 	data, err := json.Marshal(&firewall)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-firewall", edgeID), data)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	task := Task{}
// 	err = unmarshalBody(resp, &task)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	return task, nil
// }

type EdgeFirewall struct {
	Enabled bool               `json:"enabled"`
	Config  EdgeFirewallConfig `json:"firewall_global_config"`
	Policy  EdgeFirewallPolicy `json:"firewall_default_policy"`
	Rules   []EdgeFirewallRule `json:"firewall_rules"`
}

type EdgeFirewallPolicy struct {
	DefaultAction  string `json:"firewall_action_type"`
	LoggingEnabled bool   `json:"logging_enabled"`
}

type EdgeFirewallConfig struct {
	TCPPickOngoingConnections bool `json:"tcp_pick_ongoing_connections"`
	TCPAllowOutOfWindowPakets bool `json:"tcp_allow_out_of_window_packets"`
	TCPSendClosedVSEPortReset bool `json:"tcp_send_reset_for_closed_vse_ports"`
	DropInvalidTraffic        bool `json:"drop_invalid_traffic"`
	LogInvalidTraffic         bool `json:"log_invalid_traffic"`
	TCPTimeoutOpen            int  `json:"tcp_timeout_open"`
	TCPTimeoutEstablished     int  `json:"tcp_timeout_established"`
	TCPTimeoutClose           int  `json:"tcp_timeout_close"`
	UDPTimeout                int  `json:"udp_timeout"`
	ICMPTimeout               int  `json:"icmp_timeout"`
	ICMP6Timeout              int  `json:"icmp6_timeout"`
	IPGenericTimeout          int  `json:"ip_generic_timeout"`
	EnableSynFloodProtection  bool `json:"enable_syn_flood_protection"`
	LogICMPErrors             bool `json:"log_icmp_errors"`
	DropICMReplays            bool `json:"drop_icmp_replays"`
	EnableSNMPAlg             bool `json:"enable_snmp_alg"`
	EnableFTPAlg              bool `json:"enable_ftp_alg"`
	EnableTFTPAlg             bool `json:"enable_tftp_alg"`
}

type EdgeFirewallRule struct {
	ID              int                         `json:"id"`
	RuleTag         int                         `json:"rule_tag"`
	Name            string                      `json:"name"`
	Type            string                      `json:"rule_type"`
	Description     string                      `json:"description"`
	Source          EdgeFirewallRuleOrigin      `json:"source"`
	Destination     EdgeFirewallRuleOrigin      `json:"destination"`
	Application     EdgeFirewallRuleApplication `json:"application"`
	Enabled         bool                        `json:"enabled"`
	Action          string                      `json:"action_type"`
	LoggingEnabled  bool                        `json:"logging_enabled"`
	MatchTranslated bool                        `json:"match_translated"`
}

type EdgeFirewallRuleOrigin struct {
	Exclude          bool     `json:"exclude"`
	IPAddresses      []string `json:"ip_addresses"`
	GroupingObjectID []string `json:"grouping_object_id"`
	VNicGroupID      []string `json:"vnic_group_id"`
}

type EdgeFirewallRuleApplication struct {
	ID       string                    `json:"application_id,omitempty"`
	Services []EdgeFirewallRuleService `json:"service"`
}

type EdgeFirewallRuleService struct {
	Protocol   string   `json:"protocol"`
	Port       []string `json:"port"`
	SourcePort []string `json:"source_port"`
	ICMPType   string   `json:"icmp_type,omitempty"`
}

func (s *edgeService) GetFirewall(edgeID string) (EdgeFirewall, error) {
	firewall := EdgeFirewall{}
	err := s.client.getObject(fmt.Sprintf("/v1/edge-gateways/%s/firewall", edgeID), &firewall)
	if err != nil {
		return EdgeFirewall{}, err
	}
	return firewall, nil
}

func (s *edgeService) UpdateFirewall(edgeID string, firewall EdgeFirewall) (EdgeFirewall, error) {
	data, err := json.Marshal(&firewall)
	if err != nil {
		return EdgeFirewall{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/edge-gateways/%s/firewall", edgeID), data)
	if err != nil {
		return EdgeFirewall{}, err
	}
	output := EdgeFirewall{}
	err = unmarshalBody(resp, &output)
	if err != nil {
		return EdgeFirewall{}, err
	}
	return output, nil
}

func (s *edgeService) GetNAT(edgeID string) (EdgeNAT, error) {
	nat := EdgeNAT{}
	err := s.client.getObject(fmt.Sprintf("/v1/edge-gateways/%s/nat", edgeID), &nat)
	if err != nil {
		return EdgeNAT{}, err
	}
	return nat, nil
}

func (s *edgeService) UpdateNAT(edgeID string, nat EdgeNAT) (EdgeNAT, error) {
	data, err := json.Marshal(&nat)
	if err != nil {
		return EdgeNAT{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/edge-gateways/%s/nat", edgeID), data)
	if err != nil {
		return EdgeNAT{}, err
	}
	output := EdgeNAT{}
	err = unmarshalBody(resp, &output)
	if err != nil {
		return EdgeNAT{}, err
	}
	return output, nil
}

type EdgeNAT struct {
	Enabled    bool            `json:"enabled"`
	Rules      []EdgeNATRule   `json:"nat_rules"`
	NAT64Rules []EdgeNAT64Rule `json:"nat64_rules"`
}

type EdgeNATRule struct {
	ID                          int    `json:"rule_id"`
	Tag                         int    `json:"rule_tag"`
	LoggingEnabled              bool   `json:"logging_enabled"`
	Enabled                     bool   `json:"enabled"`
	Description                 string `json:"description"`
	TranslatedAddress           string `json:"translated_address"`
	Type                        string `json:"rule_type"`
	Action                      string `json:"action"`
	VNIC                        int    `json:"vnic"`
	OriginalAddress             string `json:"original_address"`
	DNATMatchSourceAddress      string `json:"dnat_match_source_address,omitempty"`
	SNATMatchDestinationAddress string `json:"snat_match_destination_address,omitempty"`
	Protocol                    string `json:"protocol"`
	OriginalPort                string `json:"original_port"`
	TranslatedPort              string `json:"translated_port"`
	DNATMatchSourcePort         string `json:"dnat_match_source_port,omitempty"`
	DNATMatchDestinationPort    string `json:"snat_match_destination_port,omitempty"`
}

type EdgeNAT64Rule struct {
	ID                         int    `json:"rule_id"`
	Tag                        int    `json:"rule_tag"`
	LoggingEnabled             bool   `json:"logging_enabled"`
	Enabled                    bool   `json:"enabled"`
	Description                string `json:"description"`
	MatchIPV6DestinationPrefix string `json:"match_ipv6_destination_prefix"`
	TranslatedIPV4SourcePrefix string `json:"translated_ipv4_source_prefix"`
}

// type EdgeNAT struct {
// 	Enabled bool          `json:"enabled"`
// 	Rules   []EdgeNATRule `json:"rules"`
// }

// type EdgeNATRule struct {
// 	ID             int    `json:"id"`
// 	Index          int    `json:"rule_index"`
// 	Description    string `json:"description"`
// 	Type           string `json:"rule_type"`
// 	Enabled        bool   `json:"enabled"`
// 	ICMPSubType    string `json:"icmp_sub_type"`
// 	OriginalIP     string `json:"original_ip"`
// 	OriginalPort   string `json:"original_port"`
// 	Protocol       string `json:"protocol"`
// 	TranslatedIP   string `json:"translated_ip"`
// 	TranslatedPort string `json:"translated_port"`
// 	Interface      string `json:"interface_name"`
// }

// type updateEdgeNAT struct {
// 	Enabled bool                `json:"enabled"`
// 	Rules   []updateEdgeNATRule `json:"rules"`
// }

// type updateEdgeNATRule struct {
// 	ID             int    `json:"id"`
// 	Index          int    `json:"idx"`
// 	Description    string `json:"description,omitempty"`
// 	Type           string `json:"rule_type"`
// 	Enabled        bool   `json:"enabled"`
// 	ICMPSubType    string `json:"icmp_sub_type"`
// 	OriginalIP     string `json:"original_ip"`
// 	OriginalPort   string `json:"original_port,omitempty"`
// 	Protocol       string `json:"protocol,omitempty"`
// 	TranslatedIP   string `json:"translated_ip"`
// 	TranslatedPort string `json:"translated_port,omitempty"`
// 	Interface      string `json:"interface"`
// }

// func (s *edgeService) GetNAT(edgeID string) (EdgeNAT, error) {
// 	nat := EdgeNAT{}
// 	err := s.client.getObject(fmt.Sprintf("/v1/edges/%s/nat", edgeID), &nat)
// 	if err != nil {
// 		return EdgeNAT{}, err
// 	}
// 	return nat, nil
// }

// func (s *edgeService) UpdateNATRules(edgeID string, rules []EdgeNATRule) (Task, error) {
// 	nat, err := s.GetNAT(edgeID)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	updatedRules := []updateEdgeNATRule{}
// 	for i, rule := range rules {
// 		updateRule := updateEdgeNATRule{
// 			ID:             rule.ID,
// 			Index:          i + 1,
// 			Description:    rule.Description,
// 			Type:           rule.Type,
// 			Enabled:        rule.Enabled,
// 			ICMPSubType:    rule.ICMPSubType,
// 			OriginalIP:     rule.OriginalIP,
// 			OriginalPort:   rule.OriginalPort,
// 			Protocol:       rule.Protocol,
// 			TranslatedIP:   rule.TranslatedIP,
// 			TranslatedPort: rule.TranslatedPort,
// 			Interface:      rule.Interface,
// 		}
// 		updatedRules = append(updatedRules, updateRule)
// 	}
// 	updateNAT := updateEdgeNAT{
// 		Enabled: nat.Enabled,
// 		Rules:   updatedRules,
// 	}
// 	data, err := json.Marshal(&updateNAT)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	task := Task{}
// 	err = unmarshalBody(resp, &task)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	return task, nil
// }

// func (s *edgeService) EnableNAT(edgeID string) (Task, error) {
// 	nat, err := s.GetNAT(edgeID)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	nat.Enabled = true
// 	data, err := json.Marshal(&nat)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	task := Task{}
// 	err = unmarshalBody(resp, &task)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	return task, nil
// }

// func (s *edgeService) DisableNAT(edgeID string) (Task, error) {
// 	nat, err := s.GetNAT(edgeID)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	nat.Enabled = false
// 	data, err := json.Marshal(&nat)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	task := Task{}
// 	err = unmarshalBody(resp, &task)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	return task, nil
// }
