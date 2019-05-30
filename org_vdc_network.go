package iland

import (
	"encoding/json"
	"fmt"
)

type OrgVdcNetwork struct {
	ID              string    `json:"uuid"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Gateway         string    `json:"gateway"`
	Netmask         string    `json:"netmask"`
	IPRanges        []IPRange `json:"ip_ranges"`
	FenceMode       string    `json:"fence_mode"`
	PrimaryDNS      string    `json:"primary_dns"`
	SecondaryDNS    string    `json:"secondary_dns"`
	DNSSuffix       string    `json:"dns_suffix"`
	Inherited       bool      `json:"inherited"`
	Shared          bool      `json:"shared"`
	ParentNetworkID string    `json:"parent_network_id"`
	EdgeID          string    `json:"edge_uuid"`
	VdcID           string    `json:"vdc_uuid"`
	OrgID           string    `json:"org_uuid"`
	CompanyID       string    `json:"company_id"`
	LocationID      string    `json:"location_id"`
	UpdatedDate     int       `json:"updated_date"`
}

type orgVdcNetworkService struct {
	client *client
}

func (s *orgVdcNetworkService) Get(networkID string) (OrgVdcNetwork, error) {
	network := OrgVdcNetwork{}
	err := s.client.getObject(fmt.Sprintf("/v1/org-vdc-networks/%s", networkID), &network)
	if err != nil {
		return OrgVdcNetwork{}, err
	}
	return network, nil
}

type UpdateOrgVdcNetworkParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Shared      bool   `json:"shared"`
}

func (s *orgVdcNetworkService) Update(networkID string, params UpdateOrgVdcNetworkParams) (Task, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/org-vdc-networks/%s", networkID), data)
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
