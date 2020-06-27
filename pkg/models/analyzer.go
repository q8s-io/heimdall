package models

import (
	"encoding/json"
)

type JobAnalyzerInfo struct {
	TaskID      string   `json:"task_id"`
	JobID       string   `json:"job_id"`
	JobStatus   string   `json:"job_status"`
	ImageName   string   `json:"image_name"`
	ImageDigest string   `json:"image_digest"`
	ImageLayers []string `json:"image_layers"`
	CreateTime  string   `json:"create_time"`
}

type JobAnalyzerData struct {
	TaskID      string `json:"task_id"`
	JobID       string `json:"job_id"`
	JobStatus   string `json:"job_status"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	ImageLayers string `json:"image_layers"`
	CreateTime  string `json:"create_time"`
}

type JobAnalyzerMsg struct {
	TaskID    string `json:"task_id"`
	JobID     string `json:"job_id"`
	ImageName string `json:"image_name"`
}

func (jobAnalyzerMsg *JobAnalyzerMsg) Encode() ([]byte, error) {
	return json.Marshal(jobAnalyzerMsg)
}

func (jobAnalyzerMsg *JobAnalyzerMsg) Length() int {
	encoded, _ := json.Marshal(jobAnalyzerMsg)
	return len(encoded)
}
