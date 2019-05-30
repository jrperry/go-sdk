package iland

type Role struct {
	ID          string   `json:"uuid"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	CompanyID   string   `json:"company_id"`
	Policies    []Policy `json:"policies"`
}

type Policy struct {
	EntityID    string   `json:"entity_uuid"`
	Type        string   `json:"type"`
	DomainType  string   `json:"domain_type"`
	Permissions []string `json:"permissions"`
}
