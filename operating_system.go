package iland

type OperatingSystem struct {
	ID                        string   `json:"id"`
	Name                      string   `json:"internal_name"`
	Family                    string   `json:"family"`
	Description               string   `json:"name"`
	DefaultDiskAdapterType    string   `json:"default_disk_adapter_type"`
	MinimumDiskSizeGigabytes  int      `json:"minimum_disk_size_gigabytes"`
	MinimumMemoryMebibytes    int      `json:"minimum_memory_mebibytes"`
	X64                       bool     `json:"x64"`
	MaximumCPUCountField      int      `json:"maximum_cpu_count_field"`
	MinimumHardwareVersion    int      `json:"minimum_hardware_version"`
	PersonalizationEnabled    bool     `json:"personalization_enabled"`
	PersonalizationAuto       bool     `json:"personalization_auto"`
	SysPrepPackagingSupported bool     `json:"sys_prep_packaging_supported"`
	SupportsMemoryHotAdd      bool     `json:"supports_memory_hot_add"`
	SupportedForCreate        bool     `json:"supported_for_create"`
	SupportedVNICTypes        []string `json:"supported_vnic_types"`
}
