package entity

type JobScanner struct {
	ID          uint   `gorm:"AUTO_INCREMENT;not null"`
	TaskID      string `gorm:"type:varchar(32);not null;primary_key"`
	JobID       string `gorm:"type:varchar(32);not null"`
	JobStatus   string `gorm:"type:enum"`
	JobData     string `gorm:"type:longtext"`
	ImageName   string `gorm:"type:varchar(255)"`
	ImageDigest string `gorm:"type:varchar(255)"`
	CreateTime  string `gorm:"column:create_time"`
	Active      int    `gorm:"type:tinyint"`
}

// table: job_anchore
type JobAnchore struct {
	JobScanner
}

func (JobAnchore) TableName() string {
	return "job_anchore"
}

// table: job_trivy
type JobTrivy struct {
	JobScanner
}

func (JobTrivy) TableName() string {
	return "job_trivy"
}

// table: job_clair
type JobClair struct {
	JobScanner
}

func (JobClair) TableName() string {
	return "job_clair"
}
