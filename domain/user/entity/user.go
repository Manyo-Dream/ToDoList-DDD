package entity

import "time"

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	PassWord  string    `json:"pass_word"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
