package ui

import (
	"fmt"
	"strconv"

	cli "cli-app"
	"cli-app/components"
	"cli-app/services"

	"github.com/google/uuid"
)

func projectsFlow(projSvc *services.ProjectService, createdBy uuid.UUID) {
	options := []string{
		"Create Project",
		"Get All Projects",
		"Delete Project",
		"Count working time per month",
		"Back to Main Menu",
	}

	for {
		fmt.Println("\n=== Project Management ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		switch choice {
		case 1:
			if err := handleCreateProject(projSvc, createdBy); err != nil {
				fmt.Printf("Failed to create project: %v\n", err)
			} else {
				fmt.Println("Project created successfully!")
			}
		case 2:
			handleListProjects(projSvc)
		case 3:
			handleDeleteProject(projSvc)
		case 4:
		case 5:
			fmt.Println("Returning to main menu...")
			return
		default:
			fmt.Println("Invalid option selected.")
		}
	}
}

func handleCreateProject(projSvc *services.ProjectService, createdBy uuid.UUID) error {
	titles := []string{"Name", "Notes", "Working Time (hours)"}
	input := components.TextInput[cli.InputAddProject](titles)

	project, err := convertToAddProject(input, createdBy)
	if err != nil {
		return fmt.Errorf("invalid input: %w", err)
	}

	if err := validateAddProject(project); err != nil {
		return err
	}

	return projSvc.Create(project)
}

func convertToAddProject(in cli.InputAddProject, createdBy uuid.UUID) (cli.AddProject, error) {
	workingTime, err := strconv.Atoi(in.WorkingTime)
	if err != nil {
		return cli.AddProject{}, fmt.Errorf("working time must be a valid number: %w", err)
	}

	if workingTime < 0 {
		return cli.AddProject{}, fmt.Errorf("working time cannot be negative")
	}

	return cli.AddProject{
		Name:        in.Name,
		Notes:       in.Notes,
		WorkingTime: workingTime,
		CreatedBy:   createdBy,
	}, nil
}

func validateAddProject(p cli.AddProject) error {
	if p.Name == "" {
		return fmt.Errorf("project name is required")
	}
	return nil
}

func handleListProjects(projSvc *services.ProjectService) {
	projects, err := projSvc.GetAll()
	if err != nil {
		fmt.Printf("Error fetching projects: %v\n", err)
		return
	}

	if len(projects) == 0 {
		fmt.Println("No projects found.")
		return
	}

	fmt.Println("\n=== Your Projects ===")
	for _, p := range projects {
		fmt.Printf("- %s (Working Time: %d hours)\n  Notes: %s\n \n\n",
			p.Name, p.WorkingTime, p.Notes)
	}
}

// handleDeleteProject lets user delete a project by ID or name
func handleDeleteProject(projSvc *services.ProjectService) {
	// projects, err := projSvc.GetAll()
	// if err != nil {
	// 	fmt.Printf("Error loading projects: %v\n", err)
	// 	return
	// }
	//
	// if len(projects) == 0 {
	// 	fmt.Println("No projects to delete.")
	// 	return
	// }
	//
	// fmt.Println("\n=== Delete Project ===")
	// for i, p := range projects {
	// 	fmt.Printf("%d. %s\n", i+1, p.Name)
	// }
	//
	// choice := readInt("Select project to delete (1-"+fmt.Sprint(len(projects))+"), or 0 to cancel: ", 0, len(projects))
	// if choice == 0 {
	// 	fmt.Println("Delete cancelled.")
	// 	return
	// }
	//
	// projectToDelete := projects[choice-1]
	//
	// confirm := components.ReadString(fmt.Sprintf("Type 'DELETE' to confirm deletion of '%s': ", projectToDelete.Name))
	// if confirm != "DELETE" {
	// 	fmt.Println("Deletion cancelled.")
	// 	return
	// }
	//
	// if err := projSvc.Delete(projectToDelete.Id); err != nil {
	// 	fmt.Printf("Failed to delete project: %v\n", err)
	// } else {
	// 	fmt.Println("Project deleted successfully.")
	// }
}
