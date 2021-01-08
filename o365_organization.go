package iland

import "fmt"

type o365Service struct {
	client *client
}

type O365Organization struct {
	ID                    string `json:"uuid"`
	Name                  string `json:"name"`
	CompanyID             string `json:"crm"`
	LocationID            string `json:"location_id"`
	ContractUUID          string `json:"contract_uuid"`
	Type                  string `json:"type"`
	Region                string `json:"region"`
	FirstBackupTime       int    `json:"first_backup_time"`
	LastBackupTime        int    `json:"last_backup_time"`
	TotalUsers            int    `json:"total_users"`
	TotalBackedUpUsers    int    `json:"total_backedup_users"`
	TotalLicensedUsers    int    `json:"total_licensed_users"`
	TotalLicensesConsumed int    `json:"total_licenses_consumed"`
	UnprotectedUsers      int    `json:"unprotected_users"`
	BackedUp              bool   `json:"backed_up"`
	ExchangeOnline        bool   `json:"exchange_online"`
	Trial                 bool   `json:"trial"`
	SharepointOnline      bool   `json:"share_point_online"`
}

func (s *o365Service) GetOrganization(id string) (O365Organization, error) {
	org := O365Organization{}
	err := s.client.getObject(fmt.Sprintf("/v1/o365-organizations/%s", id), &org)
	if err != nil {
		return O365Organization{}, err
	}
	return org, nil
}
