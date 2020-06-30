package models

type ImageRequestInfo struct {
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}

type ImageVulnInfo struct {
	TaskID      string `json:"task_id"`
	TaskStatus  string `json:"task_status"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	CreateTime  string `json:"create_time"`
}

type ImageVulnData struct {
	TaskID      string
	TaskStatus  string
	ImageName   string
	ImageDigest string
	CreateTime  string
	Active      int
}
