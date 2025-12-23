package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

const employeeCSVPath = "./csv/employees.csv"

func MenuEmployee(s *services.EmployeeService, currentUserID uuid.UUID) {
	options := []string{
		"Create New Employee",
		"Read Employee",
		"Read All Employees",
		"Update Employee",
		"Delete Employee",
		"Import Employees from CSV",
		"Export Employeest to CSV",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Employee Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-8): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			if err := createEmployeeFlow(s, currentUserID); err != nil {
				fmt.Printf("Failed to create employee: %v\n", err)
			} else {
				fmt.Println("✓ New employee created successfully!")
				exportOnChange(s)
			}
		case 2:
			readEmployeeMenu(s)
		case 3:
			readAllEmployeesFlow(s)
		case 4:
			fmt.Println("Update Employee - Still working on it...")
		case 5:
			if err := deleteEmployeeFlow(s); err != nil {
				fmt.Printf("Failed to delete employee: %v\n", err)
			} else {
				fmt.Println("✓ Employee deleted successfully!")
				exportOnChange(s)
			}
		case 6:
			importEmployeesFlow(s, currentUserID)
		case 7:
			exportOnChange(s)
		case 8:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
