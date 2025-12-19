package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuSalary(empSvc *services.EmployeeService, salSvc *services.SalaryService, currentUserID *uuid.UUID) {
	options := []string{
		"Calculate Salary for a Month",
		"Export Salary to CSV",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Salary Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-%d): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			if err := calculateSalaryFlow(salSvc); err != nil {
				fmt.Printf("Failed to calculate salary: %v\n", err)
			} else {
				fmt.Println("✓ Salary calculation completed!")
			}
		case 2:
			if err := exportSalaryFlow(salSvc); err != nil {
				fmt.Printf("Failed to export salary: %v\n", err)
			} else {
				fmt.Println("✓ Salary exported successfully!")
			}
		case 3:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
