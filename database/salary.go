package database

import (
	cli "cli-app"
	"database/sql"
)

type SalaryRepository interface {
	GetTimeAndAmount(month int) ([]cli.TimeAndAmount, error)
}

type salaryRepo struct {
	db *sql.DB
}

func InitSalaryRepo(db *sql.DB) SalaryRepository {
	return &salaryRepo{db: db}
}

func (sr *salaryRepo) GetTimeAndAmount(month int) ([]cli.TimeAndAmount, error) {
	rows, err := sr.db.Query(sumWorkingTime, month)

	if err != nil {
		return []cli.TimeAndAmount{}, err
	}

	var result []cli.TimeAndAmount
	var temp cli.TimeAndAmount

	for rows.Next() {
		err := rows.Scan(&temp.Name, &temp.WorkingTime, &temp.Amount)

		if err != nil {
			return result, err
		}

		result = append(result, temp)
	}

	return result, nil
}
