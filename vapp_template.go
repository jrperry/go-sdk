package iland

import (
	"encoding/json"
	"fmt"
)

type VAppTemplate struct {
	ID                    string                   `json:"uuid"`
	Name                  string                   `json:"name"`
	Description           string                   `json:"description"`
	Status                int                      `json:"status"`
	SizeGB                float64                  `json:"size"`
	GoldMaster            bool                     `json:"gold_master"`
	IsPublic              bool                     `json:"is_public"`
	Expired               bool                     `json:"expired"`
	Customizable          bool                     `json:"customizable"`
	CustomizationRequired bool                     `json:"customization_required"`
	VirtualMachines       []VirtualMachineTemplate `json:"vm_templates"`
	CatalogID             string                   `json:"catalog_uuid"`
	StorageProfileID      string                   `json:"storage_profile_uuid"`
	VdcID                 string                   `json:"vdc_uuid"`
	OrgID                 string                   `json:"org_uuid"`
	CompanyID             string                   `json:"company_id"`
	LocationID            string                   `json:"location_id"`
	CreatedDate           int                      `json:"created_date"`
	UpdatedDate           int                      `json:"updated_date"`
}

type VirtualMachineTemplate struct {
	ID          string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
	SizeGB      float64 `json:"size"`
}

type vappTemplateService struct {
	client *client
}

func (s *vappTemplateService) Get(vappTemplateID string) (VAppTemplate, error) {
	vappTemplate := VAppTemplate{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-templates/%s", vappTemplateID), &vappTemplate)
	if err != nil {
		return VAppTemplate{}, err
	}
	return vappTemplate, nil
}

type UpdateVAppTemplateParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *vappTemplateService) Update(vappTemplateID string, params UpdateVAppTemplateParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/vapp-templates/%s", vappTemplateID), data)
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

func (s *vappTemplateService) Delete(vappTemplateID string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vapp-templates/%s", vappTemplateID))
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

func (s *vappTemplateService) GetVirtualMachines(vappTemplateID string) ([]VirtualMachineTemplate, error) {
	schema := struct {
		VirtualMachines []VirtualMachineTemplate `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-templates/%s/vms", vappTemplateID), &schema)
	if err != nil {
		return []VirtualMachineTemplate{}, err
	}
	return schema.VirtualMachines, nil
}

type VAppTemplateConfig struct {
	ID              string                         `json:"uuid"`
	Name            string                         `json:"name"`
	Description     string                         `json:"description"`
	VirtualMachines []VirtualMachineTemplateConfig `json:"vms"`
	Networks        []NetworkTemplateConfig        `json:"networks"`
}

type VirtualMachineTemplateConfig struct {
	ID                      string               `json:"uuid"`
	Name                    string               `json:"name"`
	ComputerName            string               `json:"computer_name"`
	Description             string               `json:"description"`
	CPUCount                int                  `json:"number_of_cpus"`
	CoresPerSocket          int                  `json:"number_of_cores_per_socket"`
	MemoryBytes             int                  `json:"memory_in_bytes"`
	OperatingSystem         string               `json:"operating_system_version"`
	HardwareVersion         string               `json:"hardware_version"`
	NestedHypervisorEnabled bool                 `json:"expose_cpu_virtualization"`
	StorageProfileID        string               `json:"storage_profile_uuid"`
	Disks                   []DiskTemplateConfig `json:"disks"`
	Nics                    []NicTemplateConfig  `json:"vnics"`
}

type DiskTemplateConfig struct {
	Name      string `json:"name"`
	SizeBytes int    `json:"size_in_bytes"`
	Type      string `json:"disk_type"`
}

type NicTemplateConfig struct {
	NetworkName        string `json:"network_name"`
	IPAddress          string `json:"ip_address"`
	IPAddressingMode   string `json:"ip_assignment_mode"`
	AdapterType        string `json:"network_adapter_type"`
	IsPrimary          bool   `json:"primary_vnic"`
	IsConnected        bool   `json:"connected"`
	NeedsCustomization bool   `json:"needs_customization"`
}

type NetworkTemplateConfig struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	FenceMode         string  `json:"fence_mode"`
	IPScope           IPScope `json:"ip_scope"`
	ParentNetworkName string  `json:"parent_network_name"`
}

func (s *vappTemplateService) GetConfig(vappTemplateID string) (VAppTemplateConfig, error) {
	config := VAppTemplateConfig{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapp-templates/%s/configuration", vappTemplateID), &config)
	if err != nil {
		return VAppTemplateConfig{}, err
	}
	return config, nil
}

func (s *vappTemplateService) SyncSubscription(vappTemplateID string) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/vapp-templates/%s/actions/sync", vappTemplateID), []byte{})
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
