package ui

import (
	"cli-app/services"
	"fmt"
	"strings"
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
