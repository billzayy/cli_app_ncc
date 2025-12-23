package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuRole(empSvc *services.EmployeeService, rolSvc *services.RoleService, currentUserID uuid.UUID) {
	options := []string{
		"Add Role Info (level, position, branch)",
		"Get Role Info (level, position, branch)",
		"Assign role for Employee",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Role Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			if err := addRoleInfoFlow(rolSvc, currentUserID); err != nil {
				fmt.Println(err)
			}
		case 2:
			if err := getRoleInfo(rolSvc); err != nil {
				fmt.Println(err)
			}
		case 3:
			assignEmployeeRoleFlow(empSvc, rolSvc, currentUserID)
		case 4:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
