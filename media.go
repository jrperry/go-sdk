package iland

type Media struct {
	ID               string  `json:"uuid"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Status           int     `json:"status"`
	SizeGB           float64 `json:"size"`
	IsPublic         bool    `json:"is_public"`
	CatalogID        string  `json:"catalog_uuid"`
	StorageProfileID string  `json:"storage_profile_uuid"`
	VdcID            string  `json:"vdc_uuid"`
	OrgID            string  `json:"org_uuid"`
	CompanyID        string  `json:"company_id"`
	LocationID       string  `json:"location_id"`
	UpdatedDate      int     `json:"updated_date"`
}
