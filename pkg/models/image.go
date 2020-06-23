package models

type ImageInfoRequest struct {
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}

type ImageVulnInfo struct {
	ImageName      string `json:"image_name"`
	ImageDigest    string `json:"image_digest"`
	TaskID         string `json:"task_id"`
	TaskStatus     string `json:"task_status"`
	AnalyzerTaskID string `json:"analyzer_task_id"`
	AnchoreTaskID  string `json:"anchore_task_id"`
	AnchoreData    string `json:"anchore_data"`
}
