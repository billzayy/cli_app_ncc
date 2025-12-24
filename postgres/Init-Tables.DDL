-- drop table level_defaults;
-- drop table banks;
-- drop table identifiers;
-- drop table address;
-- drop table employees_roles;
-- drop table employees_benefit;
-- drop table benefits;
-- drop table levels;
-- drop table branches;
-- drop table positions;
-- drop table projects_tasks;
-- drop table tasks;
-- drop table employees_projects;
-- drop table projects;
-- drop table employee_departments;
-- drop table departments;
-- drop table employees;

-- CREATE TYPE genderType AS ENUM ('male', 'female', 'do not tell');

CREATE TABLE employees(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email VARCHAR(100) NOT NULL UNIQUE,
	full_name VARCHAR(100) NOT NULL ,
	code VARCHAR(50) NOT NULL UNIQUE,
	gender genderType NOT NULL DEFAULT 'do not tell',
	phone VARCHAR(11) NULL,
	dob DATE NOT NULL,	
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NULL,
	FOREIGN KEY (created_by) REFERENCES employees(id)
);

CREATE TABLE banks(
	account_number VARCHAR(12),
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	employee_id UUID,
	PRIMARY KEY (employee_id),
	FOREIGN KEY(created_by) REFERENCES employees(id),
	FOREIGN KEY(employee_id) REFERENCES employees(id)
);

CREATE TABLE identifiers(
	card_number VARCHAR(11) NOT NULL,
	date_of_issue DATE NOT NULL,
	issued_by VARCHAR(100) NOT NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	employee_id UUID,
	PRIMARY KEY(employee_id),
	FOREIGN KEY(created_by) REFERENCES employees(id),
	FOREIGN KEY(employee_id) REFERENCES employees(id)
);

CREATE TABLE address(
	origin VARCHAR(100) NOT NULL,
	residence VARCHAR(100) NULL,
	current_location VARCHAR(100) NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	employee_id UUID,
	PRIMARY KEY(employee_id),
	FOREIGN KEY(created_by) REFERENCES employees(id),
	FOREIGN KEY (employee_id) REFERENCES employees(id)
);

CREATE TABLE benefits(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE levels(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL UNIQUE,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE positions(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL UNIQUE,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE branches(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL UNIQUE,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE employees_roles(
	employee_id UUID,
	level_id UUID,
	position_id UUID,
	branch_id UUID,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
  	created_by   UUID NOT NULL,
	PRIMARY KEY(employee_id, level_id, position_id),
	FOREIGN KEY(created_by) REFERENCES employees(id),
	FOREIGN KEY(employee_id) REFERENCES employees(id),
	FOREIGN KEY(level_id) REFERENCES levels(id),
	FOREIGN KEY(position_id) REFERENCES positions(id),
	FOREIGN KEY(branch_id) REFERENCES branches(id)
);

CREATE TABLE employees_benefit(
	employee_id UUID,
	benefit_id UUID,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by   UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id),
	PRIMARY KEY(employee_id, benefit_id),
	FOREIGN KEY(employee_id) REFERENCES employees(id),
	FOREIGN KEY(benefit_id) REFERENCES benefits(id)
);

CREATE TABLE projects(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL UNIQUE,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE tasks(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	notes VARCHAR(255) NOT NULL,
	working_time TIMESTAMP NOT NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE projects_tasks(
	project_id UUID NOT NULL,
	task_id UUID NOT NULL,
	employee_id UUID NOT NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by   UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id),
	PRIMARY KEY (project_id, task_id, employee_id),
	FOREIGN KEY (project_id) REFERENCES projects(id),
	FOREIGN KEY (task_id) REFERENCES tasks(id),
	FOREIGN KEY	(employee_id) REFERENCES employees(id)
);

CREATE TABLE employees_projects(
	project_id UUID NOT NULL,
	employee_id UUID NOT NULL,
	roles VARCHAR(100) NULL,
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by   UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id),
	PRIMARY KEY (project_id, employee_id),
	FOREIGN KEY (project_id) REFERENCES projects(id),
	FOREIGN KEY	(employee_id) REFERENCES employees(id)
);

CREATE TABLE departments(
	name VARCHAR(100) PRIMARY KEY,
	address VARCHAR(255) NOT NULL,
	department_head_id UUID NOT NULL,	
	created_time TIMESTAMP NOT NULL DEFAULT NOW(),
	created_by UUID NOT NULL,
	FOREIGN KEY (department_head_id) REFERENCES employees(id),
	FOREIGN KEY(created_by) REFERENCES employees(id)
);

CREATE TABLE employee_departments(
	employee_id UUID,
	department_name VARCHAR(100),
	PRIMARY KEY(employee_id, department_name),
	FOREIGN KEY(employee_id) REFERENCES employees(id),
	FOREIGN KEY(department_name) REFERENCES departments(name)
);

CREATE TABLE level_defaults(
  level_id     UUID NOT NULL,
  value_type   VARCHAR(50) NOT NULL DEFAULT 'base_salary',
  amount       INT NOT NULL DEFAULT 0,
  created_time TIMESTAMP NOT NULL DEFAULT NOW(),
  created_by   UUID NOT NULL,
	FOREIGN KEY(created_by) REFERENCES employees(id),
	PRIMARY KEY (level_id, value_type),
	FOREIGN KEY(level_id) REFERENCES levels(id)
);
