// ui/table_helpers.go
package ui

import (
	cli "cli-app"
	"cli-app/components"
	"cli-app/services"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/google/uuid"
)

func selectEmployee(s *services.EmployeeService) string {
	columns := []table.Column{
		{Title: "Email", Width: 25},
		{Title: "Full Name", Width: 20},
		{Title: "Code", Width: 10},
		{Title: "Gender", Width: 10},
		{Title: "Phone", Width: 11},
		{Title: "Dob", Width: 20},
	}

	emps, err := s.GetAllEmployees()
	if err != nil {
		fmt.Printf("Error loading employees: %v\n", err)
		return ""
	}
	if len(emps) == 0 {
		fmt.Println("No employees found.")
		return ""
	}

	rows := employeesToTableRows(emps)
	email := components.Table(columns, rows)

	id, err := s.GetEmployeeId(email)

	if err != nil {
		return ""
	}

	return id.String()
}

func selectProject(p *services.ProjectService) string {
	columns := []table.Column{
		{Title: "Name", Width: 20},
	}

	projects, err := p.GetAll()
	if err != nil {
		fmt.Printf("Error loading projects: %v\n", err)
		return ""
	}
	if len(projects) == 0 {
		fmt.Println("No projects found.")
		return ""
	}

	rows := projectsToTableRows(projects)
	selectedName := components.Table(columns, rows)

	projectId, err := p.GetIdByName(selectedName)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if projectId == uuid.Nil {
		fmt.Println("Project not found")
		return ""
	}

	return projectId.String()
}

func selectTask(p *services.ProjectService) string {
	columns := []table.Column{
		{Title: "ID", Width: 36},
		{Title: "Name", Width: 30},
		{Title: "Created At", Width: 20},
		{Title: "Created By", Width: 36},
	}

	tasks, err := p.GetAllTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return ""
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return ""
	}

	rows := tasksToTableRows(tasks)
	selectedID := components.Table(columns, rows)
	return selectedID
}

func selectRole(roleType string, r *services.RoleService) string {
	columns := []table.Column{
		{Title: "ID", Width: 36},
		{Title: "Name", Width: 30},
	}

	roleList, err := r.GetRoleInfo(roleType)

	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return ""
	}
	if len(roleList) == 0 {
		fmt.Println("No tasks available.")
		return ""
	}

	rows := employeeRoleToTableRows(roleList)
	selectedID := components.Table(columns, rows)
	return selectedID
}

func tasksToTableRows(tasks []cli.TaskDTO) []table.Row {

	rows := make([]table.Row, len(tasks))
	for i, t := range tasks {
		rows[i] = table.Row{
			t.Name,
			t.Notes,
			strconv.Itoa(t.WorkingTime),
		}
	}
	return rows
}

func employeesToTableRows(emps []cli.EmployeeDTO) []table.Row {
	rows := make([]table.Row, len(emps))
	for i, e := range emps {
		rows[i] = table.Row{
			e.Email,
			e.FullName,
			e.Code,
			string(e.Gender),
			e.Phone,
			e.Dob.Format("2006-01-02"),
		}
	}
	return rows
}

func projectsToTableRows(projs []cli.ProjectDTO) []table.Row {
	rows := make([]table.Row, len(projs))
	for i, p := range projs {
		rows[i] = table.Row{
			p.Name,
		}
	}
	return rows
}

func employeeRoleToTableRows(role []cli.GetRoleInfo) []table.Row {
	rows := make([]table.Row, len(role))
	for i, p := range role {
		rows[i] = table.Row{
			p.Id.String(),
			p.Name,
		}
	}

	return rows
}

func selectEmailEmployee(s *services.EmployeeService) string {
	columns := []table.Column{
		{Title: "Email", Width: 25},
		{Title: "Full Name", Width: 20},
		{Title: "Code", Width: 10},
		{Title: "Gender", Width: 10},
		{Title: "Phone", Width: 11},
		{Title: "Dob", Width: 20},
	}

	emps, err := s.GetAllEmployees()
	if err != nil {
		fmt.Printf("Error loading employees: %v\n", err)
		return ""
	}
	if len(emps) == 0 {
		fmt.Println("No employees found.")
		return ""
	}

	rows := employeesToTableRows(emps)
	return components.Table(columns, rows)
}
