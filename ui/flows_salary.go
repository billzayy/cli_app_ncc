package ui

import (
	cli "cli-app"
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

func salaryMonthFlow(e *services.EmployeeService, s *services.SalaryService) {
	options := []string{
		"Calculate salary per employee",
		"Calculate salary for all employees",
		"Back to Main menu",
	}

	for {
		fmt.Println("\n=== Salary Per Month ===")
		for i, opt := range options {
			fmt.Printf("%d. %s\n", i+1, opt)
		}

		choice := readInt("Choose an option (1-3): ", len(options))
		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			err := calculateSalaryPerMonthFlow(e, s, "one")
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := calculateSalaryPerMonthFlow(e, s, "all")
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			return
		default:
			fmt.Println("Please choose option correctly !")
			continue
		}
	}
}

func calculateSalaryPerMonthFlow(e *services.EmployeeService, s *services.SalaryService, many string) error {
	var results []cli.TimeAndAmount

	month, err := getValidMonth()
	if err != nil {
		return err
	}

	switch many {
	case "one":
		email := selectEmailEmployee(e)

		person, err := s.CalculateMonthSalary(month, []string{email})
		if err != nil {
			return err
		}

		results = person
	case "all":
		var mailList []string
		list, err := e.GetAllEmployees()

		if err != nil {
			return fmt.Errorf("")
		}

		for _, v := range list {
			mailList = append(mailList, v.Email)
		}

		person, err := s.CalculateMonthSalary(month, mailList)
		if err != nil {
			return err
		}

		results = person
	case "":
		return fmt.Errorf("missing many type")
	default:
		return fmt.Errorf("Only accept one or all type")
	}

	if len(results) == 0 {
		fmt.Printf("No salary data found for month %d.\n", month)
		return nil
	}

	fmt.Printf("\nSalary Calculation for Month %d\n\n", month)
	fmt.Printf("%-30s %-15s %-15s %-15s %-15s\n", "Employee", "Working Hours", "Base Amount", "Amount (VND)", "Month")
	fmt.Println("-------------------------------------------------------------------------------------")

	for _, r := range results {
		fmt.Printf("%-30s %-15d %-15d %-15d %-12d\n", r.Name, r.WorkingTime, r.BaseAmount, r.Amount, r.Month)
	}
	fmt.Println("-------------------------------------------------------------------------------------")

	return nil
}

func exportSalaryFlow(s *services.SalaryService) error {
	month, err := getValidMonth()
	if err != nil {
		return err
	}

	results, err := s.CalculateMonthSalary(month, []string{})
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
