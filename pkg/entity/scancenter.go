package entity

// type TaskImageScan struct {
// 	TaskID      string `db:"task_id"`
// 	TaskStatus  string `db:"task_status"`
// 	ImageName   string `db:"image_name"`
// 	ImageDigest string `db:"image_digest"`
// 	CreateTime  string `db:"create_time"`
// 	Active      int    `db:"active"`
// }

type TaskImageScan struct {
	ID          uint   `gorm:"AUTO_INCREMENT;not null"`
	TaskID      string `gorm:"type:varchar(32);not null;primary_key"`
	TaskStatus  string `gorm:"type:enum"`
	ImageName   string `gorm:"type:varchar(255)"`
	ImageDigest string `gorm:"type:varchar(255)"`
	CreateTime  string `gorm:"column:create_time"`
	Active      int    `gorm:"type:tinyint"`
}

// table: image_vuln
type ImageVuln struct {
	TaskImageScan
}

func (ImageVuln) TableName() string {
	return "image_vuln"
}
