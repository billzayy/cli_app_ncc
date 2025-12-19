// ui/table_helpers.go
package ui

import (
	cli "cli-app"
	"cli-app/components"
	"cli-app/services"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

func selectEmployee(s *services.EmployeeService) string {
	columns := []table.Column{
		{Title: "ID", Width: 36},
		{Title: "Email", Width: 25},
		{Title: "Full Name", Width: 20},
		{Title: "Code", Width: 10},
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
	selectedID := components.Table(columns, rows)
	return selectedID
}

func selectProject(p *services.ProjectService) string {
	columns := []table.Column{
		{Title: "ID", Width: 36},
		{Title: "Name", Width: 20},
		{Title: "Notes", Width: 20},
		{Title: "Working Time", Width: 12},
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
	selectedID := components.Table(columns, rows)
	return selectedID
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

func tasksToTableRows(tasks []cli.Tasks) []table.Row {
	rows := make([]table.Row, len(tasks))
	for i, t := range tasks {
		rows[i] = table.Row{
			t.Id.String(),
			t.Name,
			t.CreatedTime.Format("2006-01-02 15:04"),
			t.CreatedBy.String(),
		}
	}
	return rows
}

func employeesToTableRows(emps []cli.Employees) []table.Row {
	rows := make([]table.Row, len(emps))
	for i, e := range emps {
		rows[i] = table.Row{
			e.Id.String(),
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

func projectsToTableRows(projs []cli.Projects) []table.Row {
	rows := make([]table.Row, len(projs))
	for i, p := range projs {
		rows[i] = table.Row{
			p.Id.String(),
			p.Name,
			p.Notes,
			strconv.Itoa(p.WorkingTime),
		}
	}
	return rows
}
