package models

import "time"

type CommentModel struct {
	CommentId int       `json:"comment_id"`
	TaskId    int       `json:"task_id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
