package database

import (
	cli "cli-app"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	Create(input []cli.AddEmployee) (bool, error)
	GetID(email string) (uuid.UUID, error)
	GetByInput(input string, filter string) (cli.EmployeeDTO, error)
	GetAll() ([]cli.EmployeeDTO, error)
	Delete(email string) error
}

type employeeRepo struct {
	db *sql.DB
}

func InitEmployeeRepo(db *sql.DB) EmployeeRepository {
	return &employeeRepo{
		db: db,
	}
}

func (er *employeeRepo) Create(input []cli.AddEmployee) (bool, error) {
	values := []string{}
	args := []interface{}{}
	argPos := 1

	for _, e := range input {
		values = append(values,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				argPos, argPos+1, argPos+2,
				argPos+3, argPos+4, argPos+5, argPos+6,
			),
		)

		args = append(args,
			e.Email,
			e.FullName,
			e.Code,
			e.Dob,
			e.Phone,
			e.Gender,
			e.CreatedBy,
		)

		argPos += 7
	}

	query := fmt.Sprintf(createEmployee, strings.Join(values, ","))

	query += "ON CONFLICT (email) DO NOTHING" // Add confict query will do nothing when conflict email import

	// Thực thi query
	res, err := er.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return false, err
	}

	if affected == 0 {
		log.Println("No rows affected")
		return false, nil // hoặc error tùy yêu cầu
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)
	return true, nil
}

func (er *employeeRepo) GetID(email string) (uuid.UUID, error) {
	var result uuid.UUID
	rows, err := er.db.Query(getIdEmployee, email)

	if err != nil {
		return result, nil
	}

	defer rows.Close()

	var temp string
	for rows.Next() {
		err := rows.Scan(&temp)

		if err != nil {
			return result, err
		}

		result = uuid.MustParse(temp)
	}
	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (er *employeeRepo) GetByInput(input string, filter string) (cli.EmployeeDTO, error) {
	query := getEmployeeBy

	switch filter {
	case "email":
		query += "WHERE email = $1"
	case "code":
		query += "WHERE code = $1"
	default:
		return cli.EmployeeDTO{}, fmt.Errorf("Filter is empty")
	}

	rows, err := er.db.Query(query, input)

	if err != nil {
		return cli.EmployeeDTO{}, err
	}

	defer rows.Close()

	var result cli.EmployeeDTO
	var phone sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&result.Email, &result.FullName, &result.Code,
			&result.Gender, &phone, &result.Dob)

		if err != nil {
			return cli.EmployeeDTO{}, err
		}

		if phone.Valid {
			result.Phone = phone.String
		}
	}

	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (er *employeeRepo) GetAll() ([]cli.EmployeeDTO, error) {
	rows, err := er.db.Query(getEmployeeBy)

	if err != nil {
		return []cli.EmployeeDTO{}, err
	}

	defer rows.Close()

	var result []cli.EmployeeDTO
	var temp cli.EmployeeDTO
	var phone sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&temp.Email, &temp.FullName, &temp.Code,
			&temp.Gender, &phone, &temp.Dob)

		if err != nil {
			return []cli.EmployeeDTO{}, err
		}

		if phone.Valid {
			temp.Phone = phone.String
		}

		result = append(result, temp)
	}

	log.Printf("[SQL RESULT] %+v", result)

	return result, nil
}

func (er *employeeRepo) Delete(email string) error {
	rows, err := er.db.Exec(deleteEmployee, email)

	if err != nil {
		return err
	}

	affected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found employee")
	}

	log.Printf("[SQL RESULT] rows_affected=%d", affected)

	return nil
}
