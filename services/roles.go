package services

import (
	cli "cli-app"
	"cli-app/database"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type RoleService struct {
	repo database.RoleRepository
}

func NewRoleService(er database.RoleRepository) *RoleService {
	return &RoleService{
		repo: er,
	}
}

func (rs RoleService) CreateRoleInfo(roleType, name string, createdBy uuid.UUID) error {
	if roleType != "levels" && roleType != "positions" && roleType != "branches" {
		return fmt.Errorf("invalid role type: must be levels, positions, or branches")
	}

	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("name is empty")
	}

	if createdBy == uuid.Nil {
		return fmt.Errorf("missing creator (createdBy is required)")
	}

	return rs.repo.AddRoleInfo(roleType, name, createdBy)
}

func (rs RoleService) GetRoleInfo(roleType string) ([]cli.GetRoleInfo, error) {
	if roleType != "levels" && roleType != "positions" && roleType != "branches" {
		return []cli.GetRoleInfo{}, fmt.Errorf("invalid role type: must be levels, positions, or branches")
	}

	result, err := rs.repo.GetRoleInfo(roleType)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (rs *RoleService) AssignEmployeeRole(in cli.AssignEmployeeRoleDTO) error {
	if in.EmployeeId == uuid.Nil {
		return fmt.Errorf("employee ID is required")
	}
	if in.LevelId == uuid.Nil {
		return fmt.Errorf("level ID is required")
	}
	if in.PositionId == uuid.Nil {
		return fmt.Errorf("position ID is required")
	}
	if in.BranchId == uuid.Nil {
		return fmt.Errorf("branch ID is required")
	}
	if in.CreatedBy == uuid.Nil {
		return fmt.Errorf("created by (user ID) is required")
	}

	if err := rs.repo.AddEmployeeRole(in); err != nil {
		return fmt.Errorf("failed to assign employee role: %w", err)
	}

	return nil
}
