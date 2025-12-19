package ui

import (
	"cli-app/services"
	"fmt"
)

func readInt(prompt string, max int) int {
	for {
		fmt.Print(prompt)
		var input int
		_, err := fmt.Scan(&input)
		if err != nil || input < 1 || input > max {
			fmt.Println("Invalid input. Please enter a number between 1 and", max)
			// Clear remaining input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return input
	}
}

func exportOnChange(s *services.EmployeeService) {
	if err := s.ExportEmployee(employeeCSVPath); err != nil {
		fmt.Printf("Warning: Failed to export employees: %v\n", err)
	}
}
