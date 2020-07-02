package entity

type JobScanner struct {
	TaskID      string `db:"task_id"`
	JobID       string `db:"job_id"`
	JobStatus   string `db:"job_status"`
	JobData     string `db:"job_data"`
	ImageName   string `db:"image_name"`
	ImageDigest string `db:"image_digest"`
	CreateTime  string `db:"create_time"`
	Active      int    `db:"active"`
}
