package EmployeesGo

type Department struct {
	ID             uint
	DepartmentName string
	ManagerID      uint
	LocationID     uint
	CompanyID      uint
}
type DepartmentResponse struct {
	ID             uint
	DepartmentName string
	Manager        Employee
	Location       Location
	CompanyID      uint
}

func (r Department) TableName() string {
	return tablePrefix + "departments"
}
