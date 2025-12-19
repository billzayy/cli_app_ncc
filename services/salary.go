package services

import (
	cli "cli-app"
	"cli-app/database"
	"fmt"
)

type SalaryService struct {
	repo database.SalaryRepository
}

func NewSalaryService(sr database.SalaryRepository) *SalaryService {
	return &SalaryService{
		repo: sr,
	}
}

func (ss *SalaryService) CalculateMonthSalary(month int) ([]cli.TimeAndAmount, error) {
	rawData, err := ss.CalculateMonthSalary(month)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch working time data: %w", err)
	}

	if len(rawData) == 0 {
		return []cli.TimeAndAmount{}, nil
	}

	result := make([]cli.TimeAndAmount, len(rawData))
	const standardHour = 172

	for i, v := range rawData {
		if v.WorkingTime < standardHour {
			calculated := v.Amount - ((standardHour - v.WorkingTime) * (v.Amount / standardHour))

			result[i] = cli.TimeAndAmount{
				Name:        v.Name,
				WorkingTime: v.WorkingTime,
				Amount:      calculated,
			}
			return result, nil
		}

		result[i] = rawData[i]
	}

	return result, nil
}

func (ss *SalaryService) ExportSalary(in [][]string, month int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("invalid month: %d", month)
	}

	if len(in) == 0 {
		return fmt.Errorf("no salary data to export")
	}

	filePath := "./csv/salaries.csv"

	format := [][]string{
		{"Name", "Month", "Salary"},
	}

	for _, v := range in {
		format = append(format, v)
	}

	return exportCSV(filePath, format)
}
