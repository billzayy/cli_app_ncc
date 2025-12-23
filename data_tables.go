package cli_app

import (
	"time"

	"github.com/google/uuid"
)

type Branches struct {
	Id          uuid.UUID `json:"id " db:"id"`
	Name        string    `json:"name" db:"name"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

type Levels struct {
	Id          uuid.UUID `json:"id " db:"id"`
	Name        string    `json:"name" db:"name"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

type Positions struct {
	Id          uuid.UUID `json:"id " db:"id"`
	Name        string    `json:"name" db:"name"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Employees Struct
type Employees struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	FullName    string    `json:"full_name" db:"full_name"`
	Code        string    `json:"code" db:"code"`
	Gender      Gender    `json:"gender" db:"gender"`
	Phone       string    `json:"phone" db:"phone"`
	Dob         time.Time `json:"dob" db:"dob"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Identifiers Struct
type Identifiers struct {
	EmployeeId  uuid.UUID `json:"employee_id" db:"employee_id"`
	CardNumber  string    `json:"card_number" db:"card_number"`
	DateOfIssue time.Time `json:"date_of_issue" db:"date_of_issue"`
	IssuedBy    string    `json:"issued_by" db:"issued_by"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Banks Struct
type Banks struct {
	EmployeeId    uuid.UUID `json:"employee_id" db:"employee_id"`
	AccountNumber string    `json:"account_number" db:"account_number"`
	CreatedTime   time.Time `json:"created_time" db:"created_time"`
	CreatedBy     uuid.UUID `json:"created_by" db:"created_by"`
}

// Address Struct
type Address struct {
	Id         uuid.UUID `json:"id" db:"id"`
	Origin     string    `json:"origin" db:"origin"`
	Residence  string    `json:"residence" db:"residence"`
	Current    string    `json:"current" db:"current"`
	EmployeeId uuid.UUID `json:"employee_id" db:"employee_id"`
}

// Benefits Struct
type Benefits struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Projects Struct
type Projects struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Notes       string    `json:"note" db:"note"`
	WorkingTime int       `json:"working_time" db:"working_time"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Tasks Struct
type Tasks struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Departments Struct
type Departments struct {
	DepartMentHeadId uuid.UUID `json:"department_head_id" db:"department_head_id"`
	Name             string    `json:"name" db:"name"`
	Address          string    `json:"address" db:"address"`
	CreatedTime      time.Time `json:"created_time" db:"created_time"`
	CreatedBy        uuid.UUID `json:"created_by" db:"created_by"`
}

// Default salary struct
type LevelDefault struct {
	LevelId     uuid.UUID `json:"level_id" db:"level_id"`
	ValueType   string    `json:"value_type" db:"value_type"`
	Amount      int       `json:"amount" db:"amount"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}

// Employee - Project table
type EmployeeProject struct {
	ProjectId   uuid.UUID `json:"project_id"`
	EmployeeId  uuid.UUID `json:"employee_id"`
	Roles       string    `json:"roles"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
}
