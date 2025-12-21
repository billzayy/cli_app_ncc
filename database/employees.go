package database

import (
	cli "cli-app"
	"database/sql"
	"fmt"
	"log"
	"strings"

	// "strings"

	"github.com/google/uuid"
)

type EmployeeRepository interface {
	Create(input []cli.AddEmployee) (bool, error)
	GetID(email string) (uuid.UUID, error)
	GetByEmail(email string) (cli.EmployeeDTO, error)
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

	log.Printf("Successfully inserted %d row(s)", affected)
	return true, nil
}

func (er *employeeRepo) GetID(email string) (uuid.UUID, error) {
	var result uuid.UUID
	rows, err := er.db.Query(getIdEmployee, email)

	if err != nil {
		return result, nil
	}

	var temp string
	for rows.Next() {
		err := rows.Scan(&temp)

		if err != nil {
			return result, err
		}

		result = uuid.MustParse(temp)
	}

	return result, nil
}

func (er *employeeRepo) GetByEmail(email string) (cli.EmployeeDTO, error) {
	rows, err := er.db.Query(getEmployeeByEmail, email)

	if err != nil {
		return cli.EmployeeDTO{}, err
	}

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

	return result, nil
}

func (er *employeeRepo) GetAll() ([]cli.EmployeeDTO, error) {
	rows, err := er.db.Query(getAllEmployees)

	if err != nil {
		return []cli.EmployeeDTO{}, err
	}

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

	return nil
}

func Debug() {

	// // In ra câu lệnh SQL gốc
	// log.Printf("\nExecuting SQL: %s", createEmployee)
	//
	// // In ra các tham số truyền vào (theo thứ tự)
	// log.Printf("Parameters: Email=%s, FullName=%s, Code=%s, Dob=%s, Phone=%s, Gender=%s, CreatedBy=%s",
	// 	input.Email, input.FullName, input.Code, input.Dob, input.Phone, input.Gender, input.CreatedBy)
	//
	// // Tạo phiên bản SQL "bound" (thay ? bằng giá trị thực, chỉ để debug - chú ý escape quote nếu cần)
	// boundSQL := createEmployee
	//
	// params := []interface{}{
	// 	input.Email, input.FullName, input.Code, input.Dob, input.Phone, input.Gender, input.CreatedBy,
	// }
	//
	// for _, p := range params {
	// 	if s, ok := p.(string); ok {
	// 		// Escape single quote đơn giản cho debug
	// 		escaped := strings.ReplaceAll(s, "'", "''")
	// 		boundSQL = strings.Replace(boundSQL, "?", fmt.Sprintf("'%s'", escaped), 1)
	// 	} else {
	// 		boundSQL = strings.Replace(boundSQL, "?", fmt.Sprintf("%v", p), 1)
	// 	}
	// }
	//
	// log.Printf("Bound SQL (for debugging): %s", boundSQL)
}
