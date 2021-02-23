package iland

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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

type O365User struct {
	Name             string `json:"name"`
	DisplayName      string `json:"display_name"`
	OrganizationName string `json:"organization_name"`
	OrganizationUUID string `json:"organization_uuid"`
	Type             string `json:"type"`
	NativeID         string `json:"native_id"`
	BackedUp         bool   `json:"is_backed_up"`
	DeletedFromOrg   bool   `json:"is_deleted_from_org"`
}

func (s *o365Service) GetUsers(id string) ([]O365User, error) {
	schema := struct {
		Data     []O365User `json:"data"`
		Page     int        `json:"page"`
		PageSize int        `json:"page_size"`
	}{}
	limit := 100
	page := 0
	users := []O365User{}
	for {
		err := s.client.getObject(fmt.Sprintf("/v1/o365-organizations/%s/users?pageSize=%d&page=%d", id, limit, page), &schema)
		if err != nil {
			return []O365User{}, err
		}
		users = append(users, schema.Data...)
		if len(schema.Data) < limit {
			break
		}
		page++
	}
	return users, nil
}

func (s *o365Service) GetUserReport(id string) ([]byte, error) {
	resp, err := s.client.request(fmt.Sprintf("/v1/o365-organizations/%s/users-export", id), http.MethodPost, "application/vnd.ilandcloud.api.v1.0+octet-stream", []byte{})
	if err != nil {
		return []byte{}, err
	}
	defer resp.Close()
	return ioutil.ReadAll(resp)
}
