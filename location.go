package EmployeesGo

type Location struct {
	ID            uint
	StreetAddress string
	PostalCode    string
	City          string
	StateProvince string
	CountryID     uint
}

func (r Location) TableName() string {
	return tablePrefix + "locations"
}
