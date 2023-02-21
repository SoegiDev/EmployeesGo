package EmployeesGo

import "time"

type JobHistory struct {
	ID           uint
	EmployeeID   uint
	StartDate    time.Time
	EndDate      time.Time
	JobID        uint
	DepartmentID uint
}

func (r JobHistory) TableName() string {
	return tablePrefix + "job_history"
}
