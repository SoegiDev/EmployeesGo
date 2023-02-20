package EmployeesGo

type Job struct {
	ID        uint
	JobTitle  string
	MinSalary uint64
	MaxSalary uint64
}

func (r Job) TableName() string {
	return tablePrefix + "jobs"
}
