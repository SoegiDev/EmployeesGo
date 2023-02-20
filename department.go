package EmployeesGo

type Department struct {
	ID             uint
	DepartmentName string
	ManagerID      uint
	LocationID     uint
	CompanyID      uint
}

func (r Department) TableName() string {
	return tablePrefix + "departments"
}
