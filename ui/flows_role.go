package ui

import (
	"bufio"
	cli "cli-app"
	"cli-app/services"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func addRoleInfoFlow(s *services.RoleService, createdBy uuid.UUID) error {
	options := []string{
		"Add level",
		"Add position",
		"Add branch",
		"Back to Main menu",
	}

	for {
		fmt.Println("\n=== Add Role Information ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			name, err := inputNameRole("levels")

			if err != nil {
				continue
			}

			if err := s.CreateRoleInfo("levels", name, createdBy); err != nil {
				return fmt.Errorf("failed to create level role: %w", err)
			}

			fmt.Println("Level role created successfully!")
			return nil
		case 2:
			name, err := inputNameRole("positions")

			if err != nil {
				continue
			}

			if err := s.CreateRoleInfo("positions", name, createdBy); err != nil {
				return fmt.Errorf("failed to create position role: %w", err)
			}

			fmt.Println("Position role created successfully!")
			return nil
		case 3:
			name, err := inputNameRole("branches")

			if err != nil {
				continue
			}
			if err := s.CreateRoleInfo("branches", name, createdBy); err != nil {
				return fmt.Errorf("failed to create branch role: %w", err)
			}

			fmt.Println("Branch role created successfully!")
			return nil
		case 4:
			return nil
		default:
			fmt.Println("Please choose other options : ")
			continue
		}
	}
}

func getRoleInfo(s *services.RoleService) error {
	options := []string{
		"Get levels",
		"Get positions",
		"Get branches",
		"Back to Main menu",
	}

	for {
		fmt.Println("\n=== Get Role Information ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-4): ", len(options))
		if choice == 0 {
			continue
		}

		if choice == 4 {
			return nil
		}

		var roleType string
		var label string
		switch choice {
		case 1:
			roleType = "levels"
			label = "Level"
		case 2:
			roleType = "positions"
			label = "Position"
		case 3:
			roleType = "branches"
			label = "Branche"
		default:
			fmt.Println("Invalid choice.")
			continue
		}

		results, err := s.GetRoleInfo(roleType)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", roleType, err)
			continue
		}

		if len(results) == 0 {
			fmt.Printf("No %ss found.\n", roleType)
		} else {
			fmt.Printf("\n--- %ss ---\n", label)
			for i, item := range results {
				fmt.Printf("%d. %s (ID: %s)\n", i+1, item.Name, item.Id)
			}
			fmt.Println()
		}
	}
}

func assignEmployeeRoleFlow(empSvc *services.EmployeeService, rolSvc *services.RoleService, createdBy uuid.UUID) error {
	employeeID := selectEmployee(empSvc)
	if employeeID == "" {
		return fmt.Errorf("no employee selected")
	}

	levelID := selectRole("levels", rolSvc)
	if levelID == "" {
		return fmt.Errorf("no level selected")
	}

	positionID := selectRole("positions", rolSvc)
	if positionID == "" {
		return fmt.Errorf("no positions selected")
	}

	branchID := selectRole("branches", rolSvc)
	if branchID == "" {
		return fmt.Errorf("no branch selected")
	}

	input := cli.AssignEmployeeRoleDTO{
		EmployeeId: uuid.MustParse(employeeID),
		LevelId:    uuid.MustParse(levelID),
		PositionId: uuid.MustParse(positionID),
		BranchId:   uuid.MustParse(branchID),
		CreatedBy:  createdBy,
	}
	err := rolSvc.AssignEmployeeRole(input)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func inputNameRole(typeName string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input new %s name (or empty to cancel): ", typeName)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	name := strings.TrimSpace(input)
	if name == "" {
		fmt.Println("Cancelled.")
	}
	return name, nil
}
