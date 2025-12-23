package database

import (
	cli "cli-app"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type RoleRepository interface {
	GetRoleInfo(roleType string) ([]cli.GetRoleInfo, error)
	AddRoleInfo(roleType string, name string, createdBy uuid.UUID) error
	AddEmployeeRole(in cli.AssignEmployeeRoleDTO) error
}

type roleRepo struct {
	db *sql.DB
}

func InitRoleRepo(db *sql.DB) RoleRepository {
	return &roleRepo{db: db}
}

// Get Levels, Positions and Branches informations list.
// Depends on an input type want to GET
func (rr *roleRepo) GetRoleInfo(roleType string) ([]cli.GetRoleInfo, error) {
	var query string
	switch roleType {
	case "levels":
		query = getLevelInfo
	case "positions":
		query = getPositionInfo
	case "branches":
		query = getBranchesInfo
	default:
		return nil, fmt.Errorf("invalid role type: %s", roleType)
	}

	rows, err := rr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query roles: %w", err)
	}
	defer rows.Close()

	var result []cli.GetRoleInfo

	for rows.Next() {
		var item cli.GetRoleInfo
		if err := rows.Scan(&item.Id, &item.Name); err != nil {
			return nil, fmt.Errorf("failed to scan role row: %w", err)
		}
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

// func (rr *roleRepo) GetEmployeeRole(email string) {}
// func (rr *roleRepo) GetDefaultSalary()            {}

func (rr *roleRepo) AddRoleInfo(roleType string, name string, createdBy uuid.UUID) error {
	var query string

	switch roleType {
	case "levels":
		query = addLevelsInfo
	case "positions":
		query = addPositionsInfo
	case "branches":
		query = addBranchesInfo
	default:
		return fmt.Errorf("invalid role type: %s", roleType)
	}

	row, err := rr.db.Exec(query, name, createdBy)

	if err != nil {
		return fmt.Errorf("failed to insert role: %w", err)
	}

	affected, err := row.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("no row inserted for role type %s with name %s", roleType, name)
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (rr *roleRepo) AddEmployeeRole(in cli.AssignEmployeeRoleDTO) error {
	row, err := rr.db.Exec(assignEmployeRole, in.EmployeeId, in.LevelId, in.PositionId, in.BranchId, in.CreatedBy)

	if err != nil {
		return fmt.Errorf("failed to insert role: %w", err)
	}

	affected, err := row.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("no employee role inserted for employee %s (possible duplicate or constraint violation)", in.EmployeeId)
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

// func (rr *roleRepo) AddDefaultSalary() {}
