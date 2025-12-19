package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuProject(empSvc *services.EmployeeService, projSvc *services.ProjectService, currentUserID *uuid.UUID) {
	options := []string{
		"Assign Member to Project",
		"View All Assignments",
		"Remove Member from Project",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Project Member Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-%d): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			if err := assignMemberFlow(empSvc, projSvc, *currentUserID); err != nil {
				fmt.Printf("Failed to assign member: %v\n", err)
			} else {
				fmt.Println("✓ Member assigned successfully!")
			}
		case 2:
			viewAllAssignmentsFlow(projSvc)
		case 3:
			if err := removeMemberFlow(empSvc, projSvc); err != nil {
				fmt.Printf("Failed to remove member: %v\n", err)
			} else {
				fmt.Println("✓ Member removed successfully!")
			}
		case 4:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
