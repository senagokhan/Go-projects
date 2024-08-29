package models

import "time"

type ProjectModel struct {
	ProjectId          int       `json:"project_id"`
	ProjectName        string    `json:"project_name"`
	ProjectDescription string    `json:"project_description"`
	ProjectStatus      string    `json:"project_status"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	ManagerId          int       `json:"manager_id"`
}
