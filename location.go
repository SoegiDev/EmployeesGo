package EmployeesGo

type Location struct {
	ID            uint
	StreetAddress string
	PostalCode    string
	City          string
	StateProvince string
	CountryID     uint
}

type LocationResponse struct {
	ID            uint
	StreetAddress string
	PostalCode    string
	City          string
	StateProvince string
	Country       Country
}

func (r Location) TableName() string {
	return tablePrefix + "locations"
}
