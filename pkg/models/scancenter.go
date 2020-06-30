package models

type ImageRequestInfo struct {
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}

type ImageVulnData struct {
	TaskID      string
	TaskStatus  string
	ImageName   string
	ImageDigest string
	CreateTime  string
	VulnData    []map[string]string
}

type TaskImageScanInfo struct {
	TaskID      string `json:"task_id"`
	TaskStatus  string `json:"task_status"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	CreateTime  string `json:"create_time"`
}

type TaskImageScanData struct {
	TaskID      string `db:"task_id"`
	TaskStatus  string `db:"task_status"`
	ImageName   string `db:"image_name"`
	ImageDigest string `db:"image_digest"`
	CreateTime  string `db:"create_time"`
	Active      int
}
