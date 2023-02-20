package EmployeesGo

type Region struct {
	ID         uint
	RegionName string
}

func (r Region) TableName() string {
	return tablePrefix + "regions"
}
