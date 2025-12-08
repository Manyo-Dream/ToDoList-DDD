package entity

import (
	"errors"
	"time"

	"github.com/manyodream/todolist-ddd/consts"
)

type Task struct {
	ID        uint      `json:"id"`
	UID       uint      `json:"uid"`
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Status    int       `json:"status"`
	Content   string    `json:"content"`
	StartTime int64     `json:"start_time"`
	EndTime   int64     `json:"end_time"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewTask(uid uint, username, title, context string) (*Task, error) {
	if uid == 0 {
		return nil, errors.New("invalid uid")
	}
	if title == "" {
		return nil, errors.New("invalid title")
	}
	now := time.Now()
	return &Task{
		UID:       uid,
		UserName:  username,
		Title:     title,
		Status:    consts.TaskStatusEmunInit,
		Content:   context,
		StartTime: now.Unix(),
		CreateAt:  now,
	}, nil
}

func (t *Task) IsExist() bool {
	return t.ID > 0
}

func (t *Task) Belong2User(uid uint) bool {
	return t.UID == uid
}
