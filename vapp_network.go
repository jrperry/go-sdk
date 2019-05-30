package iland

import (
	"encoding/json"
	"fmt"
)

type VAppNetwork struct {
	ID               string    `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Gateway          string    `json:"gateway"`
	Netmask          string    `json:"netmask"`
	IPRanges         []IPRange `json:"ip_ranges"`
	FenceMode        string    `json:"fence_mode"`
	PrimaryDNS       string    `json:"primary_dns"`
	SecondaryDNS     string    `json:"secondary_dns"`
	DNSSuffix        string    `json:"dns_suffix"`
	Inherited        bool      `json:"inherited"`
	Shared           bool      `json:"shared"`
	RouterExternalIP string    `json:"router_external_ip"`
	ParentNetworkID  string    `json:"parent_network_id"`
	VAppID           string    `json:"vapp_uuid"`
	VdcID            string    `json:"vdc_uuid"`
	OrgID            string    `json:"org_uuid"`
	CompanyID        string    `json:"company_id"`
	LocationID       string    `json:"location_id"`
	UpdatedDate      int       `json:"updated_date"`
}

type vappNetworkService struct {
	client *client
}

func (s *vappNetworkService) postAction(vappNetworkID, action string, payload []byte) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/vapp-networks/%s/actions/%s", vappNetworkID, action), payload)
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

func (s *vappNetworkService) Get(vappNetworkID string) (VAppNetwork, error) {
	network := VAppNetwork{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-networks/%s", vappNetworkID), &network)
	if err != nil {
		return VAppNetwork{}, err
	}
	return network, nil
}

type UpdateVAppNetworkParams struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	FenceMode       string    `json:"fence_mode"`
	PrimaryDNS      string    `json:"primary_dns"`
	SecondaryDNS    string    `json:"secondary_dns"`
	ParentNetworkID string    `json:"parent_network_id"`
	IPRanges        []IPRange `json:"ip_ranges"`
}

func (s *vappNetworkService) Update(vappNetworkID string, params UpdateVAppNetworkParams) (Task, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/vapp-networks/%s", vappNetworkID), data)
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

func (s *vappNetworkService) Delete(vappNetworkID string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vapp-networks/%s", vappNetworkID))
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

func (s *vappNetworkService) UpdateDHCP(vappNetworkID string, params DHCP) (Task, error) {
	payload, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappNetworkID, "update-dhcp", payload)
}

type VAppNetworkFirewall struct {
	VAppID         string                    `json:"vapp_uuid"`
	NetworkName    string                    `json:"network_name"`
	Enabled        bool                      `json:"enabled"`
	LoggingEnabled bool                      `json:"logging_enabled"`
	DefaultAction  string                    `json:"default_action"`
	Rules          []VAppNetworkFirewallRule `json:"rules"`
}

type VAppNetworkFirewallRule struct {
	ID                   string   `json:"id"`
	Index                int      `json:"rule_index"`
	Description          string   `json:"description"`
	Enabled              bool     `json:"enabled"`
	LoggingEnabled       bool     `json:"logging_enabled"`
	MatchOnTranslate     bool     `json:"match_on_translate"`
	Policy               string   `json:"policy"`
	Direction            string   `json:"direction"`
	Protocols            []string `json:"protocols"`
	SourceIP             string   `json:"source_ip"`
	SourcePort           int      `json:"source_port"`
	SourcePortRange      string   `json:"source_port_range"`
	DestinationIP        string   `json:"destination_ip"`
	DestinationPort      int      `json:"port"`
	DestinationPortRange string   `json:"destination_port_range"`
}

func (s *vappNetworkService) GetFirewall(vappNetworkID string) (VAppNetworkFirewall, error) {
	firewall := VAppNetworkFirewall{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-networks/%s/firewall", vappNetworkID), &firewall)
	if err != nil {
		return VAppNetworkFirewall{}, err
	}
	return firewall, nil
}

func (s *vappNetworkService) EnableFirewall(vappNetworkID string) (Task, error) {
	fw, err := s.GetFirewall(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	fw.Enabled = true
	return s.updateFirewall(vappNetworkID, fw)
}

func (s *vappNetworkService) DisableFirewall(vappNetworkID string) (Task, error) {
	fw, err := s.GetFirewall(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	fw.Enabled = false
	return s.updateFirewall(vappNetworkID, fw)
}

func (s *vappNetworkService) UpdateFirewallRules(vappNetworkID string, rules []VAppNetworkFirewallRule) (Task, error) {
	firewall, err := s.GetFirewall(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	firewall.Rules = rules
	return s.updateFirewall(vappNetworkID, firewall)
}

func (s *vappNetworkService) updateFirewall(vappNetworkID string, params VAppNetworkFirewall) (Task, error) {
	payload, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappNetworkID, "update-firewall", payload)
}

type VAppNetworkNAT struct {
	VAppID              string               `json:"vapp_uuid"`
	NetworkName         string               `json:"network_name"`
	Enabled             bool                 `json:"enabled"`
	Type                string               `json:"type"`
	IPMasquerade        bool                 `json:"enable_ip_masquerade"`
	IPTranslationRules  []IPTranslationRule  `json:"ip_translation_rules"`
	PortForwardingRules []PortForwardingRule `json:"port_forwarding_rules"`
}

type IPTranslationRule struct {
	MappingMode string `json:"mapping_mode"`
	ExternalIP  string `json:"external_ip"`
	VMInterface string `json:"vm_interface"`
	VMLocalID   string `json:"vm_local_id"`
}

type PortForwardingRule struct {
	ExternalPort  string `json:"external_port"`
	ForwardToPort string `json:"forward_to_port"`
	Protocol      string `json:"protocol"`
	VMInterface   string `json:"vm_interface"`
	VMLocalID     string `json:"vm_local_id"`
}

func (s *vappNetworkService) GetNAT(vappNetworkID string) (VAppNetworkNAT, error) {
	nat := VAppNetworkNAT{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-networks/%s/nat", vappNetworkID), &nat)
	if err != nil {
		return VAppNetworkNAT{}, err
	}
	return nat, nil
}

func (s *vappNetworkService) updateNAT(vappNetworkID string, params VAppNetworkNAT) (Task, error) {
	payload, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappNetworkID, "update-nat", payload)
}

func (s *vappNetworkService) UpdateNATIPTranslationRules(vappNetworkID string, rules []IPTranslationRule) (Task, error) {
	nat, err := s.GetNAT(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	nat.Type = "ipTranslation"
	nat.IPTranslationRules = rules
	return s.updateNAT(vappNetworkID, nat)
}

func (s *vappNetworkService) UpdateNATPortForwardingRules(vappNetworkID string, rules []PortForwardingRule) (Task, error) {
	nat, err := s.GetNAT(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	nat.Type = "portForwarding"
	nat.IPMasquerade = true
	nat.PortForwardingRules = rules
	return s.updateNAT(vappNetworkID, nat)
}

func (s *vappNetworkService) EnableNAT(vappNetworkID string) (Task, error) {
	nat, err := s.GetNAT(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	nat.Enabled = true
	return s.updateNAT(vappNetworkID, nat)
}

func (s *vappNetworkService) DisableNAT(vappNetworkID string) (Task, error) {
	nat, err := s.GetNAT(vappNetworkID)
	if err != nil {
		return Task{}, err
	}
	nat.Enabled = false
	return s.updateNAT(vappNetworkID, nat)
}

type VirtualMachineInterface struct {
	VirtualMachineID      string `json:"vm_uuid"`
	VirtualMachineName    string `json:"vm_name"`
	VirtualMachineLocalID string `json:"vm_local_id"`
	VAppID                string `json:"vapp_uuid"`
	VAppNetworkID         string `json:"vapp_network_uuid"`
	NicID                 int    `json:"nic_id"`
	IPAddress             string `json:"ip_address"`
	IPTranslationMapped   bool   `json:"ip_translation_mapped"`
}

func (s *vappNetworkService) GetInterfaces(vappNetwork string) ([]VirtualMachineInterface, error) {
	schema := struct {
		Interfaces []VirtualMachineInterface `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-networks/%s/vm-interfaces", vappNetwork), &schema)
	if err != nil {
		return []VirtualMachineInterface{}, err
	}
	return schema.Interfaces, nil
}
