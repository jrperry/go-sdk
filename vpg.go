package iland

import (
	"fmt"
)

type Vpg struct {
	ID                   string              `json:"uuid"`
	ZertoID              string              `json:"vpg_identifier"`
	Name                 string              `json:"name"`
	Status               string              `json:"status"`
	SubStatus            string              `json:"sub_status"`
	Priority             string              `json:"priority"`
	VirtualMachineCount  int                 `json:"vms_count"`
	SourceSite           string              `json:"source_site"`
	TargetSite           string              `json:"target_site"`
	ActualRPO            int                 `json:"actual_rpo"`
	IOPS                 int                 `json:"iops"`
	ProvisionedStorageMB int                 `json:"provisioned_storage_in_mb"`
	UsedStorageMB        int                 `json:"used_storage_in_mb"`
	ThroughputMB         float64             `json:"throughput_in_mb"`
	JournalStorageUsedMB int                 `json:"recovery_journal_used_storage_in_mb"`
	BackupEnabled        bool                `json:"backup_enabled"`
	VirtualMachines      []VpgVirtualMachine `json:"vms"`
	Entities             VpgEntities         `json:"entities"`
	ServiceProfileID     string              `json:"service_profile_uuid"`
	OrgID                string              `json:"org_uuid"`
	CompanyID            string              `json:"company_id"`
	LocationID           string              `json:"location_id"`
	UpdatedDate          int                 `json:"updated_date"`
}

type VpgEntities struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type vpgService struct {
	client *client
}

func (s *vpgService) Get(vpgID string) (Vpg, error) {
	vpg := Vpg{}
	err := s.client.getObject(fmt.Sprintf("/v1/vpgs/%s?expand=VPG_VM", vpgID), &vpg)
	if err != nil {
		return Vpg{}, err
	}
	return vpg, nil
}

type VpgVirtualMachine struct {
	ID                   string      `json:"uuid"`
	ZertoID              string      `json:"vm_identifier"`
	Name                 string      `json:"vm_name"`
	Status               string      `json:"status"`
	SubStatus            string      `json:"sub_status"`
	Priority             string      `json:"priority"`
	ProvisionedStorageMB int         `json:"provisioned_storage_in_mb"`
	UsedStorageMB        int         `json:"used_storage_in_mb"`
	IOPS                 int         `json:"iops"`
	ThroughputMB         float64     `json:"throughput_in_mb"`
	TargetSite           string      `json:"target_site"`
	SourceSite           string      `json:"source_site"`
	ActualRpo            int         `json:"actual_rpo"`
	LastTest             int         `json:"last_test"`
	Entities             VpgEntities `json:"entities"`
	VpgID                string      `json:"vpg_uuid"`
	VpgName              string      `json:"vpg_name"`
	OrgID                string      `json:"org_uuid"`
	LocationID           string      `json:"location"`
}

type VpgServiceProfile struct {
	ID                    string `json:"uuid"`
	ZertoID               string `json:"service_profile_identifier"`
	Name                  string `json:"service_profile_name"`
	Description           string `json:"description"`
	History               int    `json:"history"`
	MaxJournalSizePercent int    `json:"max_journal_size_in_percent"`
	Rpo                   int    `json:"rpo"`
	TestInterval          int    `json:"test_interval"`
	LocationID            string `json:"location"`
}

type VpgCheckpoint struct {
	ID        string `json:"checkpoint_identifier"`
	Tag       string `json:"tag"`
	Timestamp int    `json:"time_stamp"`
}

func (s *vpgService) GetCheckpoints(vpgID string) ([]VpgCheckpoint, error) {
	schema := struct {
		Checkpoints []VpgCheckpoint `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/vpgs/%s/checkpoints", vpgID), &schema)
	if err != nil {
		return []VpgCheckpoint{}, err
	}
	return schema.Checkpoints, nil
}
