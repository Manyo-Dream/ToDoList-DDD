package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"column:uid;index;not null"`
	Title     string `gorm:"column:title;not null"`
	Staust    int    `gorm:"column:status;defult:0;not null"`
	Content   string `gorm:"column:content"`
	StartTime string `gorm:"column:start_time;defult:0"`
	EndTime   string `gorm:"column:end_time;defult:0"`
}
