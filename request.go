package cli_app

import (
	"github.com/google/uuid"
	"time"
)

type InputCreateEmployee struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Code     string `json:"code"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
}

type AddEmployee struct {
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	Code       string    `json:"code"`
	Gender     string    `json:"gender"`
	Phone      string    `json:"phone"`
	Dob        time.Time `json:"dob"`
	LevelId    string    `json:"levelId"`
	PositionId string
	BranchId   string
	CreatedBy  uuid.UUID `json:"created_by"`
}

type InputAddProject struct {
	Name        string `json:"name"`
	Notes       string `json:"notes"`
	WorkingTime string `json:"working_time"`
}

type AddProject struct {
	Name        string    `json:"name"`
	Notes       string    `json:"notes"`
	WorkingTime int       `json:"working_time"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

type AddTask struct {
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}

type AssignEmployeeProject struct {
	ProjectId  uuid.UUID `json:"project_id"`
	EmployeeId uuid.UUID `json:"employee_id"`
	Roles      string    `json:"roles"`
}

type AssignTaskProject struct {
	ProjectId  uuid.UUID
	TaskId     uuid.UUID
	EmployeeId uuid.UUID
	CreatedBy  uuid.UUID
}
