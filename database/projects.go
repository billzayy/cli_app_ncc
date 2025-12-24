package database

import (
	cli "cli-app"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type ProjectRepository interface {
	Create(in cli.AddProject) error
	GetAll() ([]cli.ProjectDTO, error)
	GetIdByName(name string) (uuid.UUID, error)
	Delete(id uuid.UUID) error

	AssignEmployee(in cli.EmployeeProject) error
	GetAssignedEmployees() ([]cli.GetEmployeesProject, error)
	UnassignEmployee(projectID, employeeID uuid.UUID) error

	CreateTask(in cli.AddTask) (uuid.UUID, error)
	GetAllTasks() ([]cli.TaskDTO, error)
	DeleteTask(id uuid.UUID) error

	AssignTaskToProject(in cli.AssignTaskProject) error
	UnassignTaskFromProject(employeeID, projectID uuid.UUID) error
	GetTaskAssignments() ([]cli.GetTaskProject, error)
}

type projectRepo struct {
	db *sql.DB
}

func InitProjectRepo(db *sql.DB) ProjectRepository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(in cli.AddProject) error {
	row, err := r.db.Exec(createProject,
		in.Name,
		in.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}

	affected, err := row.RowsAffected()

	if affected == 0 {
		return fmt.Errorf("no project is added")
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)
	return nil
}

func (pr *projectRepo) GetAll() ([]cli.ProjectDTO, error) {
	rows, err := pr.db.Query(getAllProjects)

	if err != nil {
		return []cli.ProjectDTO{}, err
	}
	defer rows.Close()

	var result []cli.ProjectDTO
	for rows.Next() {
		var temp cli.ProjectDTO
		err := rows.Scan(
			&temp.Name)

		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}
	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (pr *projectRepo) GetIdByName(name string) (uuid.UUID, error) {
	rows, err := pr.db.Query(getProjectIdByName, name)

	if err != nil {
		return uuid.UUID{}, err
	}
	defer rows.Close()

	var result uuid.UUID

	for rows.Next() {
		var temp string
		err := rows.Scan(&temp)

		if err != nil {
			return uuid.UUID{}, err
		}

		result = uuid.MustParse(temp)
	}

	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (pr *projectRepo) Delete(id uuid.UUID) error {
	rows, err := pr.db.Exec(deleteProject, id)

	if err != nil {
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("project not found")
	}
	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) AssignEmployee(in cli.EmployeeProject) error {
	rows, err := pr.db.Exec(assignProject, in.ProjectId, in.EmployeeId, in.Roles, in.CreatedBy)

	if err != nil {
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("can not assign employeee to project")
	}
	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) GetAssignedEmployees() ([]cli.GetEmployeesProject, error) {
	rows, err := pr.db.Query(getAssignProject)

	if err != nil {
		return []cli.GetEmployeesProject{}, err
	}
	defer rows.Close()

	var result []cli.GetEmployeesProject
	var temp cli.GetEmployeesProject

	for rows.Next() {
		err := rows.Scan(&temp.EmployeeId, &temp.ProjectId, &temp.EmployeeName, &temp.Project, &temp.Role)

		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}
	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (pr *projectRepo) UnassignEmployee(projectId uuid.UUID, employeeId uuid.UUID) error {
	row, err := pr.db.Exec(deleteAssignProject, projectId, employeeId)

	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found data")
	}
	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) CreateTask(in cli.AddTask) (uuid.UUID, error) {
	rows, err := pr.db.Query(createTasks, in.Name, in.Notes, in.WorkingTime, in.CreatedBy)

	if err != nil {
		return uuid.UUID{}, err
	}

	defer rows.Close()

	var id uuid.UUID

	for rows.Next() {
		var temp string
		err := rows.Scan(&temp)

		if err != nil {
			return uuid.UUID{}, err
		}

		id = uuid.MustParse(temp)
	}

	if id == uuid.Nil {
	}

	return id, nil
}

func (pr *projectRepo) GetAllTasks() ([]cli.TaskDTO, error) {
	rows, err := pr.db.Query(getAllTasks)

	if err != nil {
		return []cli.TaskDTO{}, err
	}
	defer rows.Close()

	var result []cli.TaskDTO

	for rows.Next() {
		var temp cli.TaskDTO
		err := rows.Scan(&temp.Name, &temp.Notes, &temp.WorkingTime)

		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}
	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (pr *projectRepo) DeleteTask(id uuid.UUID) error {
	row, err := pr.db.Exec(deleteTasks, id)

	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found task")
	}
	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) AssignTaskToProject(in cli.AssignTaskProject) error {
	row, err := pr.db.Exec(addTaskToProject, in.ProjectId, in.TaskId, in.EmployeeId, in.CreatedBy)

	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("Assign Task to project failed")
	}
	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) UnassignTaskFromProject(eId uuid.UUID, pId uuid.UUID) error {
	row, err := pr.db.Exec(deleteTaskToProject, eId, pId)

	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found task assignment to project")
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}

func (pr *projectRepo) GetTaskAssignments() ([]cli.GetTaskProject, error) {
	rows, err := pr.db.Query(getTaskToProject)

	if err != nil {
		return []cli.GetTaskProject{}, err
	}
	defer rows.Close()

	var result []cli.GetTaskProject
	var temp cli.GetTaskProject
	for rows.Next() {
		err := rows.Scan(&temp.EmployeeName, &temp.ProjectName, &temp.TaskName)

		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}
	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}
