package iland

import (
	"fmt"
)

type Org struct {
	ID                                string `json:"uuid"`
	Name                              string `json:"name"`
	FullName                          string `json:"fullname"`
	Description                       string `json:"description"`
	Enabled                           bool   `json:"enabled"`
	VAppMaxRuntimeLease               int    `json:"vapp_max_runtime_lease"`
	VAppMaxStorageLease               int    `json:"vapp_max_storage_lease"`
	VAppDeleteOnStorageExpire         bool   `json:"vapp_delete_on_storage_expire"`
	VAppTemplateDeleteOnStorageExpire bool   `json:"vapp_template_delete_on_storage_expire"`
	ZertoTarget                       bool   `json:"zerto_target"`
	LocationID                        string `json:"location_id"`
	CompanyID                         string `json:"company_id"`
	UpdatedDate                       int    `json:"updated_date"`
}

type orgService struct {
	client *client
}

func (s *orgService) Get(orgID string) (Org, error) {
	org := Org{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s", orgID), &org)
	if err != nil {
		return Org{}, err
	}
	return org, nil
}

func (s *orgService) GetVdcs(orgID string) ([]Vdc, error) {
	schema := struct {
		Vdcs []Vdc `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/vdcs", orgID), &schema)
	if err != nil {
		return []Vdc{}, err
	}
	return schema.Vdcs, nil
}

func (s *orgService) GetEdges(orgID string) ([]Edge, error) {
	schema := struct {
		Edges []Edge `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/edges", orgID), &schema)
	if err != nil {
		return []Edge{}, err
	}
	return schema.Edges, nil
}

func (s *orgService) GetCatalogs(orgID string) ([]Catalog, error) {
	schema := struct {
		Catalogs []Catalog `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/catalogs", orgID), &schema)
	if err != nil {
		return []Catalog{}, err
	}
	return schema.Catalogs, nil
}

func (s *orgService) GetVAppTemplates(orgID string) ([]VAppTemplate, error) {
	schema := struct {
		VAppTemplates []VAppTemplate `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/vapp-templates", orgID), &schema)
	if err != nil {
		return []VAppTemplate{}, err
	}
	return schema.VAppTemplates, nil
}

func (s *orgService) GetMedia(orgID string) ([]Media, error) {
	schema := struct {
		Media []Media `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/medias", orgID), &schema)
	if err != nil {
		return []Media{}, err
	}
	return schema.Media, nil
}

func (s *orgService) GetNetworks(orgID string) ([]OrgVdcNetwork, error) {
	schema := struct {
		Networks []OrgVdcNetwork `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/org-vdc-networks", orgID), &schema)
	if err != nil {
		return []OrgVdcNetwork{}, err
	}
	return schema.Networks, nil
}

func (s *orgService) GetVApps(orgID string) ([]VApp, error) {
	schema := struct {
		VApps []VApp `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/vapps", orgID), &schema)
	if err != nil {
		return []VApp{}, err
	}
	return schema.VApps, nil
}

func (s *orgService) GetVirtualMachines(orgID string) ([]VirtualMachine, error) {
	schema := struct {
		VirtualMachines []VirtualMachine `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/vms", orgID), &schema)
	if err != nil {
		return []VirtualMachine{}, err
	}
	return schema.VirtualMachines, nil
}

func (s *orgService) GetVpgs(orgID string) ([]Vpg, error) {
	schema := struct {
		Vpgs []Vpg `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/vpgs?expand=VPG_VM", orgID), &schema)
	if err != nil {
		return []Vpg{}, err
	}
	return schema.Vpgs, nil
}

func (s *orgService) GetPublicIPs(orgID string) ([]string, error) {
	schema := struct {
		IPs []string `json:"ips"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/public-ips", orgID), &schema)
	if err != nil {
		return []string{}, err
	}
	return schema.IPs, nil
}

type PublicIPAssignment struct {
	IP                  string `json:"ip"`
	Type                string `json:"type"`
	EntityID            string `json:"entity_uuid"`
	EntityName          string `json:"entity_name"`
	ExternalNetworkID   string `json:"external_network_uuid"`
	ExternalNetworkName string `json:"external_network_name"`
}

func (s *orgService) GetPublicIPAssignments(orgID string) ([]PublicIPAssignment, error) {
	schema := struct {
		Assignments []PublicIPAssignment `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/public-ip-assignments", orgID), &schema)
	if err != nil {
		return []PublicIPAssignment{}, err
	}
	return schema.Assignments, nil
}

func (s *orgService) GetCurrentBill(vdcID string) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/billing", vdcID), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *orgService) GetBill(vdcID string, month, year int) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/orgs/%s/billing?month=%d&year=%d", vdcID, month, year), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}
