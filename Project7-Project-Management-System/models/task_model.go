package models

import "time"

type TaskModel struct {
	TaskId          int       `json:"task_id"`
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_description"`
	TaskStatus      string    `json:"task_status"`
	TaskPriority    string    `json:"task_priority"`
	ProjectId       int       `json:"project_id"`
	AssignedTo      int       `json:"assigned_to"`
	AssignedAt      time.Time `json:"assigned_at"`
	DueAt           time.Time `json:"due_at"`
}
