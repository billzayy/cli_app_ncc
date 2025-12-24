package services

import (
	cli "cli-app"
	"cli-app/database"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type EmployeeService struct {
	repo database.EmployeeRepository
}

func NewEmployeeService(er database.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: er,
	}
}

func (es EmployeeService) CreateEmployee(in []cli.AddEmployee) (uuid.UUID, error) {
	id, _, err := es.repo.Create(in)

	if err != nil {
		fmt.Printf("err existed: %e", err)
		return id, err
	}

	return id, nil
}

func (es EmployeeService) GetEmployeeId(email string) (uuid.UUID, error) {
	id, err := es.repo.GetID(email)

	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("employee not found with email: %s", email)
	}

	return id, nil
}

func (es EmployeeService) GetEmployeeBy(input string, filter string) (cli.EmployeeDTO, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return cli.EmployeeDTO{}, fmt.Errorf("%s cannot be empty", filter)
	}

	result, err := es.repo.GetByInput(input, filter)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (es EmployeeService) GetAllEmployees() ([]cli.EmployeeDTO, error) {
	result, err := es.repo.GetAll()

	if err != nil {
		return []cli.EmployeeDTO{}, err
	}

	return result, nil
}

func (es *EmployeeService) DeleteEmployee(email string) error {
	if len(email) == 0 {
		return fmt.Errorf("invalid employee's email: nil")
	}
	return es.repo.Delete(email)
}

func (es *EmployeeService) ExportEmployee(filePath string) error {
	employeeList, err := es.GetAllEmployees()

	if err != nil {
		return fmt.Errorf("failed to fetch employees: %w", err)
	}

	format := [][]string{
		{"FullName", "Email", "Code", "Gender", "Phone", "Dob"},
	}

	for _, emp := range employeeList {
		row := []string{
			emp.FullName,
			emp.Email,
			emp.Code,
			string(emp.Gender),
			emp.Phone,
			emp.Dob.Format("2006-01-02"),
		}
		format = append(format, row)
	}

	return exportCSV(filePath, format)
}

func (es *EmployeeService) ImportEmployee(filePath string, importedBy uuid.UUID) error {
	var listInput []cli.AddEmployee

	if importedBy == uuid.Nil {
		return fmt.Errorf("importedBy user ID cannot be nil")
	}

	rows, err := importCSV(filePath)

	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	dataRow := rows[1:]

	for i, v := range dataRow {
		if len(v) < 5 {
			return fmt.Errorf("row %d: insufficient columns", i+2)
		}

		email := strings.TrimSpace(v[1])

		if email == "" {
			fmt.Printf("err existed: %e", err)
		}

		parsed, err := time.Parse("2006-01-02", v[5])
		if err != nil {
			fmt.Printf("err existed: %e", err)
		}

		formatted := parsed.AddDate(0, 0, 0)
		in := cli.AddEmployee{
			Email:     v[1],
			FullName:  v[0],
			Code:      v[2],
			Gender:    v[3],
			Phone:     v[4],
			Dob:       formatted,
			CreatedBy: importedBy,
		}

		listInput = append(listInput, in)
	}

	_, _, err = es.repo.Create(listInput)

	if err != nil {
		fmt.Printf("err existed: %e", err)
	}

	return nil
}
