package iland

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"name"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Type        string `json:"user_type"`
	Locked      bool   `json:"locked"`
	Country     string `json:"country"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Address     string `json:"address"`
	Domain      Domain `json:"domain"`
	CreatedDate int    `json:"created_date"`
}

type userService struct {
	client *client
}

func (s *userService) Get(username string) (User, error) {
	user := User{}
	err := s.client.getObject(fmt.Sprintf("/v1/users/%s", username), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

type UpdateUserParams struct {
	FullName string `json:"fullname"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Address  string `json:"address"`
}

func (s *userService) Update(username string, params UpdateUserParams) (User, error) {
	data, err := json.Marshal(&params)
	if err != nil {
		return User{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/users/%s", username), data)
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

func (s *userService) GetCompanies(username string) ([]Company, error) {
	schema := struct {
		Companies []Company `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/users/%s/companies", username), &schema)
	if err != nil {
		return []Company{}, err
	}
	return schema.Companies, nil
}

func (s *userService) GetOrgs(username string) ([]Org, error) {
	schema := struct {
		Orgs []Org `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/users/%s/orgs", username), &schema)
	if err != nil {
		return []Org{}, err
	}
	return schema.Orgs, nil
}

func (s *userService) AssignRole(username, companyID, roleID string) error {
	params := struct {
		RoleUUID string `json:"role_uuid"`
	}{
		RoleUUID: roleID,
	}
	data, err := json.Marshal(&params)
	if err != nil {
		return err
	}
	_, err = s.client.Put(fmt.Sprintf("/v1/users/%s/roles/%s", username, companyID), data)
	return err
}

func (s *userService) GetRole(username, companyID string) (Role, error) {
	role := Role{}
	err := s.client.getObject(fmt.Sprintf("/v1/users/%s/roles/%s", username, companyID), &role)
	if err != nil {
		return Role{}, err
	}
	return role, err
}

func (s *userService) DeleteRole(username, companyID string) error {
	_, err := s.client.Delete(fmt.Sprintf("/v1/users/%s/roles/%s", username, companyID))
	return err
}
