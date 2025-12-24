package ui

import (
	cli "cli-app"
	"cli-app/components"
	"cli-app/services"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func createEmployeeFlow(s *services.EmployeeService, rolSvc *services.RoleService, createdBy uuid.UUID) error {
	fields := []string{
		"Email", "FullName", "Code", "Gender", "Phone", "Dob (DD-MM-YYYY)"}

	input := components.TextInput[cli.InputCreateEmployee](fields)

	addEmp, err := convertToAddEmployee(input, rolSvc, createdBy)

	if err != nil {
		return fmt.Errorf("invalid date of birth: %w", err)
	}

	levelId := selectRole("levels", rolSvc)
	positionId := selectRole("positions", rolSvc)
	branchId := selectRole("branches", rolSvc)

	addEmp.LevelId = levelId
	addEmp.PositionId = positionId
	addEmp.BranchId = branchId

	var listInput []cli.AddEmployee
	listInput = append(listInput, addEmp)

	id, err := s.CreateEmployee(listInput)

	if err != nil {
		return err
	}

	assignRole := cli.AssignEmployeeRoleDTO{
		EmployeeId: id,
		LevelId:    uuid.MustParse(levelId),
		PositionId: uuid.MustParse(positionId),
		BranchId:   uuid.MustParse(branchId),
		CreatedBy:  createdBy,
	}

	err = rolSvc.AssignEmployeeRole(assignRole)

	if err != nil {
		return err
	}

	return nil
}

func convertToAddEmployee(in cli.InputCreateEmployee, rolSvc *services.RoleService, createdBy uuid.UUID) (cli.AddEmployee, error) {
	const layout = "02-01-2006"

	dob, err := time.Parse(layout, strings.TrimSpace(in.Dob))

	if err != nil {
		return cli.AddEmployee{}, err
	}

	return cli.AddEmployee{
		Email:     strings.TrimSpace(in.Email),
		FullName:  strings.TrimSpace(in.FullName),
		Code:      strings.TrimSpace(in.Code),
		Gender:    strings.TrimSpace(in.Gender),
		Phone:     strings.TrimSpace(in.Phone),
		Dob:       dob,
		CreatedBy: createdBy,
	}, nil
}

func readEmployeeByEmailFlow(s *services.EmployeeService) {
	var email string

	fmt.Print("Enter employee email: ")
	fmt.Scan(&email)

	email = strings.TrimSpace(email)

	if email == "" {
		fmt.Println("Email cannot be empty.")
		return
	}

	emp, err := s.GetEmployeeBy(email, "email")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nEmployee Details:\n%+v\n", emp)
}

func readEmployeeByCodeFlow(s *services.EmployeeService) {
	var code string

	fmt.Print("Enter employee code: ")
	fmt.Scan(&code)

	code = strings.TrimSpace(code)

	if code == "" {
		fmt.Println("Code cannot be empty.")
		return
	}

	emp, err := s.GetEmployeeBy(code, "code")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nEmployee Details:\n%+v\n", emp)
}

func readEmployeeMenu(s *services.EmployeeService) {
	options := []string{
		"Read Employee (by Email)",
		"Read Employee (by Code)",
		"Back to Main menu",
	}

	for {
		fmt.Println("\n=== Read Employees ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-2): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			readEmployeeByEmailFlow(s)
		case 2:
			readEmployeeByCodeFlow(s)
		case 3:
			return
		default:
			fmt.Println("Please choose option correctly !")
			continue
		}
	}
}

func readAllEmployeesFlow(s *services.EmployeeService) {
	list, err := s.GetAllEmployees()

	if err != nil {
		fmt.Printf("Failed to fetch employees: %v\n", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Printf("\nTotal Employees: %d\n\n", len(list))

	components.ViewPaginationList(
		list,
		func(e cli.EmployeeDTO) string {
			return strings.Join([]string{
				"Email: " + e.Email,
				"Name: " + e.FullName,
				"Code: " + e.Code,
				"Phone: " + e.Phone,
				"Gender: " + string(e.Gender),
				"DOB: " + e.Dob.Format("2006-01-02"),
			}, " | ")
		},
		components.WithPerPage(8),
		components.WithTitle("All Employees"),
		components.WithHelp(true),
	)
}

func deleteEmployeeFlow(s *services.EmployeeService) error {
	var emailStr string

	fmt.Print("Enter employee's Email to delete: ")
	fmt.Scan(&emailStr)

	emailStr = strings.TrimSpace(emailStr)

	if emailStr == "" {
		return fmt.Errorf("Email cannot be empty")
	}

	return s.DeleteEmployee(emailStr)
}

func importEmployeesFlow(s *services.EmployeeService, importedBy uuid.UUID) {
	fmt.Println("Select a CSV file to import...")

	fileName := components.FilePicker()

	if fileName == "" {
		fmt.Println("No file selected. Import cancelled.")
		return
	}

	path := "./csv/" + fileName

	if err := s.ImportEmployee(path, importedBy); err != nil {
		fmt.Printf("Import failed: %v\n", err)
		return
	}

	fmt.Println("✓ Employees imported successfully!")

	exportOnChange(s)
}
