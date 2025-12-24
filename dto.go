package cli_app

import (
	"time"

	"github.com/google/uuid"
)

type Gender string
type ProjectType string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

const (
	NormalTime ProjectType = "Normal working hours"
	OverTime   ProjectType = "Overtime"
)

// Employee DTO Struct
type EmployeeDTO struct {
	Email    string    `json:"email" db:"email"`
	FullName string    `json:"full_name" db:"full_name"`
	Code     string    `json:"code" db:"code"`
	Gender   Gender    `json:"gender" db:"gender"`
	Phone    string    `json:"phone" db:"phone"`
	Dob      time.Time `json:"dob" db:"dob"`
}

type ProjectDTO struct {
	Name string
}

type TaskDTO struct {
	Name        string
	Notes       string
	WorkingTime int
}

type GetEmployeesProject struct {
	EmployeeId   uuid.UUID
	ProjectId    uuid.UUID
	EmployeeName string
	Project      string
	Role         string
}

type GetTaskProject struct {
	EmployeeName string
	ProjectName  string
	TaskName     string
}

type TimeAndAmount struct {
	Name        string
	WorkingTime int
	Amount      int
	Month       int
}

type GetRoleInfo struct {
	Id   uuid.UUID
	Name string
}

type AssignEmployeeRoleDTO struct {
	EmployeeId uuid.UUID
	LevelId    uuid.UUID
	PositionId uuid.UUID
	BranchId   uuid.UUID
	CreatedBy  uuid.UUID
}
