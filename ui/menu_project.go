package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuProject(empSvc *services.EmployeeService, projSvc *services.ProjectService, currentUserID *uuid.UUID) {
	options := []string{
		"Projects",
		"Tasks",
		"Assign Member Project",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Project Menu ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			projectsFlow(projSvc, *currentUserID)
		case 2:
			MenuTask(empSvc, projSvc, currentUserID)
		case 3:
			MenuMember(empSvc, projSvc, currentUserID)
		case 4:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
