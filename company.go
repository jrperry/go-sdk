package iland

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	ID      string `json:"company_id"`
	Name    string `json:"name"`
	HasIAAS bool   `json:"has_iaas"`
	HasVCC  bool   `json:"has_vcc"`
	Domain  Domain `json:"domain"`
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
	FullName string `json:"full_name"`
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
