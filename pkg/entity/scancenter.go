package entity

type TaskImageScan struct {
	TaskID      string `db:"task_id"`
	TaskStatus  string `db:"task_status"`
	ImageName   string `db:"image_name"`
	ImageDigest string `db:"image_digest"`
	CreateTime  string `db:"create_time"`
	Active      int    `db:"active"`
}