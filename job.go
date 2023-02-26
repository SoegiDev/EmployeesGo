package EmployeesGo

type Job struct {
	ID             uint
	JobTitle       string
	JobDescription string `gorm:"type:text;unique"`
	MinSalary      uint64
	MaxSalary      uint64
}

type JobResponse struct {
	ID             uint
	JobTitle       string
	JobDescription string
	MinSalary      uint64
	MaxSalary      uint64
}

func (r Job) TableName() string {
	return tablePrefix + "jobs"
}
