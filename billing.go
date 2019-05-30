package iland

type Billing struct {
	EntityID          string              `json:"entity_uuid"`
	EntityName        string              `json:"entity_name"`
	EntityType        string              `json:"entity_type"`
	CurrencyCode      string              `json:"currency_code"`
	TotalCost         float64             `json:"total_cost"`
	TotalCostEstimate float64             `json:"total_cost_estimage"`
	Year              int                 `json:"year"`
	Month             int                 `json:"month"`
	TestDrive         bool                `json:"test_drive"`
	LineItems         []BillingLineItem   `json:"line_items"`
	CPU               BillingResource     `json:"cpu"`
	Memory            BillingResource     `json:"memory"`
	Bandwidth         BillingResource     `json:"bandwidth"`
	Disk              BillingDiskResource `json:"disk"`
}

type BillingLineItem struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	ProductID string  `json:"product_id"`
}

type BillingResource struct {
	Total    BillingResourceUsage `json:"total"`
	Reserved BillingResourceUsage `json:"reserved"`
	Burst    BillingResourceUsage `json:"burst"`
}

type BillingResourceUsage struct {
	Cost  float64 `json:"cost"`
	Usage float64 `json:"usage"`
}

type BillingDiskResource struct {
	Total   BillingResource `json:"total"`
	HDD     BillingResource `json:"hdd"`
	SSD     BillingResource `json:"ssd"`
	Archive BillingResource `json:"archive"`
}
