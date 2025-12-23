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

func assignMemberFlow(empSvc *services.EmployeeService, projSvc *services.ProjectService, createdBy uuid.UUID) error {
	employeeID := selectEmployee(empSvc)
	if employeeID == "" {
		return fmt.Errorf("no employee selected")
	}

	projectID := selectProject(projSvc)
	if projectID == "" {
		return fmt.Errorf("no project selected")
	}

	var roles string
	fmt.Print("Enter roles (comma-separated, e.g., Developer, Tester): ")
	fmt.Scanln(&roles)
	roles = strings.TrimSpace(roles)
	if roles == "" {
		return fmt.Errorf("roles cannot be empty")
	}

	assignment := cli.EmployeeProject{
		EmployeeId: uuid.MustParse(employeeID),
		ProjectId:  uuid.MustParse(projectID),
		Roles:      roles,
		CreatedBy:  createdBy,
	}

	return projSvc.AssignMember(assignment)
}

func viewAllAssignmentsFlow(projSvc *services.ProjectService) {
	assignments, err := projSvc.GetAssignProject()
	if err != nil {
		fmt.Printf("Error fetching assignments: %v\n", err)
		return
	}

	if len(assignments) == 0 {
		fmt.Println("No project assignments found.")
		return
	}

	columns := []table.Column{
		{Title: "Employee", Width: 20},
		{Title: "Project", Width: 20},
		{Title: "Roles", Width: 30},
	}

	rows := make([]table.Row, len(assignments))
	for i, a := range assignments {
		rows[i] = table.Row{a.EmployeeName, a.Project, a.Role}
	}

	fmt.Println("\nCurrent Project Assignments:")
	components.Table(columns, rows)
}

func removeMemberFlow(eSvc *services.EmployeeService, projSvc *services.ProjectService) error {
	eId, pId := showTableEP(eSvc, projSvc)

	empID, err := uuid.Parse(strings.TrimSpace(eId))
	if err != nil {
		return fmt.Errorf("invalid employee ID: %w", err)
	}
	projID, err := uuid.Parse(strings.TrimSpace(pId))
	if err != nil {
		return fmt.Errorf("invalid project ID: %w", err)
	}

	return projSvc.DeleteAssignProject(empID, projID)
}
