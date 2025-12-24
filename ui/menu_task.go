package ui

import (
	"cli-app/services"
	"fmt"

	"github.com/google/uuid"
)

func MenuTask(empSvc *services.EmployeeService, projSvc *services.ProjectService, currentUserID *uuid.UUID) {
	options := []string{
		"Create New Task",
		"Read Task (by ID)",
		"List All Tasks",
		"Update Task",
		"Delete Task",
		"Assign Task to Employee + Project",
		"View All Task Assignments",
		"Remove Task Assignment",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Task Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-%d): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			if err := createTaskFlow(empSvc, projSvc, *currentUserID); err != nil {
				fmt.Printf("Failed to create task: %v\n", err)
			} else {
				fmt.Println("✓ Task created successfully!")
			}
		case 2:
			fmt.Println("Read single task - Coming soon...")
		case 3:
			listAllTasksFlow(projSvc)
		case 4:
			fmt.Println("Update Task - Still working on it...")
		case 5:
			if err := deleteTaskFlow(projSvc); err != nil {
				fmt.Printf("Failed to delete task: %v\n", err)
			} else {
				fmt.Println("✓ Task deleted successfully!")
			}
		case 6:
			if err := assignTaskFlow(empSvc, projSvc, *currentUserID); err != nil {
				fmt.Printf("Failed to assign task: %v\n", err)
			} else {
				fmt.Println("✓ Task assigned successfully!")
			}
		case 7:
			viewTaskAssignmentsFlow(projSvc)
		case 8:
			if err := removeTaskAssignmentFlow(empSvc, projSvc); err != nil {
				fmt.Printf("Failed to remove assignment: %v\n", err)
			} else {
				fmt.Println("✓ Assignment removed successfully!")
			}
		case 9:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
