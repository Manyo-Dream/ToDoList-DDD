package entity

import "time"

type Task struct {
	ID        uint      `json:"id"`
	UID       uint      `json:"uid"`
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Status    int       `json:"status"`
	Content   string    `json:"content"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
