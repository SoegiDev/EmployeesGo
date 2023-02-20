package EmployeesGo

type Country struct {
	ID          uint
	CountryName string
	RegionID    uint
}

func (r Country) TableName() string {
	return tablePrefix + "countries"
}
