package services

import (
	cli "cli-app"
	"cli-app/database"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ProjectService struct {
	repo database.ProjectRepository
}

func NewProjectService(er database.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: er,
	}
}

func (ps *ProjectService) Create(in cli.AddProject) error {
	if in.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	return ps.repo.Create(in)
}

func (ps *ProjectService) GetAll() ([]cli.ProjectDTO, error) {
	return ps.repo.GetAll()
}

func (ps *ProjectService) GetIdByName(name string) (uuid.UUID, error) {
	if strings.TrimSpace(name) == "" {
		return uuid.UUID{}, fmt.Errorf("name is empty")
	}
	result, err := ps.repo.GetIdByName(name)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (ps *ProjectService) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("project ID cannot be nil")
	}
	return ps.repo.Delete(id)
}

func (ps *ProjectService) AssignMember(in cli.EmployeeProject) error {
	if in.EmployeeId == uuid.Nil {
		return fmt.Errorf("employee ID cannot be nil")
	}
	if in.ProjectId == uuid.Nil {
		return fmt.Errorf("project ID cannot be nil")
	}
	if in.Roles == "" {
		return fmt.Errorf("roles cannot be empty")
	}
	if in.CreatedBy == uuid.Nil {
		return fmt.Errorf("createdBy user ID cannot be nil")
	}

	return ps.repo.AssignEmployee(in)
}

func (ps *ProjectService) GetAssignProject() ([]cli.GetEmployeesProject, error) {
	return ps.repo.GetAssignedEmployees()
}

func (ps *ProjectService) DeleteAssignProject(employeeID uuid.UUID, projectID uuid.UUID) error {
	if employeeID == uuid.Nil {
		return fmt.Errorf("employee ID cannot be nil")
	}
	if projectID == uuid.Nil {
		return fmt.Errorf("project ID cannot be nil")
	}

	return ps.repo.UnassignEmployee(projectID, employeeID)
}

func (ps *ProjectService) CreateTask(in cli.AddTask) (uuid.UUID, error) {
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" {
		return uuid.UUID{}, fmt.Errorf("task name cannot be empty")
	}
	if in.CreatedBy == uuid.Nil {
		return uuid.UUID{}, fmt.Errorf("createdBy user ID cannot be nil")
	}

	taskId, err := ps.repo.CreateTask(in)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("")
	}

	return taskId, nil
}

func (ps *ProjectService) GetAllTasks() ([]cli.TaskDTO, error) {
	return ps.repo.GetAllTasks()
}

func (ps *ProjectService) DeleteTask(id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("task ID cannot be nil")
	}
	return ps.repo.DeleteTask(id)
}

func (ps *ProjectService) AssignTaskProject(in cli.AssignTaskProject) error {
	if in.EmployeeId == uuid.Nil {
		return fmt.Errorf("employee ID cannot be nil")
	}
	if in.ProjectId == uuid.Nil {
		return fmt.Errorf("project ID cannot be nil")
	}
	if in.TaskId == uuid.Nil {
		return fmt.Errorf("task ID cannot be nil")
	}
	if in.CreatedBy == uuid.Nil {
		return fmt.Errorf("createdBy user ID cannot be nil")
	}

	return ps.repo.AssignTaskToProject(in)
}

func (ps *ProjectService) DeleteTaskProject(employeeID uuid.UUID, projectID uuid.UUID) error {
	if employeeID == uuid.Nil {
		return fmt.Errorf("employee ID cannot be nil")
	}
	if projectID == uuid.Nil {
		return fmt.Errorf("project ID cannot be nil")
	}

	return ps.repo.UnassignTaskFromProject(employeeID, projectID)
}

func (ps *ProjectService) GetTaskProject() ([]cli.GetTaskProject, error) {
	return ps.repo.GetTaskAssignments()
}
