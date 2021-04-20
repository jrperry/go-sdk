package iland

type BackupGroupStats struct {
	BackupGroupUID string      `json:"backup_group_uid"`
	Stats          BackupStats `json:"stats"`
}

type BackupStats struct {
	CancelledRuns                int `json:"num_cancelled_runs"`
	FailedRuns                   int `json:"num_failed_runs"`
	SlaViolations                int `json:"num_sla_violations"`
	SuccessfulRuns               int `json:"num_successful_runs"`
	RunningRuns                  int `json:"num_running_runs"`
	AverageRunTimeMs             int `json:"average_run_time_millis"`
	FastestRunTimeMs             int `json:"fastest_run_time_millis"`
	SlowestRunTimeMs             int `json:"slowest_run_time_millis"`
	TotalBytesReadFromSource     int `json:"total_bytes_read_from_source"`
	TotalLogicalBackupSizeBytes  int `json:"total_logical_backup_size_bytes"`
	TotalPhysicalBackupSizeBytes int `json:"total_physical_backup_size_bytes"`
}
