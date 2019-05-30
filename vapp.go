package iland

import (
	"encoding/json"
	"fmt"
	"time"
)

type VApp struct {
	ID                string   `json:"uuid"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Status            string   `json:"status"`
	Deployed          bool     `json:"deployed"`
	IsExpired         bool     `json:"is_expired"`
	AllocationModel   string   `json:"allocation_model"`
	StorageProfileIDs []string `json:"storage_profiles"`
	VdcID             string   `json:"vdc_uuid"`
	OrgID             string   `json:"org_uuid"`
	CompanyID         string   `json:"company_id"`
	LocationID        string   `json:"location_id"`
	CreatedDate       int      `json:"created_date"`
	UpdatedDate       int      `json:"updated_date"`
}

type vappService struct {
	client *client
}

func (s *vappService) postAction(vappID, action string, params []byte) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/vapps/%s/actions/%s", vappID, action), params)
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

func (s *vappService) Get(vappID string) (VApp, error) {
	vapp := VApp{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s", vappID), &vapp)
	if err != nil {
		return VApp{}, err
	}
	return vapp, nil
}

func (s *vappService) Delete(vappID string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vapps/%s", vappID))
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

func (s *vappService) GetVirtualMachines(vappID string) ([]VirtualMachine, error) {
	schema := struct {
		VirtualMachines []VirtualMachine `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/vms", vappID), &schema)
	if err != nil {
		return []VirtualMachine{}, err
	}
	return schema.VirtualMachines, nil
}

func (s *vappService) GetNetworks(vappID string) ([]VAppNetwork, error) {
	schema := struct {
		Networks []VAppNetwork `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/networks", vappID), &schema)
	if err != nil {
		return []VAppNetwork{}, err
	}
	return schema.Networks, nil
}

func (s *vappService) AddOrgNetwork(vappID, orgVdcNetworkID string) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/vapps/%s/org-vdc-network/%s", vappID, orgVdcNetworkID), []byte{})
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

func (s *vappService) UpdateName(vappID, name string) (Task, error) {
	params := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "update-name", data)
}

func (s *vappService) UpdateDescription(vappID, description string) (Task, error) {
	params := struct {
		Description string `json:"description"`
	}{
		Description: description,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "update-name", data)
}

type CopyVAppParams struct {
	Name  string `json:"name"`
	VdcID string `json:"vdc_uuid"`
}

func (s *vappService) Copy(vappID string, params CopyVAppParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "copy", data)
}

type MoveVAppParams struct {
	Name  string `json:"name"`
	VdcID string `json:"vdc_uuid"`
}

func (s *vappService) Move(vappID string, params MoveVAppParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "move", data)
}

type BuildVirtualMachineParams struct {
	Name                     string           `json:"name"`
	Description              string           `json:"description"`
	ComputerName             string           `json:"computer_name"`
	VAppTemplateID           string           `json:"vapp_template_uuid"`
	VirtualMachineTemplateID string           `json:"vm_template_uuid"`
	StorageProfileID         string           `json:"storage_profile_uuid"`
	CPUCount                 int              `json:"number_of_cpus,omitempty"`
	CPUCoresPerSocket        int              `json:"cpu_cores_per_socket,omitempty"`
	EnableCPUVirtualization  bool             `json:"expose_cpu_virtualization"`
	MemoryMB                 int              `json:"ram,omitempty"`
	HardwareVersion          int              `json:"hardware_version,omitempty"`
	BootDelay                int              `json:"boot_delay,omitempty"`
	Disks                    []Disk           `json:"disks"`
	Nics                     []BuildNicParams `json:"nics"`
}

type BuildNicParams struct {
	IPAssignment string `json:"ip_assignment,omitempty"`
	IPAddress    string `json:"ip_address,omitempty"`
	Primary      bool   `json:"primary_vnic"`
	Type         string `json:"network_adapter_type,omitempty"`
	NetworkID    string `json:"network_uuid,omitempty"`
}

func (s *vappService) BuildVirtualMachines(vappID string, params []BuildVirtualMachineParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "build-vms", data)
}

type AddTemplateVirtualMachineParams struct {
	Name                     string `json:"name"`
	Description              string `json:"description"`
	NetworkID                string `json:"network_uuid,omitempty"`
	VAppTemplateID           string `json:"vapp_template_uuid"`
	TemplateVirtualMachineID string `json:"vm_template_uuid"`
	IPAddress                string `json:"ip_address,omitempty"`
	StorageProfileID         string `json:"storage_profile_uuid,omitempty"`
	IPAddressingMode         string `json:"ip_address_mode,omitempty"`
}

func (s *vappService) AddTemplateVirtualMachines(vappID string, params []AddTemplateVirtualMachineParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "add-vms-from-templates", data)
}

type CreateVAppNetworkParams struct {
	Name            string    `json:"name"`
	Description     string    `json:"description,omitempty"`
	ParentNetworkID string    `json:"parent_network_uuid,omitempty"`
	Gateway         string    `json:"gateway_address"`
	Netmask         string    `json:"network_mask"`
	IPRanges        []IPRange `json:"ip_ranges"`
	PrimaryDNS      string    `json:"primary_dns,omitempty"`
	SecondaryDNS    string    `json:"secondary_dns,omitempty"`
	DNSSuffix       string    `json:"dns_suffix,omitempty"`
}

func (s *vappService) CreateNetwork(vappID string, params CreateVAppNetworkParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "add-vapp-network", data)
}

func (s *vappService) PowerOn(vappID string) (Task, error) {
	return s.postAction(vappID, "poweron", []byte{})
}

func (s *vappService) PowerOff(vappID string) (Task, error) {
	return s.postAction(vappID, "poweroff", []byte{})
}

func (s *vappService) Shutdown(vappID string) (Task, error) {
	return s.postAction(vappID, "shutdown", []byte{})
}

func (s *vappService) Reboot(vappID string) (Task, error) {
	return s.postAction(vappID, "reboot", []byte{})
}

func (s *vappService) Reset(vappID string) (Task, error) {
	return s.postAction(vappID, "reset", []byte{})
}

func (s *vappService) Suspend(vappID string) (Task, error) {
	return s.postAction(vappID, "suspend", []byte{})
}

func (s *vappService) GetCurrentBill(vappID string) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/billing", vappID), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *vappService) GetBill(vappID string, month, year int) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/billing?month=%d&year=%d", vappID, month, year), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *vappService) GetAvailableStorageProfiles(vappID string) ([]StorageProfile, error) {
	schema := struct {
		StorageProfiles []StorageProfile `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/available-storage-profiles", vappID), &schema)
	if err != nil {
		return []StorageProfile{}, err
	}
	return schema.StorageProfiles, nil
}

func (s *vappService) GetMetadata(vappID string) ([]Metadata, error) {
	schema := struct {
		Metadata []Metadata `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/metadata", vappID), &schema)
	if err != nil {
		return []Metadata{}, err
	}
	return schema.Metadata, nil
}

func (s *vappService) UpdateMetadata(vappID string, metadata []Metadata) (Task, error) {
	payload, err := json.Marshal(&metadata)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/vapps/%s/metadata", vappID), payload)
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

func (s *vappService) DeleteMetadata(vappID, metadataKey string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vapps/%s/metadata/%s", vappID, metadataKey))
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

func (s *vappService) HasSnapshot(vappID string) (bool, error) {
	schema := struct {
		HasSnapshot bool `json:"has_snapshot"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/has-snapshot", vappID), &schema)
	if err != nil {
		return false, err
	}
	return schema.HasSnapshot, nil
}

func (s *vappService) GetSnapshot(vappID string) (Snapshot, error) {
	snapshot := Snapshot{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/snapshot", vappID), &snapshot)
	if err != nil {
		return Snapshot{}, err
	}
	return snapshot, nil
}

func (s *vappService) CreateSnapshot(vappID string) (Task, error) {
	params := struct {
		Memory  bool `json:"memory"`
		Quiesce bool `json:"quiesce"`
	}{
		Memory:  false,
		Quiesce: false,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "create-snapshot", data)
}

func (s *vappService) RestoreSnapshot(vappID string) (Task, error) {
	return s.postAction(vappID, "restore-snapshot", []byte{})
}

func (s *vappService) RemoveSnapshot(vappID string) (Task, error) {
	return s.postAction(vappID, "remove-snapshot", []byte{})
}

type VAppStartupSetting struct {
	VirtualMachineName string `json:"vm_name"`
	Order              int    `json:"ord"`
	StartAction        string `json:"startup_action"`
	StopAction         string `json:"stop_action"`
	StartDelay         int    `json:"start_delay"`
	StopDelay          int    `json:"stop_delay"`
}

func (s *vappService) GetStartupSettings(vappID string) ([]VAppStartupSetting, error) {
	schema := struct {
		Settings []VAppStartupSetting `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/startup-section", vappID), &schema)
	if err != nil {
		return []VAppStartupSetting{}, err
	}
	return schema.Settings, nil
}

func (s *vappService) UpdateStartupSettings(vappID string, params []VAppStartupSetting) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(vappID, "update-startup-section", data)
}

func (s *vappService) GetPerformanceCounters(vappID string) ([]PerformanceCounter, error) {
	schema := struct {
		Counters []PerformanceCounter `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/performance-counters", vappID), &schema)
	if err != nil {
		return []PerformanceCounter{}, err
	}
	return schema.Counters, nil
}

func (s *vappService) GetPerformance(vappID string, counter PerformanceCounter, start, end time.Time) (Performance, error) {
	startNano := getUnixMilliseconds(start)
	endNano := getUnixMilliseconds(end)
	performance := Performance{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/performance/%s::%s::%s?start=%d&end=%d", vappID, counter.Group, counter.Name, counter.Type, startNano, endNano), &performance)
	if err != nil {
		return Performance{}, err
	}
	return performance, nil
}

type VAppSummary struct {
	NumberOfVms     int     `json:"number_of_vms"`
	ReservedCPU     float64 `json:"reserved_cpu"`
	ConsumedCPU     float64 `json:"consumed_cpu"`
	ReservedMemory  float64 `json:"reserved_mem"`
	ConsumedMemory  float64 `json:"consumed_mem"`
	ProvisionedDisk float64 `json:"provisioned_disk"`
	ConsumedDisk    float64 `json:"consumed_disk"`
}

func (s *vappService) GetSummary(vdcID string) (VAppSummary, error) {
	summary := VAppSummary{}
	err := s.client.getObject(fmt.Sprintf("/v1/vapps/%s/summary", vdcID), &summary)
	if err != nil {
		return VAppSummary{}, err
	}
	return summary, nil
}
