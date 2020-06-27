package models

type JobAnchoreInfo struct {
	TaskID      string              `json:"task_id"`
	JobID       string              `json:"job_id"`
	JobStatus   string              `json:"job_status"`
	JobData     []map[string]string `json:"job_data"`
	ImageName   string              `json:"image_name"`
	ImageDigest string              `json:"image_digest"`
	CreateTime  string              `json:"create_time"`
}

type JobAnchoreData struct {
	TaskID      string `json:"task_id"`
	JobID       string `json:"job_id"`
	JobStatus   string `json:"job_status"`
	JobData     string `json:"job_data"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	CreateTime  string `json:"create_time"`
}
