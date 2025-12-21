package ui

import (
	"cli-app/database"
	"cli-app/services"
	"database/sql"
	"flag"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type App struct {
	db      *sql.DB
	empSvc  *services.EmployeeService
	projSvc *services.ProjectService
	salSvc  *services.SalaryService
}

func NewApp(db *sql.DB) *App {
	return &App{
		db:      db,
		empSvc:  services.NewEmployeeService(database.InitEmployeeRepo(db)),
		projSvc: services.NewProjectService(database.InitProjectRepo(db)),
		salSvc:  services.NewSalaryService(database.InitSalaryRepo(db)),
	}
}

func (a *App) Run() error {
	id, err := a.promptEmployeeID()

	if err != nil {
		return fmt.Errorf("failed to identify employee: %w", err)
	}

	for {
		choice := a.ShowMainMenu()

		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			a.ShowEmployeeMenu(id)
		case 2:
			a.ShowAdditionMenu()
		case 3:
			a.ShowRoleMenu()
		case 4:
			a.ShowProjectMenu(id)
		case 5:
			a.ShowSalaryMenu(id)
		default:
			a.printUsage()
			return fmt.Errorf("invalid menu option: %q", choice)
		}
	}

}

func (a *App) RunWFlag() error {
	menu := flag.String("menu", "", "Menu option: all, employee, addition, role, project, salary (or numbers 1-5)")
	flag.Parse()

	input := *menu
	choice := a.parseChoice(input)

	switch {
	case input == "" || input == "all" || choice == 0:
		a.ShowMainMenu()
		return nil
	case input == "employee" || choice == 1:
		id, err := a.promptEmployeeID()
		if err != nil {
			return fmt.Errorf("failed to identify employee: %w", err)
		}
		a.ShowEmployeeMenu(id)
		return nil
	case input == "addition" || choice == 2:
		a.ShowAdditionMenu()
		return nil
	case input == "role" || choice == 3:
		a.ShowRoleMenu()
		return nil
	case input == "project" || choice == 4:
		id, err := a.promptEmployeeID()
		if err != nil {
			return fmt.Errorf("failed to identify employee: %w", err)
		}
		a.ShowProjectMenu(id)
		return nil
	case input == "salary" || choice == 5:
		id, err := a.promptEmployeeID()
		if err != nil {
			return fmt.Errorf("failed to identify employee: %w", err)
		}
		a.ShowSalaryMenu(id)
		return nil
	default:
		a.printUsage()
		return fmt.Errorf("invalid menu option: %q", input)
	}
}

func (a *App) parseChoice(input string) int {
	if n, err := strconv.Atoi(input); err == nil {
		return n
	}
	return 0
}

func (a *App) ShowMainMenu() int {
	options := []string{
		"Employee",
		"Additional Employee",
		"Roles",
		"Projects",
		"Salary",
	}

	fmt.Println("\n=== Main Menu ===")
	for i, opt := range options {
		fmt.Printf(" %d. %s\n", i+1, opt)
	}

	choose := readInt("Please select an option: ", len(options))

	return choose
}

// TODO: check email format
func (a *App) promptEmployeeID() (uuid.UUID, error) {
	var email string
	fmt.Print("\nWho are you? (Input email): ")
	_, err := fmt.Scan(&email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to read email: %w", err)
	}
	if email == "" {
		return uuid.Nil, fmt.Errorf("email cannot be empty")
	}

	id, err := a.empSvc.GetEmployeeId(email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("employee not found for email %q: %w", email, err)
	}

	fmt.Printf("Welcome %s\n\n", email)

	return id, nil
}

func (a *App) ShowEmployeeMenu(employeeID uuid.UUID) {
	MenuEmployee(a.empSvc, employeeID)
}

func (a *App) ShowAdditionMenu() {
	fmt.Println("Additional Employee menu - Still working on it...")
}

func (a *App) ShowRoleMenu() {
	fmt.Println("Roles menu - Still working on it...")
}

func (a *App) ShowProjectMenu(employeeID uuid.UUID) {
	MenuProject(a.empSvc, a.projSvc, &employeeID)
}

func (a *App) ShowSalaryMenu(employeeID uuid.UUID) {
	MenuSalary(a.empSvc, a.salSvc, &employeeID)
}

func (a *App) printUsage() {
	fmt.Println("Usage: go run ./cmd/main.go --menu=<option>")
	fmt.Println("\nAvailable options:")
	fmt.Println(" (empty) or 'all' - Show this menu")
	fmt.Println(" employee or 1 - Employee menu")
	fmt.Println(" addition or 2 - Additional Employee (WIP)")
	fmt.Println(" role or 3 - Roles menu (WIP)")
	fmt.Println(" project or 4 - Projects menu")
	fmt.Println(" salary or 5 - Salary menu")
	fmt.Println("\nExample: go run . --menu=project")
}
