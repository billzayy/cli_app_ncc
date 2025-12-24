package database

import (
	cli "cli-app"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type SalaryRepository interface {
	GetTimeAndAmount(month int, nameList []string) ([]cli.TimeAndAmount, error)
}

type salaryRepo struct {
	db *sql.DB
}

func InitSalaryRepo(db *sql.DB) SalaryRepository {
	return &salaryRepo{db: db}
}

func (sr *salaryRepo) GetTimeAndAmount(month int, nameList []string) ([]cli.TimeAndAmount, error) {
	if len(nameList) == 0 {
		return []cli.TimeAndAmount{}, nil
	}

	placeholders := make([]string, len(nameList))
	args := make([]any, len(nameList)+1)
	args[0] = month

	for i, name := range nameList {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = name
	}

	inClause := strings.Join(placeholders, ", ")

	query := fmt.Sprintf(sumWorkingTime, inClause)

	rows, err := sr.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var result []cli.TimeAndAmount
	for rows.Next() {
		var item cli.TimeAndAmount
		if err := rows.Scan(&item.Name, &item.WorkingTime, &item.Amount, &item.Month); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		item.BaseAmount = item.Amount
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	log.Printf("[SQL RESULT] %+v", result)
	return result, nil
}
