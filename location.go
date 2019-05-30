package iland

import "fmt"

type Location struct {
	ID string `json:"location_id"`
}

type locationService struct {
	client *client
}

func (s *locationService) GetPublicCatalogs(locationID string) ([]Catalog, error) {
	schema := struct {
		Catalogs []Catalog `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/locations/%s/public-catalogs", locationID), &schema)
	if err != nil {
		return []Catalog{}, err
	}
	return schema.Catalogs, nil
}

func (s *locationService) GetPublicVAppTemplates(locationID string) ([]VAppTemplate, error) {
	schema := struct {
		VAppTemplates []VAppTemplate `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/locations/%s/public-vapp-templates", locationID), &schema)
	if err != nil {
		return []VAppTemplate{}, err
	}
	return schema.VAppTemplates, nil
}

func (s *locationService) GetPublicMedia(locationID string) ([]Media, error) {
	schema := struct {
		Media []Media `json:"data"`
	}{}
	err := s.client.getObject(fmt.Sprintf("/v1/locations/%s/public-media", locationID), &schema)
	if err != nil {
		return []Media{}, err
	}
	return schema.Media, nil
}
