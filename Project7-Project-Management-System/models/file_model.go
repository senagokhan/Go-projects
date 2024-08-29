package models

import "time"

type FileModel struct {
	FileId     int       `json:"file_id"`
	FileName   string    `json:"file_name"`
	FilePath   string    `json:"file_path"`
	TaskId     int       `json:"task_id"`
	UploadedBy int       `json:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at"`
}
