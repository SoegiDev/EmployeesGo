package EmployeesGo

type Employee struct {
	ID           uint
	CustomID     string
	EmployeeID   string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	JobID        uint
	Job          Job
	Salary       uint64
	Picture      string
	ManagerID    uint
	DepartmentID uint
	Department   Department
}

func (a Employee) TableName() string {
	return tablePrefix + "employees"
}
