package models

import (
	"encoding/json"
)

type JobImageAnalyzerInfo struct {
	TaskID      string   `json:"task_id"`
	JobID       string   `json:"job_id"`
	JobStatus   string   `json:"job_status"`
	ImageName   string   `json:"image_name"`
	ImageDigest string   `json:"image_digest"`
	ImageLayers []string `json:"image_layers"`
	CreateTime  string   `json:"create_time"`
}

type JobImageAnalyzerData struct {
	TaskID      string
	JobID       string
	JobStatus   string
	ImageName   string
	ImageDigest string
	ImageLayers string
	CreateTime  string
	Active      int
}

type JobImageAnalyzerMsg struct {
	TaskID    string `json:"task_id"`
	JobID     string `json:"job_id"`
	ImageName string `json:"image_name"`
}

func (jobAnalyzerMsg *JobImageAnalyzerMsg) Encode() ([]byte, error) {
	return json.Marshal(jobAnalyzerMsg)
}

func (jobAnalyzerMsg *JobImageAnalyzerMsg) Length() int {
	encoded, _ := json.Marshal(jobAnalyzerMsg)
	return len(encoded)
}
