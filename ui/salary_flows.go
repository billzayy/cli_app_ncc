package ui

import (
	"cli-app/services"
	"fmt"
	"strconv"
)

func getValidMonth() (int, error) {
	for {
		month := readInt("Enter month (1-12): ", 12)
		if month >= 1 && month <= 12 {
			return month, nil
		}
		fmt.Println("Please enter a valid month between 1 and 12.")
	}
}

func calculateSalaryFlow(s *services.SalaryService) error {
	month, err := getValidMonth()
	if err != nil {
		return err
	}

	results, err := s.CalculateMonthSalary(month)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Printf("No salary data found for month %d.\n", month)
		return nil
	}

	fmt.Printf("\nSalary Calculation for Month %d\n\n", month)
	fmt.Printf("%-30s %-15s %-15s\n", "Employee", "Working Hours", "Amount (VND)")
	fmt.Println("-----------------------------------------------------------------")

	for _, r := range results {
		fmt.Printf("%-30s %-15d %-15d\n", r.Name, r.WorkingTime, r.Amount)
	}
	fmt.Println("-----------------------------------------------------------------")

	return nil
}

func exportSalaryFlow(s *services.SalaryService) error {
	month, err := getValidMonth()
	if err != nil {
		return err
	}

	results, err := s.CalculateMonthSalary(month)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Printf("No data to export for month %d.\n", month)
		return nil
	}

	data := make([][]string, len(results)+1)
	data[0] = []string{"Employee Name", "Working Hours", "Salary Amount (VND)"}

	for i, r := range results {
		data[i+1] = []string{
			r.Name,
			strconv.Itoa(r.WorkingTime),
			strconv.Itoa(r.Amount),
		}
	}

	if err := s.ExportSalary(data, month); err != nil {
		return err
	}

	fmt.Printf("Salary data for month %d exported to CSV successfully!\n", month)
	return nil
}
