-- INSERT EMPLOYEES
INSERT INTO employees (email,full_name,code,gender,phone,dob,created_by)
VALUES
('emp01@test.com', 'Adam Miller', 'EMP001', 'male', '0900000001', to_date('12-01-1995','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp02@test.com', 'Bella Johnson', 'EMP002', 'female', '0900000002', to_date('23-02-1997','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp03@test.com', 'Chris Brown', 'EMP003', 'male', '0900000003', to_date('05-03-1994','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a' ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp04@test.com', 'Diana Wilson', 'EMP004', 'female', '0900000004', to_date('14-04-1998','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp05@test.com', 'Ethan Moore', 'EMP005', 'male', '0900000005', to_date('19-05-1996','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp06@test.com', 'Fiona Taylor', 'EMP006', 'female', '0900000006', to_date('08-06-1999','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp07@test.com', 'George Anderson', 'EMP007', 'male', '0900000007', to_date('30-07-1993','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp08@test.com', 'Hannah Thomas', 'EMP008', 'female', '0900000008', to_date('11-08-1997','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp09@test.com', 'Ian Jackson', 'EMP009', 'male', '0900000009', to_date('21-09-1995','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp10@test.com', 'Julia White', 'EMP010', 'female', '0900000010', to_date('02-10-1998','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp11@test.com', 'Kevin Harris', 'EMP011', 'male', '0900000011', to_date('17-11-1994','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp12@test.com', 'Laura Martin', 'EMP012', 'female', '0900000012', to_date('29-12-1996','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp13@test.com', 'Michael Clark', 'EMP013', 'male', '0900000013', to_date('04-01-1992','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp14@test.com', 'Nina Lewis', 'EMP014', 'female', '0900000014', to_date('16-02-1999','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp15@test.com', 'Oliver Walker', 'EMP015', 'male', '0900000015', to_date('07-03-1995','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp16@test.com', 'Paula Hall', 'EMP016', 'female', '0900000016', to_date('25-04-1997','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp17@test.com', 'Quentin Allen', 'EMP017', 'male', '0900000017', to_date('13-05-1993','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp18@test.com', 'Rachel Young', 'EMP018', 'female', '0900000018', to_date('09-06-1998','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp19@test.com', 'Samuel King', 'EMP019', 'male', '0900000019', to_date('18-07-1996','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('emp20@test.com', 'Tina Scott', 'EMP020', 'female', '0900000020', to_date('01-08-1999','DD-MM-YYYY'),
 CASE WHEN random() < 0.5 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT LEVELS
INSERT INTO levels (name, created_by)
VALUES
('Intern',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('Junior',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('Middle',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('Senior',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),
('Lead',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT POSITIONS
INSERT INTO positions (name, created_by)
VALUES
('Backend Developer',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Frontend Developer',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Mobile Developer',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('QA Engineer',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Project Manager',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT BRANCHES
INSERT INTO branches (name, created_by)
VALUES
('Head Office',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Hanoi Branch',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Ho Chi Minh Branch',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Da Nang Branch',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT PROJECTS
INSERT INTO projects (name, created_by)
VALUES
('Employee Attendance System',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Document Management System',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Inventory Management',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Customer Support Portal',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Performance Evaluation System',
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT TASKS
INSERT INTO tasks (name, notes, working_time, created_by)
VALUES
('Requirement Analysis',
 'Gather and analyze business requirements',
 6,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('System Architecture Design',
 'Design system architecture and components',
 10,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('API Integration',
 'Integrate third-party APIs',
 8,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Database Optimization',
 'Optimize queries and indexes',
 7,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('UI/UX Design',
 'Design wireframes and UI components',
 9,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Frontend Development',
 'Implement frontend features',
 14,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Backend Development',
 'Implement backend business logic',
 16,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Security Review',
 'Review system security and permissions',
 5,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Deployment',
 'Deploy system to production',
 4,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('Maintenance & Support',
 'Ongoing system maintenance',
 12,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT LEVEL DEFAULT
INSERT INTO level_defaults (level_id, value_type, amount, created_by)
VALUES
('b8c54109-adec-41a0-b7ce-01809fcd598b', 'base_salary', 2000000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('2ee71543-596e-4a00-8e5f-de046c927054', 'base_salary', 2100000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('eb1e0fd0-b322-4c71-a106-1d4adcd2e264', 'base_salary', 2200000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('7126fcac-0ef5-430d-9416-ade7cb8b00ca', 'base_salary', 2300000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('86a2ec4b-50a9-4d35-a275-50319db21766', 'base_salary', 2500000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('ac10d4b3-b449-48fb-8e46-366d1c99edeb', 'base_salary', 3000000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('e6de0d31-07e7-4f60-bb5b-e58d2a8114ee', 'base_salary', 4000000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END),

('dc706a9c-e421-4c26-9e7c-cf7462c50b01', 'base_salary', 5000000,
 CASE WHEN random() < 0.5
 THEN 'a746a1f5-764a-4ea9-a39c-12950c2b414a'::uuid
 ELSE '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'::uuid END);

-- INSERT EMPLOYEE_ROLE
INSERT INTO employees_roles(employee_id, level_id, position_id, branch_id, created_by)
VALUES
-- Chris Brown
('06764e90-4504-4ef1-b7f7-6331e4581c52', 'eb1e0fd0-b322-4c71-a106-1d4adcd2e264', 'aa247a37-e14c-4035-be60-775d90d6e134', '57612ed3-6eb5-4532-81f9-e084679067dd', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Diana Wilson
('49f8e9e6-1791-4a64-9914-582616a439ed', '7126fcac-0ef5-430d-9416-ade7cb8b00ca', '3488d94d-73d7-4491-abd9-7c34131ff2e0', '703d1b52-b079-453b-8ad8-ee2b3698b39c', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Ethan Moore
('4e310af6-e3d4-48c4-8bbe-319bbc04b940', '86a2ec4b-50a9-4d35-a275-50319db21766', 'c72a566c-da50-41fd-b0a8-9178acc8142f', '169a42b5-5fda-4575-9425-be6c331c4405', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Fiona Taylor
('6267d537-feb1-4cd1-8e0e-5fa79c498a48', 'ac10d4b3-b449-48fb-8e46-366d1c99edeb', '3205dcdf-e4a9-4de9-b793-3ee625108d44', '77fd1177-59df-4869-8846-83d9bb1f7420', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- George Anderson
('19418e72-8e06-4a2b-8bcd-f7bccbdbebc0', 'e6de0d31-07e7-4f60-bb5b-e58d2a8114ee', '44944680-f8d5-4e16-9333-e46b8dfbaf5e', '7bc88162-10d5-41ce-aa43-36864942d505', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Hannah Thomas
('d1ec4ab4-8671-491f-8780-d5d47786fa43', 'dc706a9c-e421-4c26-9e7c-cf7462c50b01', 'e883c467-8081-49ff-8a0d-af86174c866e', '57612ed3-6eb5-4532-81f9-e084679067dd', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Ian Jackson
('c138a32b-3405-4461-93f5-39e9c6740159', 'b8c54109-adec-41a0-b7ce-01809fcd598b', '3488d94d-73d7-4491-abd9-7c34131ff2e0', '703d1b52-b079-453b-8ad8-ee2b3698b39c', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Julia White
('3030f351-d5c4-4828-9b06-ca87ee95a466', '2ee71543-596e-4a00-8e5f-de046c927054', 'c72a566c-da50-41fd-b0a8-9178acc8142f', '169a42b5-5fda-4575-9425-be6c331c4405', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Kevin Harris
('36a87902-acd4-4ef5-b149-8e6193964c53', '7126fcac-0ef5-430d-9416-ade7cb8b00ca', 'aa247a37-e14c-4035-be60-775d90d6e134', '77fd1177-59df-4869-8846-83d9bb1f7420', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Laura Martin
('55c75319-fde6-4814-aaea-12abf50de442', '86a2ec4b-50a9-4d35-a275-50319db21766', '3205dcdf-e4a9-4de9-b793-3ee625108d44', '7bc88162-10d5-41ce-aa43-36864942d505', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Michael Clark
('b22bbb1a-1bf1-4574-a2e6-43c5f5bf5a71', 'ac10d4b3-b449-48fb-8e46-366d1c99edeb', '44944680-f8d5-4e16-9333-e46b8dfbaf5e', '57612ed3-6eb5-4532-81f9-e084679067dd', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Nina Lewis
('573d1f35-53b9-4280-9620-0d4fdcfee498', 'e6de0d31-07e7-4f60-bb5b-e58d2a8114ee', 'e883c467-8081-49ff-8a0d-af86174c866e', '703d1b52-b079-453b-8ad8-ee2b3698b39c', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Oliver Walker
('784df0e4-f862-457a-9a06-4de96cabf9e4', 'dc706a9c-e421-4c26-9e7c-cf7462c50b01', '3488d94d-73d7-4491-abd9-7c34131ff2e0', '169a42b5-5fda-4575-9425-be6c331c4405', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Paula Hall
('1236591c-f4ad-4957-8437-b917004221ed', 'b8c54109-adec-41a0-b7ce-01809fcd598b', 'c72a566c-da50-41fd-b0a8-9178acc8142f', '77fd1177-59df-4869-8846-83d9bb1f7420', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Quentin Allen
('754bb477-94ae-4090-a9ad-8aec03979f8b', '2ee71543-596e-4a00-8e5f-de046c927054', 'aa247a37-e14c-4035-be60-775d90d6e134', '7bc88162-10d5-41ce-aa43-36864942d505', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Rachel Young
('f2b3d5c4-e452-45dc-b12c-5577534b2e84', 'eb1e0fd0-b322-4c71-a106-1d4adcd2e264', '44944680-f8d5-4e16-9333-e46b8dfbaf5e', '57612ed3-6eb5-4532-81f9-e084679067dd', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
-- Samuel Kin
('40822a0f-b3ef-46b7-9397-4c01da3214ec', '7126fcac-0ef5-430d-9416-ade7cb8b00ca', '3488d94d-73d7-4491-abd9-7c34131ff2e0', '703d1b52-b079-453b-8ad8-ee2b3698b39c', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
-- Adam Miller
('3beb8d1d-1cef-4901-b7d2-38b700011024', '86a2ec4b-50a9-4d35-a275-50319db21766', 'c72a566c-da50-41fd-b0a8-9178acc8142f', '169a42b5-5fda-4575-9425-be6c331c4405', 'a746a1f5-764a-4ea9-a39c-12950c2b414a');

-- INSERT PROJECT_TASKS
INSERT INTO projects_tasks (project_id, task_id, employee_id, created_by) VALUES
-- Employee Attendance System tasks
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', 'a6863b90-4f1c-49f9-9140-895c3dc092ce', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', '5138ec39-6288-4620-85c1-4d68573e36f4', '06764e90-4504-4ef1-b7f7-6331e4581c52', '06764e90-4504-4ef1-b7f7-6331e4581c52'),

-- Document Management System tasks
('32b6a17b-43b6-4055-bf9f-aa31534f2455', '71353492-b81f-41ef-8a32-e682830502d1', '49f8e9e6-1791-4a64-9914-582616a439ed', '49f8e9e6-1791-4a64-9914-582616a439ed'),

-- Inventory Management tasks
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', '78f835c2-5e14-460f-b685-51b9d90e8925', '4e310af6-e3d4-48c4-8bbe-319bbc04b940', '4e310af6-e3d4-48c4-8bbe-319bbc04b940'),

-- Customer Support Portal tasks
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '5f422f5e-dd6d-4a63-a500-66bc484c639c', '6267d537-feb1-4cd1-8e0e-5fa79c498a48', '6267d537-feb1-4cd1-8e0e-5fa79c498a48'),
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '71353492-b81f-41ef-8a32-e682830502d1', '19418e72-8e06-4a2b-8bcd-f7bccbdbebc0', '19418e72-8e06-4a2b-8bcd-f7bccbdbebc0'),

-- Performance Evaluation System tasks
('3e06b366-5f92-4e1c-bb74-5972804b037a', '09f5f998-a88e-4658-a992-49d5bf17ae27', 'd1ec4ab4-8671-491f-8780-d5d47786fa43', 'd1ec4ab4-8671-491f-8780-d5d47786fa43'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '2915a366-ddd7-4d37-af1e-5371a07e3004', 'c138a32b-3405-4461-93f5-39e9c6740159', 'c138a32b-3405-4461-93f5-39e9c6740159'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '81d0a1e1-5f17-435d-bc2b-826bdfcc85de', '3030f351-d5c4-4828-9b06-ca87ee95a466', '3030f351-d5c4-4828-9b06-ca87ee95a466'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '8de15acd-4782-47d2-a98d-0b4c8c3b5484', '36a87902-acd4-4ef5-b149-8e6193964c53', '36a87902-acd4-4ef5-b149-8e6193964c53'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '3b6c628e-88f0-41c0-b1e4-1291e602fd17', '55c75319-fde6-4814-aaea-12abf50de442', '55c75319-fde6-4814-aaea-12abf50de442');

-- INSERT EMPLOYEE_PROJECT
INSERT INTO employees_projects (project_id, employee_id, roles, created_by) VALUES
-- Employee Attendance System
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', 'a746a1f5-764a-4ea9-a39c-12950c2b414a', 'Project Manager', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916', 'Developer', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', '06764e90-4504-4ef1-b7f7-6331e4581c52', 'Tester', '06764e90-4504-4ef1-b7f7-6331e4581c52'),
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', '49f8e9e6-1791-4a64-9914-582616a439ed', 'QA', '49f8e9e6-1791-4a64-9914-582616a439ed'),
('b0177b6d-0b3a-4ee4-93ac-7a6e6f5e6de2', '4e310af6-e3d4-48c4-8bbe-319bbc04b940', 'Business Analyst', '4e310af6-e3d4-48c4-8bbe-319bbc04b940'),

-- Document Management System
('32b6a17b-43b6-4055-bf9f-aa31534f2455', '6267d537-feb1-4cd1-8e0e-5fa79c498a48', 'Project Manager', '6267d537-feb1-4cd1-8e0e-5fa79c498a48'),
('32b6a17b-43b6-4055-bf9f-aa31534f2455', '19418e72-8e06-4a2b-8bcd-f7bccbdbebc0', 'Developer', '19418e72-8e06-4a2b-8bcd-f7bccbdbebc0'),
('32b6a17b-43b6-4055-bf9f-aa31534f2455', 'd1ec4ab4-8671-491f-8780-d5d47786fa43', 'Tester', 'd1ec4ab4-8671-491f-8780-d5d47786fa43'),
('32b6a17b-43b6-4055-bf9f-aa31534f2455', 'c138a32b-3405-4461-93f5-39e9c6740159', 'QA', 'c138a32b-3405-4461-93f5-39e9c6740159'),
('32b6a17b-43b6-4055-bf9f-aa31534f2455', '3030f351-d5c4-4828-9b06-ca87ee95a466', 'Business Analyst', '3030f351-d5c4-4828-9b06-ca87ee95a466'),

-- Inventory Management
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', '36a87902-acd4-4ef5-b149-8e6193964c53', 'Project Manager', '36a87902-acd4-4ef5-b149-8e6193964c53'),
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', '55c75319-fde6-4814-aaea-12abf50de442', 'Developer', '55c75319-fde6-4814-aaea-12abf50de442'),
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', 'b22bbb1a-1bf1-4574-a2e6-43c5f5bf5a71', 'Tester', 'b22bbb1a-1bf1-4574-a2e6-43c5f5bf5a71'),
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', '573d1f35-53b9-4280-9620-0d4fdcfee498', 'QA', '573d1f35-53b9-4280-9620-0d4fdcfee498'),
('fc69c30b-5039-4bb1-9af0-abdd90d717d1', '784df0e4-f862-457a-9a06-4de96cabf9e4', 'Business Analyst', '784df0e4-f862-457a-9a06-4de96cabf9e4'),

-- Customer Support Portal
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '1236591c-f4ad-4957-8437-b917004221ed', 'Project Manager', '1236591c-f4ad-4957-8437-b917004221ed'),
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '754bb477-94ae-4090-a9ad-8aec03979f8b', 'Developer', '754bb477-94ae-4090-a9ad-8aec03979f8b'),
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', 'f2b3d5c4-e452-45dc-b12c-5577534b2e84', 'Tester', 'f2b3d5c4-e452-45dc-b12c-5577534b2e84'),
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '40822a0f-b3ef-46b7-9397-4c01da3214ec', 'QA', '40822a0f-b3ef-46b7-9397-4c01da3214ec'),
('bb3ee3ad-ed71-4daf-8789-d7657c46a5e3', '3beb8d1d-1cef-4901-b7d2-38b700011024', 'Business Analyst', '3beb8d1d-1cef-4901-b7d2-38b700011024'),

-- Performance Evaluation System
('3e06b366-5f92-4e1c-bb74-5972804b037a', '5feae934-0032-41ea-9ce9-dc0d07491baa', 'Project Manager', '5feae934-0032-41ea-9ce9-dc0d07491baa'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', 'b85fedd7-74d8-4672-84b2-a575ee780339', 'Developer', 'b85fedd7-74d8-4672-84b2-a575ee780339'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', 'a746a1f5-764a-4ea9-a39c-12950c2b414a', 'Tester', 'a746a1f5-764a-4ea9-a39c-12950c2b414a'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916', 'QA', '5e5a2d85-07d9-44ab-a7ca-d2dd3e580916'),
('3e06b366-5f92-4e1c-bb74-5972804b037a', '06764e90-4504-4ef1-b7f7-6331e4581c52', 'Business Analyst', '06764e90-4504-4ef1-b7f7-6331e4581c52');

-- UPDATE TASKS DATA
UPDATE tasks SET working_time = 6
WHERE id = 'a6863b90-4f1c-49f9-9140-895c3dc092ce';

UPDATE tasks SET working_time = 8
WHERE id = '5138ec39-6288-4620-85c1-4d68573e36f4';

UPDATE tasks SET working_time = 5
WHERE id = '71353492-b81f-41ef-8a32-e682830502d1';

UPDATE tasks SET working_time = 4
WHERE id = '78f835c2-5e14-460f-b685-51b9d90e8925';

UPDATE tasks SET working_time = 7
WHERE id = '5f422f5e-dd6d-4a63-a500-66bc484c639c';

UPDATE tasks SET working_time = 8
WHERE id = '09f5f998-a88e-4658-a992-49d5bf17ae27';

UPDATE tasks SET working_time = 8
WHERE id = '2915a366-ddd7-4d37-af1e-5371a07e3004';

UPDATE tasks SET working_time = 3
WHERE id = '81d0a1e1-5f17-435d-bc2b-826bdfcc85de';

UPDATE tasks SET working_time = 2
WHERE id = '8de15acd-4782-47d2-a98d-0b4c8c3b5484';

UPDATE tasks SET working_time = 6
WHERE id = '3b6c628e-88f0-41c0-b1e4-1291e602fd17';

