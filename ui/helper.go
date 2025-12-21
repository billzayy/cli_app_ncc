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
			fmt.Println("Invalid input. Please enter a number between 1 and", max)
			var discard string
			fmt.Scan(&discard)
			continue
		}
		return input
	}
}

func exportOnChange(s *services.EmployeeService) {
	if err := s.ExportEmployee(employeeCSVPath); err != nil {
		fmt.Printf("Warning: Failed to export employees: %v\n", err)
	}

	fmt.Println("Export successul!")
}

func validateEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	return email, nil
}
