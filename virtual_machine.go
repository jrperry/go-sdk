package iland

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type VirtualMachine struct {
	ID                         string   `json:"uuid"`
	Name                       string   `json:"name"`
	Description                string   `json:"description"`
	Status                     string   `json:"status"`
	Deployed                   bool     `json:"deployed"`
	LocalID                    string   `json:"vm_local_id"`
	OperatingSystemName        string   `json:"os"`
	OperatingSystemDescription string   `json:"os_description"`
	CPUCount                   int      `json:"cpus_number"`
	CoresPerSocket             int      `json:"cores_per_socket"`
	MemoryMB                   int      `json:"memory_size"`
	StorageProfileIDs          []string `json:"storage_profiles"`
	HardwareVersion            string   `json:"hardware_version"`
	MediaInserted              bool     `json:"media_inserted"`
	MediaName                  string   `json:"inserted_media_name"`
	NestedHypervisorEnabled    bool     `json:"nested_hypervisor_enabled"`
	AllocationModel            string   `json:"allocation_model"`
	VCloudHref                 string   `json:"vcloud_href"`
	VCenterMoref               string   `json:"vcenter_moref"`
	VCenterName                string   `json:"vcenter_name"`
	VCenterInstanceUUID        string   `json:"vcenter_instance_uuid"`
	VAppID                     string   `json:"vapp_uuid"`
	VdcID                      string   `json:"vdc_uuid"`
	OrgID                      string   `json:"org_uuid"`
	CompanyID                  string   `json:"company_id"`
	LocationID                 string   `json:"location_id"`
	CreatedDate                int      `json:"created_date"`
	UpdatedDate                int      `json:"updated_date"`
}

type virtualMachineService struct {
	client *client
}

func (s *virtualMachineService) postAction(virtualMachineID, action string, params []byte) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/vms/%s/actions/%s", virtualMachineID, action), params)
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

func (s *virtualMachineService) Get(virtualMachineID string) (VirtualMachine, error) {
	virtualMachine := VirtualMachine{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s", virtualMachineID), &virtualMachine)
	if err != nil {
		return VirtualMachine{}, err
	}
	return virtualMachine, nil
}

func (s *virtualMachineService) Delete(virtualMachineID string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vms/%s", virtualMachineID))
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

func (s *virtualMachineService) UpdateName(virtualMachineID, name string) (Task, error) {
	params := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-name", data)
}

func (s *virtualMachineService) UpdateDescription(virtualMachineID, description string) (Task, error) {
	params := struct {
		Description string `json:"description"`
	}{
		Description: description,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-description", data)
}

func (s *virtualMachineService) PowerOn(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "poweron", []byte{})
}

func (s *virtualMachineService) PowerOnForceCustomization(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "poweron?forceGuestCustomization=true", []byte{})
}

func (s *virtualMachineService) PowerOff(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "poweroff", []byte{})
}

func (s *virtualMachineService) Reboot(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "reboot", []byte{})
}

func (s *virtualMachineService) Reset(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "reset", []byte{})
}

func (s *virtualMachineService) Shutdown(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "shutdown", []byte{})
}

func (s *virtualMachineService) Suspend(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "suspend", []byte{})
}

type CopyVirtualMachineParams struct {
	Name   string `json:"name"`
	VAppID string `json:"vapp_uuid"`
}

func (s *virtualMachineService) Copy(virtualMachineID string, params CopyVirtualMachineParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "copy", data)
}

type MoveVirtualMachineParams struct {
	Name   string `json:"name"`
	VAppID string `json:"vapp_uuid"`
}

func (s *virtualMachineService) Move(virtualMachineID string, params MoveVirtualMachineParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "move", data)
}

type Summary struct {
	ReservedCpu     float64 `json:"reserved_cpu"`
	ConsumedCpu     float64 `json:"consumed_cpu"`
	ReservedMemory  float64 `json:"reserved_mem"`
	ConsumedMemory  float64 `json:"consumed_mem"`
	ConfiguredDisk  float64 `json:"configured_disk"`
	ProvisionedDisk float64 `json:"provisioned_disk"`
}

func (s *virtualMachineService) GetSummary(virtualMachineID string) (Summary, error) {
	summary := Summary{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/summary", virtualMachineID), &summary)
	if err != nil {
		return Summary{}, err
	}
	return summary, nil
}

func (s *virtualMachineService) GetAvailableStorageProfiles(virtualMachineID string) ([]StorageProfile, error) {
	schema := struct {
		StorageProfiles []StorageProfile `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/available-storage-profiles", virtualMachineID), &schema)
	if err != nil {
		return []StorageProfile{}, err
	}
	return schema.StorageProfiles, nil
}

func (s *virtualMachineService) ChangeStorageProfile(virtualMachineID, storageProfileID string) (Task, error) {
	schema := struct {
		StorageProfileID string `json:"storage_profile"`
	}{
		StorageProfileID: storageProfileID,
	}
	data, err := json.Marshal(&schema)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "relocate", data)
}

func (s *virtualMachineService) EnableNestedHypervisor(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "enable-nested-hypervisor", []byte{})
}

func (s *virtualMachineService) DisableNestedHypervisor(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "disable-nested-hypervisor", []byte{})
}

func (s *virtualMachineService) InsertMedia(virtualMachineID, mediaID string) (Task, error) {
	schema := struct {
		MediaID string `json:"media"`
	}{
		MediaID: mediaID,
	}
	data, err := json.Marshal(&schema)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "insert-media", data)
}

func (s *virtualMachineService) EjectMedia(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "eject-media", []byte{})
}

type GuestCustomization struct {
	Enabled               bool   `json:"enabled"`
	Required              bool   `json:"required"`
	ComputerName          string `json:"computer_name"`
	VirtualMachineLocalID string `json:"virtual_machine_id"`
	ChangeSid             bool   `json:"change_sid"`
	AdminPasswordEnabled  bool   `json:"admin_password_enabled"`
	GenerateAdminPassword bool   `json:"admin_password_auto"`
	AdminPassword         string `json:"admin_password"`
	AdminAutoLoginEnabled bool   `json:"admin_auto_logon_enabled"`
	AdminAutoLogonCount   int    `json:"admin_auto_logon_count"`
	ResetPasswordRequired bool   `json:"reset_password_required"`
	UseOrgSettings        bool   `json:"use_org_settings"`
	JoinDomain            bool   `json:"join_domain"`
	DomainName            string `json:"domain_name"`
	DomainUserName        string `json:"domain_user_name"`
	DomanUserPassword     string `json:"domain_user_password"`
	MachineObjectOU       string `json:"machine_object_ou"`
}

func (s *virtualMachineService) GetGuestCustomization(virtualMachineID string) (GuestCustomization, error) {
	guestCustomization := GuestCustomization{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/guest-customization", virtualMachineID), &guestCustomization)
	if err != nil {
		return GuestCustomization{}, err
	}
	return guestCustomization, nil
}

func (s *virtualMachineService) UpdateGuestCustomization(virtualMachineID string, params GuestCustomization) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-guest-customization", data)
}

type HotAdd struct {
	CPUEnabled    bool `json:"cpu_hot_add_enabled"`
	MemoryEnabled bool `json:"memory_hot_add_enabled"`
}

func (s *virtualMachineService) GetHotAdd(virtualMachineID string) (HotAdd, error) {
	hotAdd := HotAdd{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/capabilities", virtualMachineID), &hotAdd)
	if err != nil {
		return HotAdd{}, err
	}
	return hotAdd, nil
}

func (s *virtualMachineService) UpdateHotAdd(virtualMachineID string, params HotAdd) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-capabilities", data)
}

type BootOptions struct {
	BootDelay int  `json:"boot_delay"`
	EnterBios bool `json:"is_enter_bios"`
}

func (s *virtualMachineService) GetBootOptions(virtualMachineID string) (BootOptions, error) {
	bootOptions := BootOptions{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/boot-options", virtualMachineID), &bootOptions)
	if err != nil {
		return BootOptions{}, err
	}
	return bootOptions, nil
}

func (s *virtualMachineService) UpdateBootOptions(virtualMachineID string, params BootOptions) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-boot-options", data)
}

func (s *virtualMachineService) UpdateHardwareVersion(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "update-virtual-hardware-version", []byte{})
}

type VMwareTools struct {
	Status        string `json:"status"`
	RunningStatus string `json:"running_status"`
	Version       string `json:"version"`
}

func (s *virtualMachineService) GetVMwareTools(virtualMachineID string) (VMwareTools, error) {
	tools := VMwareTools{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/guest-tools", virtualMachineID), &tools)
	if err != nil {
		return VMwareTools{}, err
	}
	return tools, nil
}

func (s *virtualMachineService) UpgradeVMwareTools(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "upgrade-guest-tools", []byte{})
}

func (s *virtualMachineService) InstallVMwareTools(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "install-vmware-tools", []byte{})
}

type ReconfigureParams struct {
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	GuestCustomization GuestCustomization `json:"guest_customization_section"`
	Cpu                UpdateCPUParams    `json:"cpu_spec"`
	Disks              []DiskParams       `json:"disk_spec"`
	Memory             UpdateMemoryParams `json:"memory_spec"`
}

func (s *virtualMachineService) Reconfigure(virtualMachineID string, params ReconfigureParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "reconfigure", data)
}

type Disk struct {
	Name             string `json:"name,omitempty"`
	SizeMB           int    `json:"size"`
	Type             string `json:"type,omitempty"`
	StorageProfileID string `json:"storage_profile_uuid,omitempty"`
}

func (s *virtualMachineService) GetDisks(virtualMachineID string) ([]Disk, error) {
	schema := struct {
		Disks []Disk `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/virtual-disks", virtualMachineID), &schema)
	if err != nil {
		return []Disk{}, err
	}
	return schema.Disks, nil
}

type DiskParams struct {
	Name             string `json:"name,omitempty"`
	Type             string `json:"type"`
	SizeMB           int    `json:"size"`
	StorageProfileID string `json:"storage_profile_uuid"`
}

func (s *virtualMachineService) AddDisk(virtualMachineID string, params DiskParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "add-virtual-disk", data)
}

func (s *virtualMachineService) UpdateDisk(virtualMachineID string, params DiskParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-virtual-disk", data)
}

func (s *virtualMachineService) UpdateDisks(virtualMachineID string, params []DiskParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-virtual-disks", data)
}

func (s *virtualMachineService) DeleteDisk(virtualMachineID string, diskName string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vms/%s/disks/%s", virtualMachineID, diskName))
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

func (s *virtualMachineService) GetRecommendedBusType(virtualMachineID string) (string, error) {
	schema := struct {
		BusType string `json:"bus_type"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/recommended-disk-bus-type", virtualMachineID), &schema)
	if err != nil {
		return "", err
	}
	return schema.BusType, nil
}

type Nic struct {
	ID               int    `json:"vnic_id"`
	IPAddress        string `json:"ip_address"`
	IPAddressingMode string `json:"ip_addressing_mode"`
	MacAddress       string `json:"mac_address,omitempty"`
	AdapterType      string `json:"adapter_type"`
	NetworkName      string `json:"network_name"`
	IsConnected      bool   `json:"is_connected"`
	IsPrimary        bool   `json:"is_primary"`
}

func (s *virtualMachineService) GetNics(virtualMachineID string) ([]Nic, error) {
	schema := struct {
		Nics []Nic `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/vnics", virtualMachineID), &schema)
	if err != nil {
		return []Nic{}, err
	}
	return schema.Nics, nil
}

func (s *virtualMachineService) DeleteNic(virtualMachineID string, nicID int) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vms/%s/vnics/%d", virtualMachineID, nicID))
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

func (s *virtualMachineService) UpdateNics(virtualMachineID string, params []Nic) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	fmt.Println(string(data))
	return s.postAction(virtualMachineID, "update-vnics", data)
}

type UpdateCPUParams struct {
	CPUCount       int `json:"cpus_number"`
	CoresPerSocket int `json:"cores_per_socket"`
}

func (s *virtualMachineService) UpdateCPU(virtualMachineID string, params UpdateCPUParams) (Task, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-cpu-count", data)
}

func (s *virtualMachineService) UpdateCPUCount(virtualMachineID string, cpuCount int) (Task, error) {
	params := UpdateCPUParams{
		CPUCount:       cpuCount,
		CoresPerSocket: 1,
	}
	return s.UpdateCPU(virtualMachineID, params)
}

type UpdateMemoryParams struct {
	MemoryMB int `json:"memory_size"`
}

func (s *virtualMachineService) UpdateMemory(virtualMachineID string, memorySize int) (Task, error) {
	params := UpdateMemoryParams{
		MemoryMB: memorySize,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "update-memory-size", data)
}

type VirtualMachineBackup struct {
	Name      string `json:"name"`
	Timestamp int    `json:"timestamp"`
}

func (s *virtualMachineService) GetBackups(virtualMachineID string) ([]VirtualMachineBackup, error) {
	schema := struct {
		Backups []VirtualMachineBackup `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/backups", virtualMachineID), &schema)
	if err != nil {
		return []VirtualMachineBackup{}, err
	}
	return schema.Backups, nil
}

func (s *virtualMachineService) RestoreBackup(virtualMachineID string, backupTimestamp int) (Task, error) {
	schema := struct {
		Time int `json:"time"`
	}{
		Time: backupTimestamp,
	}
	data, err := json.Marshal(&schema)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "restore", data)
}

func (s *virtualMachineService) RestoreBackupToVApp(virtualMachineID, vappID string, backupTimestamp int) (Task, error) {
	schema := struct {
		Time   int    `json:"time"`
		VAppID string `json:"vapp_uuid"`
	}{
		Time:   backupTimestamp,
		VAppID: vappID,
	}
	data, err := json.Marshal(&schema)
	if err != nil {
		return Task{}, err
	}
	return s.postAction(virtualMachineID, "restore-into-vapp", data)
}

func (s *virtualMachineService) HasSnapshot(virtualMachineID string) (bool, error) {
	schema := struct {
		HasSnapshot bool `json:"has_snapshot"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/has-snapshot", virtualMachineID), &schema)
	if err != nil {
		return false, err
	}
	return schema.HasSnapshot, nil
}

func (s *virtualMachineService) GetSnapshot(virtualMachineID string) (Snapshot, error) {
	snapshot := Snapshot{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/snapshot", virtualMachineID), &snapshot)
	if err != nil {
		return Snapshot{}, err
	}
	return snapshot, nil
}

func (s *virtualMachineService) CreateSnapshot(virtualMachineID string) (Task, error) {
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
	return s.postAction(virtualMachineID, "create-snapshot", data)
}

func (s *virtualMachineService) RestoreSnapshot(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "restore-snapshot", []byte{})
}

func (s *virtualMachineService) RemoveSnapshot(virtualMachineID string) (Task, error) {
	return s.postAction(virtualMachineID, "remove-snapshot", []byte{})
}

func (s *virtualMachineService) GetNetworks(virtualMachineID string) ([]VAppNetwork, error) {
	schema := struct {
		Networks []VAppNetwork `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/networks", virtualMachineID), &schema)
	if err != nil {
		return []VAppNetwork{}, err
	}
	return schema.Networks, nil
}

func (s *virtualMachineService) GetCurrentBill(virtualMachineID string) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/billing", virtualMachineID), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *virtualMachineService) GetBill(virtualMachineID string, month, year int) (Billing, error) {
	billing := Billing{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/billing?month=%d&year=%d", virtualMachineID, month, year), &billing)
	if err != nil {
		return Billing{}, err
	}
	return billing, nil
}

func (s *virtualMachineService) GetMetadata(virtualMachineID string) ([]Metadata, error) {
	schema := struct {
		Metadata []Metadata `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/metadata", virtualMachineID), &schema)
	if err != nil {
		return []Metadata{}, err
	}
	return schema.Metadata, nil
}

func (s *virtualMachineService) UpdateMetadata(virtualMachineID string, metadata []Metadata) (Task, error) {
	payload, err := json.Marshal(&metadata)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/vms/%s/metadata", virtualMachineID), payload)
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

func (s *virtualMachineService) DeleteMetadata(virtualMachineID string, metadataKey string) (Task, error) {
	resp, err := s.client.Delete(fmt.Sprintf("/v1/vms/%s/metadata/%s", virtualMachineID, metadataKey))
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

func (s *virtualMachineService) GetPerformanceCounters(virtualMachineID string) ([]PerformanceCounter, error) {
	schema := struct {
		Counters []PerformanceCounter `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/performance-counters", virtualMachineID), &schema)
	if err != nil {
		return []PerformanceCounter{}, err
	}
	return schema.Counters, nil
}

func (s *virtualMachineService) GetPerformance(virtualMachineID string, counter PerformanceCounter, start, end time.Time) (Performance, error) {
	startNano := getUnixMilliseconds(start)
	endNano := getUnixMilliseconds(end)
	performance := Performance{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/performance/%s::%s::%s?start=%d&end=%d", virtualMachineID, counter.Group, counter.Name, counter.Type, startNano, endNano), &performance)
	if err != nil {
		return Performance{}, err
	}
	return performance, nil
}

type ConsoleSession struct {
	VMX    string `json:"vmx"`
	Ticket string `json:"ticket"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

func (s *virtualMachineService) GetConsoleSession(virtualMachineID string) (ConsoleSession, error) {
	session := ConsoleSession{}
	err := s.client.getObject(fmt.Sprintf("/v1/vms/%s/mks-screen-ticket", virtualMachineID), &session)
	if err != nil {
		return session, err
	}
	return session, err
}

func (s *virtualMachineService) GetScreenThumbnail(virtualMachineID string) ([]byte, error) {
	resp, err := s.client.Get(fmt.Sprintf("/v1/vms/%s/screen", virtualMachineID))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp)
}
