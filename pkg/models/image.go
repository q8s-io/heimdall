package models

type ImageInfo struct {
	TaskID      string `json:"task_id"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
}
