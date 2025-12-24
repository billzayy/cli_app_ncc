package database

var createLevel string = `INSERT INTO levels(name, created_by) VALUES($1,$2)`

var createPosition string = `INSERT INTO positions(name, created_by) VALUES($1,$2)`

var createBranch string = `INSERT INTO branches(name,created_by) VALUES($1,$2)`

var createEmployee string = `
	INSERT INTO employees(email, full_name, code, dob, phone, gender, created_by) 
	VALUES %s ON CONFLICT (email) DO NOTHING RETURNING id`

var getIdEmployee string = `SELECT id from employees WHERE email = $1`

var getEmployeeBy string = `SELECT email, full_name, code, gender, phone, dob FROM employees `

var deleteEmployee string = `DELETE FROM employees WHERE email = $1`

var createProject string = `INSERT INTO projects(name, notes, working_time, created_by) VALUES($1, $2, $3, $4)`

var getAllProjects string = `SELECT name, notes, working_time FROM projects`

var getProjectIdByName string = `SELECT id FROM projects WHERE name = $1`

var deleteProject string = `DELETE FROM projects WHERE id =  $1`

var assignProject string = `INSERT INTO employees_projects(project_id, employee_id, roles, created_by) VALUES($1,$2,$3,$4)`

var getAssignProject string = `
	SELECT ep.employee_id, ep.project_id, e.full_name, p.Name, ep.Roles, p.Working_Time FROM employees e 
	INNER JOIN employees_projects ep ON ep.employee_id = e.Id 
	INNER JOIN projects p ON ep.project_id = p.Id `

var deleteAssignProject string = `DELETE FROM employees_projects WHERE project_id = $1 AND employee_id = $2`

var createTasks string = `INSERT INTO tasks(name, created_by) VALUES($1, $2)`

var getAllTasks string = `SELECT * FROM tasks`

var deleteTasks string = `DELETE FROM tasks WHERE id = $1`

var addTaskToProject string = `INSERT INTO projects_tasks(project_id, task_id, employee_id, created_by) VALUES ($1, $2, $3, $4)`

var getTaskToProject string = `
	SELECT e.full_name as employee_name, p.Name as project_name, t.Name as task_name FROM projects_tasks pt
	INNER JOIN employees e ON e.Id = pt.employee_id
	INNER JOIN projects p ON pt.project_id = p.Id 
	INNER JOIN tasks t ON t.Id = pt.task_id`

var deleteTaskToProject string = `DELETE FROM projects_tasks WHERE employee_id = $1 AND project_id = $2`

var sumWorkingTime string = `
	SELECT e.full_name ,sum(p.working_time), ld.amount FROM employees_projects ep 
	INNER JOIN projects p ON p.Id = ep.project_id 
	INNER JOIN employees e ON e.Id = ep.employee_id 
	INNER JOIN employees_roles er ON e.Id = er.employee_id
	INNER JOIN levels l ON l.Id = er.level_id
	INNER JOIN level_defaults ld ON ld.level_id = l.Id
	WHERE EXTRACT(MONTH FROM p.created_time) = $1
	GROUP BY ld.amount, e.full_name;`

var getLevelInfo string = `SELECT id, name FROM levels`
var getPositionInfo string = `SELECT id, name FROM positions`
var getBranchesInfo string = `SELECT id, name FROM branches`

var addLevelsInfo string = `INSERT INTO levels (name, created_by) VALUES($1,$2)`
var addPositionsInfo string = `INSERT INTO positions (name, created_by) VALUES($1,$2)`
var addBranchesInfo string = `INSERT INTO branches (name, created_by) VALUES($1,$2)`

var assignEmployeRole string = `INSERT INTO employees_roles(employee_id, level_id, position_id, branch_id, created_by) VALUES($1, $2, $3, $4, $5)`
