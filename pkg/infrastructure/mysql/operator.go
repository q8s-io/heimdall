package mysql

import (
	"github.com/jinzhu/gorm"
)

// 根据taskID查询
func QueryByTaskID(taskID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("task_id = ?", taskID)
	}
}

func QueryByJobID(jobID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("job_id = ?", jobID)
	}
}
