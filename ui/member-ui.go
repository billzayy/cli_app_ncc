package ui

import (
	cli "cli-app"
	"cli-app/components"
	"cli-app/services"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/google/uuid"
)

type getList struct {
	Name    string
	Project string
	Role    string
}

func memberUI(s *services.EmployeeService, p *services.ProjectService, id *uuid.UUID) {
	list := []string{"Assign", "Get", "Delete"}

	for i, v := range list {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	var choose int
	fmt.Print("Choose your feature : ")
	fmt.Scan(&choose)

	switch choose {
	case 1: // Assign
		eId, pId := showTableEP(s, p)

		var roles string
		var result cli.EmployeeProject

		fmt.Print("Input your roles: ")
		fmt.Scan(&roles)

		result.ProjectId = uuid.MustParse(pId)
		result.EmployeeId = uuid.MustParse(eId)
		result.Roles = roles
		result.CreatedBy = *id

		err := p.AssignMember(result)

		if err != nil {
			panic(err)
		}

		fmt.Println("Assign Member successful")
	case 2: // Get
		list, err := p.GetAssignProject()

		if err != nil {
			panic(err)
		}

		output := make([]getList, len(list))

		for i, v := range list {
			output[i] = getList{
				Name:    v.EmployeeName,
				Project: v.Project,
				Role:    v.Role,
			}
		}

		fmt.Println(output)
	case 3: // Delete
		var eId, pId string
		fmt.Print("Input your employee id and project id want to delete (eId - pId): ")
		fmt.Scanf("%s-%s", eId, pId)

		err := p.DeleteAssignProject(uuid.MustParse(eId), uuid.MustParse(pId))

		if err != nil {
			panic(err)
		}

		fmt.Println("Delete assignment of project successful")
	}

}

func showTableE(s *services.EmployeeService) string {
	var employeeId string
	employeeTableColumn := []table.Column{
		{Title: "Id", Width: 36},
		{Title: "Email", Width: 20},
		{Title: "FullName", Width: 12},
		{Title: "Code", Width: 10},
		{Title: "Gender", Width: 12},
		{Title: "Phone", Width: 11},
		{Title: "Dob", Width: 10},
	}

	outputE, err := s.GetAllEmployees()

	if err != nil {
		panic(err)
	}

	if len(outputE) > 0 {
		employeeRows := employeesToTableRows(outputE)

		employeeId = components.Table(employeeTableColumn, employeeRows)
	}

	return employeeId
}

func showTableP(p *services.ProjectService) string {
	var projectId string

	projectTableColumn := []table.Column{
		{Title: "Id", Width: 36},
		{Title: "Name", Width: 15},
		{Title: "Notes", Width: 15},
		{Title: "Working Time", Width: 12},
	}

	outputP, err := p.GetAll()

	if err != nil {
		panic(err)
	}

	if len(outputP) > 0 {
		projectRows := projectsToTableRows(outputP)
		projectId = components.Table(projectTableColumn, projectRows)
	}

	return projectId
}

func showTableEP(s *services.EmployeeService, p *services.ProjectService) (string, string) {
	return showTableE(s), showTableP(p)
}
