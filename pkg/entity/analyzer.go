package entity

type JobImageAnalyzer struct {
	ID          uint   `gorm:"AUTO_INCREMENT;not null"`
	TaskID      string `gorm:"type:varchar(32);not null;primary_key"`
	JobID       string `gorm:"type:varchar(32);not null"`
	JobStatus   string `gorm:"type:enum"`
	ImageName   string `gorm:"type:varchar(255)"`
	ImageDigest string `gorm:"type:varchar(255)"`
	ImageLayers string `gorm:"type:longtext"`
	CreateTime  string `gorm:"column:create_time"`
	Active      int    `gorm:"type:tinyint"`
}

// table: job_analyzer
type JobAnalyzer struct {
	JobImageAnalyzer
}

func (JobAnalyzer) TableName() string {
	return "job_analyzer"
}
