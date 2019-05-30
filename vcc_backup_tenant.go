package iland

import "fmt"

type VCCBackupTenant struct {
	ID                   string             `json:"uuid"`
	UID                  string             `json:"uid"`
	Name                 string             `json:"name"`
	Enabled              bool               `json:"enabled"`
	LastResult           string             `json:"last_result"`
	LastActive           int                `json:"last_active"`
	BackupCount          int                `json:"backup_count"`
	ThrottlingEnabled    bool               `json:"throttling_enabled"`
	ThrottlingSpeedLimit int                `json:"throttling_speed_limit"`
	ThrottlingSpeedUnit  string             `json:"throttling_speed_unit"`
	Resources            VCCBackupResources `json:"resources"`
	PublicIPCount        int                `json:"public_ip_count"`
	ContractID           string             `json:"contract_uuid"`
	CompanyName          string             `json:"owner_name"`
	CompanyID            string             `json:"company_id"`
	LocationID           string             `json:"location_id"`
	UpdatedDate          int                `json:"updated_date"`
}

type VCCBackupResources struct {
	Resources []VCCBackupResource `json:"resources"`
}

type VCCBackupResource struct {
	Repository VCCRepository `json:"repository"`
}

type VCCRepository struct {
	Name           string `json:"display_name"`
	StorageQuotaMB int    `json:"quota"`
	StorageUsedMB  int    `json:"used_quota"`
}

type vccBackupTenantService struct {
	client *client
}

func (s *vccBackupTenantService) Get(vccBackupTenantID string) (VCCBackupTenant, error) {
	vccBackupTenant := VCCBackupTenant{}
	err := s.client.getObject(fmt.Sprintf("/v1/vcc-backup-tenants/%s", vccBackupTenantID), &vccBackupTenant)
	if err != nil {
		return VCCBackupTenant{}, err
	}
	return vccBackupTenant, nil
}
