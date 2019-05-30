package iland

import (
	"encoding/json"
	"fmt"
)

type Catalog struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     int    `json:"version"`
	IsPublic    bool   `json:"catalog_public"`
	Shared      bool   `json:"shared"`
	Subscribed  bool   `json:"subscribed"`
	OrgID       string `json:"org_uuid"`
	CompanyID   string `json:"company_id"`
	LocationID  string `json:"location_id"`
	CreatedDate int    `json:"created_date"`
	UpdatedDate int    `json:"updated_date"`
}

type catalogService struct {
	client *client
}

func (s *catalogService) Get(catalogID string) (Catalog, error) {
	catalog := Catalog{}
	err := s.client.getObject(fmt.Sprintf("/v1/catalogs/%s", catalogID), &catalog)
	if err != nil {
		return Catalog{}, err
	}
	return catalog, nil
}

type UpdateCatalogParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *catalogService) Update(catalogID string, params UpdateCatalogParams) (Task, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Put(fmt.Sprintf("/v1/catalogs/%s", catalogID), data)
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

func (s *catalogService) GetVAppTemplates(catalogID string) ([]VAppTemplate, error) {
	schema := struct {
		VAppTemplates []VAppTemplate `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/catalogs/%s/vapp-templates", catalogID), &schema)
	if err != nil {
		return []VAppTemplate{}, err
	}
	return schema.VAppTemplates, nil
}

func (s *catalogService) GetMedia(catalogID string) ([]Media, error) {
	schema := struct {
		Media []Media `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/catalogs/%s/media", catalogID), &schema)
	if err != nil {
		return []Media{}, err
	}
	return schema.Media, nil
}

type CreateVAppTemplateParams struct {
	VAppID      string `json:"vapp_uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *catalogService) CreateVAppTemplate(catalogID string, params CreateVAppTemplateParams) (Task, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return Task{}, err
	}
	resp, err := s.client.Post(fmt.Sprintf("/v1/catalogs/%s/actions/add-vapp-template-from-vapp", catalogID), data)
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

func (s *catalogService) SyncSubscription(catalogID string) (Task, error) {
	resp, err := s.client.Post(fmt.Sprintf("/v1/catalogs/%s/actions/sync", catalogID), []byte{})
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
