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

func (es EmployeeService) CreateEmployee(in cli.AddEmployee) error {
	_, err := es.repo.Create(in)

	if err != nil {

		fmt.Printf("err existed: %e", err)
	}

	return nil
}

func (es EmployeeService) GetEmployeeId(email string) (uuid.UUID, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return uuid.Nil, fmt.Errorf("email cannot be empty")
	}
	id, err := es.repo.GetID(email)

	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("employee not found with email: %s", email)
	}

	return id, nil
}

func (es EmployeeService) GetEmployeeByEmail(email string) (cli.EmployeeDTO, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return cli.EmployeeDTO{}, fmt.Errorf("email cannot be empty")
	}

	result, err := es.repo.GetByEmail(email)

	if err != nil {
		return result, err
	}

	if result.Email == "" {
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
		return fmt.Errorf("invalid employee ID: nil")
	}
	return es.repo.Delete(email)
}

func (es *EmployeeService) ExportEmployee(filePath string) error {
	employeeList, err := es.GetAllEmployees()

	if err != nil {
		return fmt.Errorf("failed to fetch employees: %w", err)
	}

	format := [][]string{
		{"Id", "FullName", "Email", "Code", "Gender", "Phone", "Dob", "CreatedTime", "CreatedBy"},
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
	if importedBy == uuid.Nil {
		return fmt.Errorf("importedBy user ID cannot be nil")
	}

	rows, err := importCSV(filePath)

	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	dataRow := rows[1:]

	for i, v := range dataRow {
		if len(v) < 7 {
			return fmt.Errorf("row %d: insufficient columns", i+2)
		}

		email := strings.TrimSpace(v[2])

		if email == "" {

			fmt.Printf("err existed: %e", err)
			// return fmt.Errorf("row %d: email is empty", i+2)
		}

		check, err := es.GetEmployeeByEmail(v[2])

		if err == nil {
			fmt.Printf("err existed: %e", err)
		}

		if check.Email != v[2] {
			parsed, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", v[6])
			if err != nil {
				fmt.Printf("err existed: %e", err)
			}

			formatted := parsed.AddDate(0, 0, 0)
			in := cli.AddEmployee{
				Email:     v[2],
				FullName:  v[1],
				Code:      v[3],
				Gender:    v[4],
				Phone:     v[5],
				Dob:       formatted,
				CreatedBy: importedBy,
			}

			if err := es.CreateEmployee(in); err != nil {

				fmt.Printf("err existed: %e", err)
				// return fmt.Errorf("failed to import row %d (email: %s): %w", i+2, email, err)
			}
		}
	}

	return nil
}
