package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuSalary(empSvc *services.EmployeeService, salSvc *services.SalaryService, currentUserID *uuid.UUID) {
	options := []string{
		"Calculate salary for a month",
		"Calculate salary from CSV",
		"Export Salary to CSV",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Salary Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			salaryMonthFlow(empSvc, salSvc)
		case 2:
			fmt.Println("Calculate salary from CSV ... working on it")
		case 3:
			if err := exportSalaryFlow(salSvc); err != nil {
				fmt.Printf("Failed to export salary: %v\n", err)
			} else {
				fmt.Println("✓ Salary exported successfully!")
			}
		case 4:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
