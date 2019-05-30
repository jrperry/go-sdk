package iland

type Event struct {
	ID              string `json:"uuid"`
	Details         string `json:"details"`
	Type            string `json:"type"`
	EntityID        string `json:"entity_uuid"`
	EntityName      string `json:"entity_name"`
	EntityType      string `json:"entity_type"`
	OwnerType       string `json:"owner_type"`
	OwnerID         string `json:"owner_id"`
	TaskID          string `json:"task_uuid"`
	InitiatedByUser string `json:"initiated_by_username"`
	InitiatedByName string `json:"initiated_by_full_name"`
	Timestamp       int    `json:"timestamp"`
}
