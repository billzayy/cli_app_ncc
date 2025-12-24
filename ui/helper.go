package ui

import (
	"cli-app/components"
	"cli-app/services"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
)

func readInt(prompt string, max int) int {
	for {
		fmt.Print(prompt)
		var input int
		_, err := fmt.Scan(&input)
		if err != nil || input < 1 || input > max {
			fmt.Printf("Invalid input. Please enter a number between 1 and %d\n", max)
			continue
		}
		// ← ADD THIS: consume the leftover newline
		var newline string
		fmt.Scanln(&newline)
		return input
	}
}

func exportOnChange(s *services.EmployeeService) {
	if err := s.ExportEmployee(employeeCSVPath); err != nil {
		fmt.Printf("Warning: Failed to export employees: %v\n", err)
	}

	fmt.Println("✓ Export successul!")
}

func validateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	checkAtSign := strings.Contains(email, "@")

	if !checkAtSign {
		return "", fmt.Errorf("Missing @ for email format")
	}

	domain := strings.Split(email, "@")[1]
	splitDomain := strings.Split(domain, ".")

	if len(splitDomain) == 0 {
		return "", fmt.Errorf("Missing domain for email")
	}

	if splitDomain[0] != "gmail" || splitDomain[1] != "com" {
		return "", fmt.Errorf("Missing domain name format")
	}

	return email, nil
}

func showTableE(s *services.EmployeeService) string {
	var employeeId string
	employeeTableColumn := []table.Column{
		{Title: "Email", Width: 20},
		{Title: "FullName", Width: 12},
		{Title: "Code", Width: 10},
		{Title: "Gender", Width: 12},
		{Title: "Phone", Width: 11},
		{Title: "Dob", Width: 10},
	}

	outputE, err := s.GetAllEmployees()

	if err != nil {
		fmt.Println(err)
	}

	if len(outputE) > 0 {
		employeeRows := employeesToTableRows(outputE)

		employeeName := components.Table(employeeTableColumn, employeeRows)

		outId, err := s.GetEmployeeId(employeeName)

		if err != nil {
			fmt.Println(err)
		}

		employeeId = outId.String()
	}

	return employeeId
}

func showTableP(p *services.ProjectService) string {
	var projectId string

	projectTableColumn := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Notes", Width: 15},
		{Title: "Working Time", Width: 12},
	}

	outputP, err := p.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	if len(outputP) > 0 {
		projectRows := projectsToTableRows(outputP)
		projecName := components.Table(projectTableColumn, projectRows)

		id, err := p.GetIdByName(projecName)

		if err != nil {
			fmt.Println(err)
		}

		projectId = id.String()
	}

	return projectId
}

func showTableEP(s *services.EmployeeService, p *services.ProjectService) (string, string) {
	return showTableE(s), showTableP(p)
}
