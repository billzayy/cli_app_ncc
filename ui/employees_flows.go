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

func createEmployeeFlow(s *services.EmployeeService, createdBy uuid.UUID) error {
	fields := []string{
		"Email", "FullName", "Code", "Gender", "Phone", "Dob (DD-MM-YYYY)",
		"Origin", "Residence", "CurrentLocation",
	}

	input := components.TextInput[cli.InputCreateEmployee](fields)

	addEmp, err := convertToAddEmployee(input, createdBy)

	if err != nil {
		return fmt.Errorf("invalid date of birth: %w", err)
	}

	var listInput []cli.AddEmployee
	listInput = append(listInput, addEmp)

	return s.CreateEmployee(listInput)
}

func convertToAddEmployee(in cli.InputCreateEmployee, createdBy uuid.UUID) (cli.AddEmployee, error) {
	const layout = "02-01-2006"

	dob, err := time.Parse(layout, strings.TrimSpace(in.Dob))

	if err != nil {
		return cli.AddEmployee{}, err
	}

	return cli.AddEmployee{
		Email:           strings.TrimSpace(in.Email),
		FullName:        strings.TrimSpace(in.FullName),
		Code:            strings.TrimSpace(in.Code),
		Gender:          strings.TrimSpace(in.Gender),
		Phone:           strings.TrimSpace(in.Phone),
		Dob:             dob,
		Origin:          strings.TrimSpace(in.Origin),
		Residence:       strings.TrimSpace(in.Residence),
		CurrentLocation: strings.TrimSpace(in.CurrentLocation),
		CreatedBy:       createdBy,
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

	emp, err := s.GetEmployeeByEmail(email)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nEmployee Details:\n%+v\n", emp)
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

	for _, emp := range list {
		fmt.Printf("%+v\n\n", emp)
	}
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
