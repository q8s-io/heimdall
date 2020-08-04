package model

type ImageRequestInfo struct {
	ImageName   string `json:"image_name" example:"photon:2.0-20190511"`
	ImageDigest string `json:"image_digest" example:" "`
}

type ImageVulnInfo struct {
	TaskID      string                   `json:"task_id"`
	TaskStatus  string                   `json:"task_status"`
	ImageName   string                   `json:"image_name"`
	ImageDigest string                   `json:"image_digest"`
	CreateTime  string                   `json:"create_time"`
	VulnData    []map[string]interface{} `json:"vuln_data"`
}

type TaskImageScanInfo struct {
	TaskID      string `json:"task_id"`
	TaskStatus  string `json:"task_status"`
	ImageName   string `json:"image_name"`
	ImageDigest string `json:"image_digest"`
	CreateTime  string `json:"create_time"`
}

type Result struct {
	Code    int         `json:"code" example:"0"`
	Message string      `json:"message" example:""`
	Data    interface{} `json:"data" `
}

type VulnerabilityMsg []byte

func (vuln VulnerabilityMsg) Encode() ([]byte, error) {
	return vuln, nil
}

func (vuln VulnerabilityMsg) Length() int {
	return len(vuln)
}
