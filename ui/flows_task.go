package ui

import (
	cli "cli-app"
	"cli-app/components"
	"cli-app/services"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/google/uuid"
)

func createTaskFlow(p *services.ProjectService, createdBy uuid.UUID) error {
	var name string
	fmt.Print("Enter task name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("task name cannot be empty")
	}
	return p.CreateTask(name, createdBy)
}

func listAllTasksFlow(p *services.ProjectService) {
	tasks, err := p.GetAllTasks()
	if err != nil {
		fmt.Printf("Error fetching tasks: %v\n", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	columns := []table.Column{
		{Title: "ID", Width: 36},
		{Title: "Name", Width: 25},
		{Title: "Created At", Width: 20},
		{Title: "Created By", Width: 36},
	}

	rows := tasksToTableRows(tasks)
	fmt.Println("\nAll Tasks:")
	components.Table(columns, rows) // Displays table and waits for Enter
}

func deleteTaskFlow(p *services.ProjectService) error {
	taskID := selectTask(p)
	if taskID == "" {
		return fmt.Errorf("no task selected")
	}
	return p.DeleteTask(uuid.MustParse(taskID))
}

func assignTaskFlow(empSvc *services.EmployeeService, projSvc *services.ProjectService, createdBy uuid.UUID) error {
	employeeID := selectEmployee(empSvc)
	if employeeID == "" {
		return fmt.Errorf("no employee selected")
	}
	projectID := selectProject(projSvc)
	if projectID == "" {
		return fmt.Errorf("no project selected")
	}
	taskID := selectTask(projSvc)
	if taskID == "" {
		return fmt.Errorf("no task selected")
	}

	assignment := cli.AssignTaskProject{
		EmployeeId: uuid.MustParse(employeeID),
		ProjectId:  uuid.MustParse(projectID),
		TaskId:     uuid.MustParse(taskID),
		CreatedBy:  createdBy,
	}

	return projSvc.AssignTaskProject(assignment)
}

func viewTaskAssignmentsFlow(p *services.ProjectService) {
	assignments, err := p.GetTaskProject()
	if err != nil {
		fmt.Printf("Error fetching assignments: %v\n", err)
		return
	}
	if len(assignments) == 0 {
		fmt.Println("No task assignments found.")
		return
	}

	// Assuming GetTaskProject() returns a slice with fields like EmployeeName, ProjectName, TaskName
	columns := []table.Column{
		{Title: "Employee", Width: 20},
		{Title: "Project", Width: 20},
		{Title: "Task", Width: 25},
	}

	rows := make([]table.Row, len(assignments))
	for i, a := range assignments {
		rows[i] = table.Row{a.EmployeeName, a.ProjectName, a.TaskName} // Adjust field names
	}

	fmt.Println("\nCurrent Task Assignments:")
	components.Table(columns, rows)
}

func removeTaskAssignmentFlow(empSvc *services.EmployeeService, projSvc *services.ProjectService) error {
	// Optional: show current assignments first?
	employeeID := selectEmployee(empSvc)
	if employeeID == "" {
		return fmt.Errorf("no employee selected")
	}
	projectID := selectProject(projSvc)
	if projectID == "" {
		return fmt.Errorf("no project selected")
	}

	return projSvc.DeleteTaskProject(uuid.MustParse(employeeID), uuid.MustParse(projectID))
}
