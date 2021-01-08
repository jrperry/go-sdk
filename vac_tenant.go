package iland

import "fmt"

type VacTenant struct {
	ID                              string              `json:"uuid"`
	Name                            string              `json:"name"`
	CRM                             string              `json:"crm"`
	LocationID                      string              `json:"location_id"`
	ContractUUID                    string              `json:"contract_uuid"`
	OwnerName                       string              `json:"owner_name"`
	LastActive                      int                 `json:"last_active"`
	LastResult                      string              `json:"last_result"`
	InstanceUID                     string              `json:"instance_uid"`
	CloudConnectAgentUID            string              `json:"cloud_connect_agent_uid"`
	SiteName                        string              `json:"site_name"`
	VcdOrganizationUID              string              `json:"vcd_organization_uid"`
	TenantType                      string              `json:"tenant_type"`
	Description                     string              `json:"description"`
	Title                           string              `json:"title"`
	FirstName                       string              `json:"first_name"`
	LastName                        string              `json:"last_name"`
	Username                        string              `json:"user_name"`
	EmailAddress                    string              `json:"email_address"`
	IsEnabled                       bool                `json:"is_enabled"`
	TaxID                           string              `json:"tax_id"`
	Telephone                       string              `json:"telephone"`
	Country                         string              `json:"country"`
	City                            string              `json:"city"`
	Street                          string              `json:"street"`
	USState                         string              `json:"us_state"`
	ZipCode                         string              `json:"zip_code"`
	Domain                          string              `json:"domain"`
	CompanyID                       string              `json:"company_id"`
	Notes                           string              `json:"notes"`
	BackupProtectionEnabled         bool                `json:"backup_protection_enabled"`
	BackupProtectionPeriod          int                 `json:"backup_protection_period"`
	NetworkFailoverResourcesEnabled bool                `json:"network_failover_resources_enabled"`
	NumberOfPublicIP                int                 `json:"number_of_public_ip"`
	PublicIPEnabled                 bool                `json:"public_ip_enabled"`
	MaxConcurrentTasks              int                 `json:"max_concurrent_tasks"`
	BandwidthThrottlingEnabled      bool                `json:"bandwidth_throttling_enabled"`
	AllowedBandwidth                int                 `json:"allowed_bandwidth"`
	AllowedBandwidthUnits           string              `json:"allowed_bandwidth_units"`
	GatewayFailoverEnabled          bool                `json:"gateway_failover_enabled"`
	VmsBackedUp                     int                 `json:"vms_backed_up"`
	VmsReplicated                   int                 `json:"vms_replicated"`
	VmsBackedUpToCloud              int                 `json:"vms_backed_up_to_cloud"`
	ManagedPhysicalWorkstations     int                 `json:"managed_physical_workstations"`
	ManagedCloudWorkstations        int                 `json:"managed_cloud_workstations"`
	ManagedPhysicalServers          int                 `json:"managed_physical_servers"`
	ManagedCloudServers             int                 `json:"managed_cloud_servers"`
	ExpirationEnabled               bool                `json:"expiration_enabled"`
	ExpirationDate                  int                 `json:"expiration_date"`
	TotalStorageQuota               int                 `json:"total_storage_quota"`
	UsedStorageQuota                int                 `json:"used_storage_quota"`
	Endpoint                        string              `json:"endpoint"`
	BackupResources                 []VacBackupResource `json:"backup_resources"`
}

type VacBackupResource struct {
	ID                     string  `json:"id"`
	Name                   string  `json:"cloud_repository_name"`
	StorageQuota           int     `json:"storage_quota"`
	UsedStorageQuota       int     `json:"used_storage_quota"`
	VmsQuota               int     `json:"vms_quota"`
	WorkstationsQuota      int     `json:"workstations_quota"`
	ServersQuota           int     `json:"servers_quota"`
	TrafficQuota           int     `json:"traffic_quota"`
	UsedTrafficQuota       float64 `json:"used_traffic_quota"`
	QuotasAreUnlimited     bool    `json:"quotas_are_unlimited"`
	WanAccelerationEnabled bool    `json:"wan_acceleration_enabled"`
	WanAcclerator          string  `json:"wan_accelerator"`
	IntervalStartTime      int     `json:"interval_start_time"`
	IntervalEndTime        int     `json:"interval_end_time"`
}

type vacTenantService struct {
	client *client
}

func (s *vacTenantService) Get(id string) (VacTenant, error) {
	vacTenant := VacTenant{}
	err := s.client.getObject(fmt.Sprintf("/v1/vac-companies/%s", id), &vacTenant)
	if err != nil {
		return VacTenant{}, err
	}
	return vacTenant, nil
}
