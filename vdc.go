package iland

import (
	"encoding/json"
	"fmt"
	"time"
)

type Vdc struct {
	ID                 string `json:"uuid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Enabled            bool   `json:"enabled"`
	AllocationModel    string `json:"allocation_model"`
	ReservedCPU        int    `json:"reserved_cpu"`
	AllocatedCPU       int    `json:"alloc_cpu"`
	ReservedMemory     int    `json:"reserved_mem"`
	AllocatedMemory    int    `json:"allocated_memory"`
	NetworkQuota       int    `json:"network_quota"`
	UsedNetworkCount   int    `json:"used_network_count"`
	MaxHardwareVersion string `json:"max_hardware_version"`
	DiskLimit          int    `json:"disk_limit"`
	AdvancedDiskLimit  int    `json:"contracted_advanced_disk_limit"`
	SSDDiskLimit       int    `json:"contracted_ssd_disk_limit"`
	ArchiveDiskLimit   int    `json:"contracted_archive_disk_limit"`
	CompanyID          string `json:"company_id"`
	OrgID              string `json:"org_uuid"`
	VCenterName        string `json:"vcenter_name"`
	VCloudHref         string `json:"vcloud_href"`
	LocationID         string `json:"location_id"`
	UpdatedDate        int    `json:"updated_date"`
}

type vdcService struct {
	client *client
}

func (s *vdcService) Get(vdcID string) (Vdc, error) {
	vdc := Vdc{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s", vdcID), &vdc)
	if err != nil {
		return Vdc{}, err
	}
	return vdc, nil
}

func (s *vdcService) GetStorageProfiles(vdcID string) ([]StorageProfile, error) {
	schema := struct {
		StorageProfiles []StorageProfile `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/storage-profiles", vdcID), &schema)
	if err != nil {
		return []StorageProfile{}, err
	}
	return schema.StorageProfiles, nil
}

type VdcSummary struct {
	NumberOfVApps    int     `json:"number_of_vapps"`
	NumberOfVms      int     `json:"number_of_vms"`
	ReservedCPU      float64 `json:"reserved_cpu"`
	AllocatedCPU     float64 `json:"allocated_cpu"`
	ConfiguredCPU    float64 `json:"configured_cpu"`
	ConsumedCPU      float64 `json:"consumed_cpu"`
	ReservedMemory   float64 `json:"reserved_mem"`
	AllocatedMemory  float64 `json:"allocated_mem"`
	ConfiguredMemory float64 `json:"configured_mem"`
	ConsumedMemory   float64 `json:"consumed_mem"`
	ProvisionedDisk  float64 `json:"provisioned_disk"`
	ConfiguredDisk   float64 `json:"configured_disk"`
	ConsumedDisk     float64 `json:"consumed_disk"`
}

func (s *vdcService) GetSummary(vdcID string) (VdcSummary, error) {
	summary := VdcSummary{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/summary", vdcID), &summary)
	if err != nil {
		return VdcSummary{}, err
	}
	return summary, nil
}

func (s *vdcService) GetVApps(vdcID string) ([]VApp, error) {
	schema := struct {
		VApps []VApp `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/vapps", vdcID), &schema)
	if err != nil {
		return []VApp{}, err
	}
	return schema.VApps, nil
}

func (s *vdcService) GetVirtualMachines(vdcID string) ([]VirtualMachine, error) {
	schema := struct {
		VirtualMachines []VirtualMachine `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/vms", vdcID), &schema)
	if err != nil {
		return []VirtualMachine{}, err
	}
	return schema.VirtualMachines, nil
}

func (s *vdcService) GetEdges(vdcID string) ([]Edge, error) {
	schema := struct {
		Edges []Edge `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/edges", vdcID), &schema)
	if err != nil {
		return []Edge{}, err
	}
	return schema.Edges, nil
}

func (s *vdcService) GetNetworks(vdcID string) ([]OrgVdcNetwork, error) {
	schema := struct {
		Networks []OrgVdcNetwork `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/org-vdc-networks", vdcID), &schema)
	if err != nil {
		return []OrgVdcNetwork{}, err
	}
	return schema.Networks, nil
}

func (s *vdcService) GetCurrentBill(vdcID string) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/billing", vdcID), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *vdcService) GetBill(vdcID string, month, year int) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/billing?month=%d&year=%d", vdcID, month, year), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *vdcService) GetCurrentVAppBill(vdcID string) ([]Billing, error) {
	schema := struct {
		Billing []Billing `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/vapp-bills", vdcID), &schema)
	if err != nil {
		return []Billing{}, err
	}
	return schema.Billing, nil
}

func (s *vdcService) GetVAppBill(vdcID string, month, year int) ([]Billing, error) {
	schema := struct {
		Billing []Billing `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/vapp-bills?month=%d&year=%d", vdcID, month, year), &schema)
	if err != nil {
		return []Billing{}, err
	}
	return schema.Billing, nil
}

func (s *vdcService) GetPerformanceCounters(vdcID string) ([]PerformanceCounter, error) {
	schema := struct {
		Counters []PerformanceCounter `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/performance-counters", vdcID), &schema)
	if err != nil {
		return []PerformanceCounter{}, err
	}
	return schema.Counters, nil
}

func (s *vdcService) GetPerformance(vdcID string, counter PerformanceCounter, start, end time.Time) (Performance, error) {
	startNano := getUnixMilliseconds(start)
	endNano := getUnixMilliseconds(end)
	performance := Performance{}
	err := s.client.getObject(fmt.Sprintf("/v1/vdcs/%s/performance/%s::%s::%s?start=%d&end=%d", vdcID, counter.Group, counter.Name, counter.Type, startNano, endNano), &performance)
	if err != nil {
		return Performance{}, err
	}
	return performance, nil
}

type BuildVAppParams struct {
	Name            string                      `json:"name"`
	Description     string                      `json:"description"`
	VirtualMachines []BuildVirtualMachineParams `json:"vms"`
}

func (s *vdcService) BuildVApp(vdcID string, params BuildVAppParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/vdcs/%s/actions/build-vapp", vdcID), data)
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

type DeployVAppTemplateParams struct {
	Name             string                                   `json:"name"`
	Description      string                                   `json:"description"`
	VAppTemplateID   string                                   `json:"vapp_template_uuid"`
	PreserveNetworks bool                                     `json:"preserve_networks"`
	VirtualMachines  []DeployVAppTemplateVirtualMachineParams `json:"vms"`
}

type DeployVAppTemplateVirtualMachineParams struct {
	Name                     string                  `json:"name"`
	Description              string                  `json:"description,omitempty"`
	VirtualMachineTemplateID string                  `json:"vm_template_uuid"`
	StorageProfileID         string                  `json:"storage_profile_uuid,omitempty"`
	Nics                     []VAppTemplateNicParams `json:"vnics"`
}

type VAppTemplateNicParams struct {
	NetworkID          string `json:"network_uuid,omitempty"`
	IPAssignment       string `json:"ip_assignment"`
	IPAddress          string `json:"ip_address,omitempty"`
	PrimaryVNic        bool   `json:"primary_vnic"`
	NetworkAdapterType string `json:"network_adapter_type"`
}

func (s *vdcService) DeployVAppTemplate(vdcID string, params DeployVAppTemplateParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/vdcs/%s/actions/add-vapp-from-template", vdcID), data)
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
