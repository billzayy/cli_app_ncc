package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

const (
	optAssignMember    = 1
	optViewAssignments = 2
	optRemoveMember    = 3
)

func MenuMember(empSvc *services.EmployeeService, projSvc *services.ProjectService, currentUserID *uuid.UUID) {
	options := []string{
		"Assign Member to Project",
		"Get Assigned Member Project",
		"Delete member from project",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Project Member Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		var err error

		switch choice {
		case optAssignMember:
			err = assignMemberFlow(empSvc, projSvc, *currentUserID)
			if err != nil {
				fmt.Printf("Failed to assign member: %v\n", err)
			} else {
				fmt.Println("Member assigned successfully!")
			}

		case optViewAssignments:
			viewAllAssignmentsFlow(projSvc)

		case optRemoveMember:
			err = removeMemberFlow(empSvc, projSvc)
			if err != nil {
				fmt.Printf("Failed to remove member: %v\n", err)
			} else {
				fmt.Println("Member removed from project successfully!")
			}
		case 4:
			fmt.Println("Returning to main menu...")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
