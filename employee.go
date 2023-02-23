package EmployeesGo

import ("time")

type Employee struct {
	ID           uint
	CustomID     string `gorm:"size:50;default:null;unique"`
	EmployeeID   string `gorm:"size:50;default:null;unique"`
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	JobID        uint
	Salary       uint64
	Picture      string
	Active bool
	CreatedDate time.Time
	Contract bool
	ManagerID    uint
	DepartmentID uint
}

type EmployeeResponse struct {
	ID          uint
	CustomID    string
	EmployeeID  string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	JobID       uint
	Salary      uint64
	Picture     string
	ManagerID   uint
	Department  Department
	Job         Job
}

func (a Employee) TableName() string {
	return tablePrefix + "employees"
}
