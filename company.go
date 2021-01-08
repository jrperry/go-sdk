package iland

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	ID               string `json:"company_id"`
	Name             string `json:"name"`
	HasIAAS          bool   `json:"has_iaas"`
	HasVCC           bool   `json:"has_vcc"`
	HasVCCR          bool   `json:"has_vccr"`
	HasObjectStorage bool   `json:"has_object_storage"`
	HasO365          bool   `json:"has_o365"`
	Domain           Domain `json:"domain"`
}

type companyService struct {
	client *client
}

func (s *companyService) Get(companyID string) (Company, error) {
	company := Company{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s", companyID), &company)
	if err != nil {
		return Company{}, err
	}
	return company, nil
}

func (s *companyService) GetUsers(companyID string) ([]User, error) {
	schema := struct {
		Users []User `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/users", companyID), &schema)
	if err != nil {
		return []User{}, err
	}
	return schema.Users, nil
}

type CreateUserParams struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	DomainID string `json:"domain"`
	Password string `json:"password"`
}

func (s *companyService) CreateUser(companyID string, params CreateUserParams) (User, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return User{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/companies/%s/users", companyID), data)
	if err != nil {
		return User{}, err
	}
	user := User{}
	err = unmarshalBody(resp, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *companyService) GetRoles(companyID string) ([]Role, error) {
	schema := struct {
		Roles []Role `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/roles", companyID), &schema)
	if err != nil {
		return []Role{}, err
	}
	return schema.Roles, nil
}

func (s *companyService) GetRole(companyID, roleID string) (Role, error) {
	role := Role{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/roles/%s", companyID, roleID), &role)
	if err != nil {
		return Role{}, err
	}
	return role, nil
}

func (s *companyService) GetOrgs(companyID string) ([]Org, error) {
	orgs := []Org{}
	for _, locationID := range LocationIDs {
		schema := struct {
			Orgs []Org `json:"data"`
		}{}
		err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/location/%s/orgs", companyID, locationID), &schema)
		if err != nil {
			return []Org{}, err
		}
		orgs = append(orgs, schema.Orgs...)
	}
	return orgs, nil
}

func (s *companyService) GetLocationOrgs(companyID, locationID string) ([]Org, error) {
	schema := struct {
		Orgs []Org `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/location/%s/orgs", companyID, locationID), &schema)
	if err != nil {
		return []Org{}, err
	}
	return schema.Orgs, nil
}

func (s *companyService) GetVCCBackupTenants(companyID string) ([]VCCBackupTenant, error) {
	schema := struct {
		VCCBackupTenants []VCCBackupTenant `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/vcc-backup-tenants", companyID), &schema)
	if err != nil {
		return []VCCBackupTenant{}, err
	}
	return schema.VCCBackupTenants, nil
}

func (s *companyService) GetLocationVacTenants(companyID, location string) ([]VacTenant, error) {
	schema := struct {
		Tenants []VacTenant `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/location/%s/vac-companies", companyID, location), &schema)
	if err != nil {
		return []VacTenant{}, err
	}
	return schema.Tenants, nil
}

func (s *companyService) GetVacTenants(companyID string) ([]VacTenant, error) {
	schema := struct {
		Tenants []VacTenant `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/companies/%s/vac-companies", companyID), &schema)
	if err != nil {
		return []VacTenant{}, err
	}
	return schema.Tenants, nil
}

type CompanyInventory struct {
	CompanyID   string    `json:"company_id"`
	CompanyName string    `json:"company_name"`
	Entities    Inventory `json:"entities"`
}

type Inventory struct {
	Company                []InventoryItem `json:"COMPANY"`
	VCCBackupLocations     []InventoryItem `json:"VCC_BACKUP_LOCATION"`
	VCCBackupProducts      []InventoryItem `json:"VCC_BACKUP_PRODUCT"`
	VCCBackupTenants       []InventoryItem `json:"VCC_BACKUP_TENANT"`
	VACBackupJob           []InventoryItem `json:"VAC_BACKUP_JOB"`
	VCCFailoverPlans       []InventoryItem `json:"IAAS_VCC_FAILOVER_PLAN"`
	IaasProducts           []InventoryItem `json:"IAAS_PRODUCT"`
	IaasLocations          []InventoryItem `json:"IAAS_LOCATION"`
	IaasOrganizations      []InventoryItem `json:"IAAS_ORGANIZATION"`
	IaasVdcs               []InventoryItem `json:"IAAS_VDC"`
	IaasCatalogs           []InventoryItem `json:"IAAS_CATALOG"`
	IaasMedia              []InventoryItem `json:"IAAS_MEDIA"`
	IaasNetworks           []InventoryItem `json:"IAAS_INTERNAL_NETWORK"`
	IaasEdges              []InventoryItem `json:"IAAS_EDGE"`
	IaasVApps              []InventoryItem `json:"IAAS_VAPP"`
	IaasVAppNetworks       []InventoryItem `json:"IAAS_VAPP_NETWORK"`
	IaasVms                []InventoryItem `json:"IAAS_VM"`
	IaasVpgs               []InventoryItem `json:"IAAS_VPG"`
	ObjectStorageProducts  []InventoryItem `json:"OBJECT_STORAGE_PRODUCT"`
	ObjectStorageLocations []InventoryItem `json:"OBJECT_STORAGE_LOCATION"`
	O365Product            []InventoryItem `json:"O365_PRODUCT"`
	O365Locations          []InventoryItem `json:"O365_LOCATION"`
	O365Organization       []InventoryItem `json:"O365_ORGANIZATION"`
	O365Jobs               []InventoryItem `json:"O365_JOB"`
}

type InventoryItem struct {
	UUID       string `json:"uuid"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	ParentUUID string `json:"parent_uuid"`
	ParentType string `json:"parent_type"`
}

func (s *companyService) GetInventory(companyID string) (CompanyInventory, error) {
	schema := struct {
		Username  string             `json:"username"`
		Inventory []CompanyInventory `json:"inventory"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/users/%s/inventory?company=%s", s.client.username, companyID), &schema)
	if err != nil {
		return CompanyInventory{}, err
	}
	if len(schema.Inventory) == 0 {
		return CompanyInventory{}, nil
	}
	return schema.Inventory[0], nil
}
