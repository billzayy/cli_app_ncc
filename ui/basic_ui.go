package ui

import (
	"cli-app/database"
	"cli-app/services"
	"database/sql"
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
	var globalId uuid.UUID
	for {

		eId, err := a.promptEmployeeID()

		if err != nil {
			fmt.Println("failed to identify employee: %w", err)
			continue
		}

		globalId = eId
		break
	}

	for {
		choice := a.ShowMainMenu()

		if choice == 0 {
			continue
		}

		switch choice {
		case 1:
			a.ShowEmployeeMenu(globalId)
		case 2:
			a.ShowAdditionMenu()
		case 3:
			a.ShowRoleMenu()
		case 4:
			a.ShowProjectMenu(globalId)
		case 5:
			a.ShowSalaryMenu(globalId)
		default:
			return fmt.Errorf("invalid menu option: %q", choice)
		}
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

	email, err = validateEmail(email)

	if err != nil {
		return uuid.Nil, err
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
