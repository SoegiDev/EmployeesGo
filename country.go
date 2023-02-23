package EmployeesGo

type Country struct {
	ID          uint
	CountryName string
	RegionID    uint
}

type CountryResponse struct {
	ID          uint
	CountryName string
	Region      Region
}

func (r Country) TableName() string {
	return tablePrefix + "countries"
}
