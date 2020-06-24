package models

type AnalyzerJobInfo struct {
	TaskID      string `json:"task_id"`
	JobID       string `json:"job_id"`
	JobStatus   string `json:"job_status"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	CreateTime  string `json:"create_time"`
}
