package mysql

import "github.com/jinzhu/gorm"

// 通用方法封装

// 根据ID查询
func QuerytByTaskID(taskID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("task_id = ?", taskID)
	}
}
