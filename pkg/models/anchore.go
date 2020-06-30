package models

import (
	"encoding/json"
)

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
	TaskID      string
	JobID       string
	JobStatus   string
	JobData     string
	ImageName   string
	ImageDigest string
	CreateTime  string
	Active      int
}

type JobAnchoreMsg struct {
	TaskID      string `json:"task_id"`
	JobID       string `json:"job_id"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}

func (jobAnchoreMsg *JobAnchoreMsg) Encode() ([]byte, error) {
	return json.Marshal(jobAnchoreMsg)
}

func (jobAnchoreMsg *JobAnchoreMsg) Length() int {
	encoded, _ := json.Marshal(jobAnchoreMsg)
	return len(encoded)
}

type AnchoreRequestInfo struct {
	ImageName   string `json:"tag"`
	ImageDigest string `json:"digest"`
	CreateTime  string `json:"created_at"`
}
