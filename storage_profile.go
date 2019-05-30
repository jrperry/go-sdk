package iland

type StorageProfile struct {
	ID      string `json:"uuid"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Default bool   `json:"default_profile"`
	LimitMB int    `json:"limit"`
	UsedMB  int    `json:"storage_used_in_mb"`
	VdcID   string `json:"vdc_uuid"`
}
