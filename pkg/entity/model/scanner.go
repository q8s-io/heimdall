package model

import (
	"encoding/json"
)

type JobScannerInfo struct {
	TaskID      string              `json:"task_id"`
	JobID       string              `json:"job_id"`
	JobStatus   string              `json:"job_status"`
	JobData     []map[string]string `json:"job_data"`
	ImageName   string              `json:"image_name"`
	ImageDigest string              `json:"image_digest"`
	CreateTime  string              `json:"create_time"`
}

type JobScannerMsg struct {
	TaskID      string `json:"task_id"`
	JobID       string `json:"job_id"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}

func (jobScannerMsg *JobScannerMsg) Encode() ([]byte, error) {
	return json.Marshal(jobScannerMsg)
}

func (jobScannerMsg *JobScannerMsg) Length() int {
	encoded, _ := json.Marshal(jobScannerMsg)
	return len(encoded)
}
