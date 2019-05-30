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

type EdgeFirewall struct {
	Log           bool               `json:"log"`
	Enabled       bool               `json:"enabled"`
	DefaultAction string             `json:"default_action"`
	Rules         []EdgeFirewallRule `json:"rules"`
}

type EdgeFirewallRule struct {
	ID                   string   `json:"id"`
	Order                int      `json:"idx"`
	Description          string   `json:"description"`
	Enabled              bool     `json:"enabled"`
	Policy               string   `json:"policy"`
	SourceIP             string   `json:"source_ip,omitempty"`
	SourcePort           int      `json:"source_port,omitempty"`
	SourcePortRange      string   `json:"source_port_range,omitempty"`
	DestinationIP        string   `json:"destination_ip,omitempty"`
	DestinationPort      int      `json:"port,omitempty"`
	DestinationPortRange string   `json:"destination_port_range,omitempty"`
	Direction            string   `json:"direction,omitempty"`
	Protocols            []string `json:"protocols"`
	MatchOnTranslate     bool     `json:"match_on_translate,omitempty"`
}

func (s *edgeService) GetFirewall(edgeID string) (EdgeFirewall, error) {
	firewall := EdgeFirewall{}
	err := s.client.getObject(fmt.Sprintf("/v1/edges/%s/firewall", edgeID), &firewall)
	if err != nil {
		return EdgeFirewall{}, err
	}
	return firewall, nil
}

func (s *edgeService) UpdateFirewallRules(edgeID string, rules []EdgeFirewallRule) (Task, error) {
	firewall, err := s.GetFirewall(edgeID)
	if err != nil {
		return Task{}, err
	}
	firewall.Rules = rules
	data, err := json.Marshal(&firewall)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-firewall", edgeID), data)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	err = unmarshalBody(resp, &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

type EdgeNAT struct {
	Enabled bool          `json:"enabled"`
	Rules   []EdgeNATRule `json:"rules"`
}

type EdgeNATRule struct {
	ID             int    `json:"id"`
	Order          int    `json:"idx"`
	Description    string `json:"description"`
	Enabled        bool   `json:"enabled"`
	Type           string `json:"rule_type"`
	Interface      string `json:"interface"`
	OriginalIP     string `json:"original_ip"`
	OriginalPort   string `json:"original_port"`
	TranslatedIP   string `json:"translated_ip"`
	TranslatedPort string `json:"translated_port"`
	Protocol       string `json:"protocol"`
}

func (s *edgeService) GetNAT(edgeID string) (EdgeNAT, error) {
	nat := EdgeNAT{}
	err := s.client.getObject(fmt.Sprintf("/v1/edges/%s/nat", edgeID), &nat)
	if err != nil {
		return EdgeNAT{}, err
	}
	return nat, nil
}

func (s *edgeService) UpdateNATRules(edgeID string, rules []EdgeNATRule) (Task, error) {
	nat, err := s.GetNAT(edgeID)
	if err != nil {
		return Task{}, err
	}
	for i, rule := range rules {
		rule.Order = i + 1
		rules[i] = rule
	}
	nat.Rules = rules
	data, err := json.Marshal(&nat)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	err = unmarshalBody(resp, &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *edgeService) EnableNAT(edgeID string) (Task, error) {
	nat, err := s.GetNAT(edgeID)
	if err != nil {
		return Task{}, err
	}
	nat.Enabled = true
	data, err := json.Marshal(&nat)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	err = unmarshalBody(resp, &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *edgeService) DisableNAT(edgeID string) (Task, error) {
	nat, err := s.GetNAT(edgeID)
	if err != nil {
		return Task{}, err
	}
	nat.Enabled = false
	data, err := json.Marshal(&nat)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/edges/%s/actions/update-nat", edgeID), data)
	if err != nil {
		return Task{}, err
	}
	task := Task{}
	err = unmarshalBody(resp, &task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
