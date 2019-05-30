package iland

import (
	"io"
	"time"
)

type ConsoleService interface {
	Get(endpoint string) (io.ReadCloser, error)
	Post(endpoint string, body []byte) (io.ReadCloser, error)
	Put(endpoint string, body []byte) (io.ReadCloser, error)
	Delete(endpoint string) (io.ReadCloser, error)

	GetOperatingSystems() ([]OperatingSystem, error)
	GetLocations() []Location
	GetCompanies() ([]Company, error)
	GetOrgs() ([]Org, error)
	StreamEvents(companyID string) (chan Event, error)

	Location() LocationService
	Company() CompanyService
	User() UserService
	Org() OrgService
	Catalog() CatalogService
	VAppTemplate() VAppTemplateService
	Vdc() VdcService
	Edge() EdgeService
	OrgVdcNetwork() OrgVdcNetworkService
	VApp() VAppService
	VAppNetwork() VAppNetworkService
	VirtualMachine() VirtualMachineService
	VCCBackupTenant() VCCBackupTenantService
	Vpg() VpgService
	Task() TaskService
}

type LocationService interface {
	GetPublicCatalogs(locationID string) ([]Catalog, error)
	GetPublicVAppTemplates(locationID string) ([]VAppTemplate, error)
	GetPublicMedia(locationID string) ([]Media, error)
}

type TaskService interface {
	Get(taskID string) (Task, error)
	Track(taskID string) (Task, error)
}

type VCCBackupTenantService interface {
	Get(vccBackupTenantID string) (VCCBackupTenant, error)
}

type CompanyService interface {
	Get(companyID string) (Company, error)
	GetUsers(companyID string) ([]User, error)
	CreateUser(companyID string, params CreateUserParams) (User, error)
	GetRoles(companyID string) ([]Role, error)
	GetRole(companyID, roleID string) (Role, error)
	GetOrgs(companyID string) ([]Org, error)
	GetLocationOrgs(companyID, locationID string) ([]Org, error)
	GetVCCBackupTenants(companyID string) ([]VCCBackupTenant, error)
}

type UserService interface {
	Get(username string) (User, error)
	Update(username string, params UpdateUserParams) (User, error)
	GetCompanies(username string) ([]Company, error)
	GetOrgs(username string) ([]Org, error)
	AssignRole(username, companyID, roleID string) error
	GetRole(username, companyID string) (Role, error)
	DeleteRole(username, companyID string) error
}

type OrgService interface {
	Get(orgID string) (Org, error)
	GetVdcs(orgID string) ([]Vdc, error)
	GetEdges(orgID string) ([]Edge, error)
	GetCatalogs(orgID string) ([]Catalog, error)
	GetVAppTemplates(orgID string) ([]VAppTemplate, error)
	GetMedia(orgID string) ([]Media, error)
	GetNetworks(orgID string) ([]OrgVdcNetwork, error)
	GetVApps(orgID string) ([]VApp, error)
	GetVirtualMachines(orgID string) ([]VirtualMachine, error)
	GetVpgs(orgID string) ([]Vpg, error)
	GetPublicIPs(orgID string) ([]string, error)
	GetPublicIPAssignments(orgID string) ([]PublicIPAssignment, error)
	GetCurrentBill(orgID string) (Billing, error)
	GetBill(orgID string, month, year int) (Billing, error)
}

type CatalogService interface {
	Get(catalogID string) (Catalog, error)
	Update(catalogID string, params UpdateCatalogParams) (Task, error)
	GetVAppTemplates(catalogID string) ([]VAppTemplate, error)
	GetMedia(catalogID string) ([]Media, error)
	CreateVAppTemplate(catalogID string, params CreateVAppTemplateParams) (Task, error)
	SyncSubscription(catalogID string) (Task, error)
}

type VAppTemplateService interface {
	Get(vappTemplateID string) (VAppTemplate, error)
	Update(vappTemplateID string, params UpdateVAppTemplateParams) (Task, error)
	Delete(vappTemplateID string) (Task, error)
	GetVirtualMachines(vappTemplateID string) ([]VirtualMachineTemplate, error)
	GetConfig(vappTemplateID string) (VAppTemplateConfig, error)
	SyncSubscription(vappTemplateID string) (Task, error)
}

type VdcService interface {
	Get(vdcID string) (Vdc, error)
	GetStorageProfiles(vdcID string) ([]StorageProfile, error)
	GetSummary(vdcID string) (VdcSummary, error)
	GetVApps(vdcID string) ([]VApp, error)
	GetVirtualMachines(vdcID string) ([]VirtualMachine, error)
	GetEdges(vdcID string) ([]Edge, error)
	GetNetworks(vdcID string) ([]OrgVdcNetwork, error)
	GetCurrentBill(vdcID string) (Billing, error)
	GetBill(vdcID string, month, year int) (Billing, error)
	GetCurrentVAppBill(vdcID string) ([]Billing, error)
	GetVAppBill(vdcID string, month, year int) ([]Billing, error)
	GetPerformanceCounters(vdcID string) ([]PerformanceCounter, error)
	GetPerformance(vdcID string, counter PerformanceCounter, start, end time.Time) (Performance, error)
	BuildVApp(vdcID string, params BuildVAppParams) (Task, error)
	DeployVAppTemplate(vdcID string, params DeployVAppTemplateParams) (Task, error)
}

type EdgeService interface {
	Get(edgeID string) (Edge, error)
	GetFirewall(edgeID string) (EdgeFirewall, error)
	UpdateFirewallRules(edgeID string, rules []EdgeFirewallRule) (Task, error)
	GetNAT(edgeID string) (EdgeNAT, error)
	UpdateNATRules(edgeID string, rules []EdgeNATRule) (Task, error)
	EnableNAT(edgeID string) (Task, error)
	DisableNAT(edgeID string) (Task, error)
}

type OrgVdcNetworkService interface {
	Get(networkID string) (OrgVdcNetwork, error)
	Update(networkID string, params UpdateOrgVdcNetworkParams) (Task, error)
}

type VAppService interface {
	Get(vappID string) (VApp, error)
	Delete(vappID string) (Task, error)
	GetVirtualMachines(vappID string) ([]VirtualMachine, error)
	GetNetworks(vappID string) ([]VAppNetwork, error)
	AddOrgNetwork(vappID, orgVdcNetworkID string) (Task, error)
	UpdateName(vappID, name string) (Task, error)
	UpdateDescription(vappID, description string) (Task, error)
	Copy(vappID string, params CopyVAppParams) (Task, error)
	Move(vappID string, params MoveVAppParams) (Task, error)
	BuildVirtualMachines(vappID string, params []BuildVirtualMachineParams) (Task, error)
	AddTemplateVirtualMachines(vappID string, params []AddTemplateVirtualMachineParams) (Task, error)
	CreateNetwork(vappID string, params CreateVAppNetworkParams) (Task, error)
	PowerOn(vappID string) (Task, error)
	PowerOff(vappID string) (Task, error)
	Shutdown(vappID string) (Task, error)
	Reboot(vappID string) (Task, error)
	Reset(vappID string) (Task, error)
	Suspend(vappID string) (Task, error)
	GetCurrentBill(vappID string) (Billing, error)
	GetBill(vappID string, month, year int) (Billing, error)
	GetAvailableStorageProfiles(vappID string) ([]StorageProfile, error)
	GetMetadata(vappID string) ([]Metadata, error)
	UpdateMetadata(vappID string, metadata []Metadata) (Task, error)
	DeleteMetadata(vappID, metadataKey string) (Task, error)
	HasSnapshot(vappID string) (bool, error)
	GetSnapshot(vappID string) (Snapshot, error)
	CreateSnapshot(vappID string) (Task, error)
	RestoreSnapshot(vappID string) (Task, error)
	RemoveSnapshot(vappID string) (Task, error)
	GetStartupSettings(vappID string) ([]VAppStartupSetting, error)
	UpdateStartupSettings(vappID string, params []VAppStartupSetting) (Task, error)
	GetPerformanceCounters(vappID string) ([]PerformanceCounter, error)
	GetPerformance(vappID string, counter PerformanceCounter, start, end time.Time) (Performance, error)
	GetSummary(vappID string) (VAppSummary, error)
}

type VAppNetworkService interface {
	Get(vappNetworkID string) (VAppNetwork, error)
	Update(vappNetworkID string, params UpdateVAppNetworkParams) (Task, error)
	Delete(vappNetworkID string) (Task, error)
	UpdateDHCP(vappNetworkID string, params DHCP) (Task, error)
	GetFirewall(vappNetworkID string) (VAppNetworkFirewall, error)
	UpdateFirewallRules(vappNetworkID string, rules []VAppNetworkFirewallRule) (Task, error)
	EnableFirewall(vappNetworkID string) (Task, error)
	DisableFirewall(vappNetworkID string) (Task, error)
	GetNAT(vappNetworkID string) (VAppNetworkNAT, error)
	UpdateNATIPTranslationRules(vappNetworkID string, rules []IPTranslationRule) (Task, error)
	UpdateNATPortForwardingRules(vappNetworkID string, rules []PortForwardingRule) (Task, error)
	EnableNAT(vappNetworkID string) (Task, error)
	DisableNAT(vappNetworkID string) (Task, error)
	GetInterfaces(vappNetwork string) ([]VirtualMachineInterface, error)
}

type VirtualMachineService interface {
	Get(virtualMachineID string) (VirtualMachine, error)
	Delete(virtualMachineID string) (Task, error)
	UpdateName(virtualMachineID, name string) (Task, error)
	UpdateDescription(virtualMachineID, description string) (Task, error)
	PowerOn(virtualMachineID string) (Task, error)
	PowerOnForceCustomization(virtualMachineID string) (Task, error)
	PowerOff(virtualMachineID string) (Task, error)
	Reboot(virtualMachineID string) (Task, error)
	Reset(virtualMachineID string) (Task, error)
	Shutdown(virtualMachineID string) (Task, error)
	Suspend(virtualMachineID string) (Task, error)
	Copy(virtualMachineID string, params CopyVirtualMachineParams) (Task, error)
	Move(virtualMachineID string, params MoveVirtualMachineParams) (Task, error)

	GetSummary(virtualMachineID string) (Summary, error)
	GetAvailableStorageProfiles(virtualMachineID string) ([]StorageProfile, error)
	ChangeStorageProfile(virtualMachineID, storageProfileID string) (Task, error)
	EnableNestedHypervisor(virtualMachineID string) (Task, error)
	DisableNestedHypervisor(virtualMachineID string) (Task, error)
	InsertMedia(virtualMachineID, mediaID string) (Task, error)
	EjectMedia(virtualMachineID string) (Task, error)
	GetGuestCustomization(virtualMachineID string) (GuestCustomization, error)
	UpdateGuestCustomization(virtualMachineID string, params GuestCustomization) (Task, error)
	GetHotAdd(virtualMachineID string) (HotAdd, error)
	UpdateHotAdd(virtualMachineID string, params HotAdd) (Task, error)
	GetBootOptions(virtualMachineID string) (BootOptions, error)
	UpdateBootOptions(virtualMachineID string, params BootOptions) (Task, error)
	UpdateHardwareVersion(virtualMachineID string) (Task, error)
	GetVMwareTools(virtualMachineID string) (VMwareTools, error)
	UpgradeVMwareTools(virtualMachineID string) (Task, error)
	InstallVMwareTools(virtualMachineID string) (Task, error)
	Reconfigure(virtualMachineID string, params ReconfigureParams) (Task, error)
	GetDisks(virtualMachineID string) ([]Disk, error)
	AddDisk(virtualMachineID string, params DiskParams) (Task, error)
	UpdateDisk(virtualMachineID string, params DiskParams) (Task, error)
	UpdateDisks(virtualMachineID string, params []DiskParams) (Task, error)
	DeleteDisk(virtualMachineID string, diskName string) (Task, error)
	GetRecommendedBusType(virtualMachineID string) (string, error)
	GetNics(virtualMachineID string) ([]Nic, error)
	DeleteNic(virtualMachineID string, nicID int) (Task, error)
	UpdateNics(virtualMachineID string, params []Nic) (Task, error)
	UpdateCPU(virtualMachineID string, params UpdateCPUParams) (Task, error)
	UpdateCPUCount(virtualMachineID string, cpuCount int) (Task, error)
	UpdateMemory(virtualMachineID string, memorySize int) (Task, error)
	GetBackups(virtualMachineID string) ([]VirtualMachineBackup, error)
	RestoreBackup(virtualMachineID string, backupTimestamp int) (Task, error)
	RestoreBackupToVApp(virtualMachineID, vappID string, backupTimestamp int) (Task, error)
	HasSnapshot(virtualMachineID string) (bool, error)
	GetSnapshot(virtualMachineID string) (Snapshot, error)
	CreateSnapshot(virtualMachineID string) (Task, error)
	RestoreSnapshot(virtualMachineID string) (Task, error)
	RemoveSnapshot(virtualMachineID string) (Task, error)

	GetNetworks(virtualMachineID string) ([]VAppNetwork, error)
	GetCurrentBill(virtualMachineID string) (Billing, error)
	GetBill(virtualMachineID string, month, year int) (Billing, error)
	GetMetadata(virtualMachineID string) ([]Metadata, error)
	UpdateMetadata(virtualMachineID string, metadata []Metadata) (Task, error)
	DeleteMetadata(virtualMachineID string, metadataKey string) (Task, error)
	GetPerformanceCounters(virtualMachineID string) ([]PerformanceCounter, error)
	GetPerformance(virtualMachineID string, counter PerformanceCounter, start, end time.Time) (Performance, error)
	GetConsoleSession(virtualMachineID string) (ConsoleSession, error)
	GetScreenThumbnail(virtualMachineID string) ([]byte, error)
}

type VpgService interface {
	Get(vpgID string) (Vpg, error)
	GetCheckpoints(vpgID string) ([]VpgCheckpoint, error)
}
