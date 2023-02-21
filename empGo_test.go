package EmployeesGo_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	EmployeesGo "github.com/SoegiDev/EmployeesGo"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var prefix_test string = "employeeGo_"

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var dsn string
	if os.Getenv("env") == "testing" {
		fmt.Println("preparing testing config...")
		dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=employee_go port=5432 sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DbUser"), os.Getenv("DbPassword"))
	} else {
		dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=employee_go port=5432 sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DbUser"), os.Getenv("DbPassword"))
	}

	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestCreateRegion(t *testing.T) {
	var table string = "Region"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Region{
		RegionName: "Asia Pasific",
	}
	err := empGo.CreateRegion(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Region{}).Where("region_name = ?", data.RegionName).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateRegion(data)
	empGo.CreateRegion(data)
	empGo.CreateRegion(data)
	db.Model(EmployeesGo.Region{}).Where("region_name = ?", data.RegionName).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("region_name = ?", data.RegionName).Delete(EmployeesGo.Region{})
}

func TestCreateCountry(t *testing.T) {

	var table string = "country"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Country{
		CountryName: "Indonesia",
		RegionID:    1,
	}
	err := empGo.CreateCountry(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Country{}).Where("country_name = ?", data.CountryName).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateCountry(data)
	empGo.CreateCountry(data)
	empGo.CreateCountry(data)
	db.Model(EmployeesGo.Country{}).Where("country_name = ?", data.CountryName).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("country_name = ?", data.CountryName).Delete(EmployeesGo.Country{})
}

func TestCreateLocation(t *testing.T) {

	var table string = "location"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Location{
		StreetAddress: "Jl. Gading 2",
		PostalCode:    "12312",
		City:          "Jakarta Timur",
		StateProvince: "DKI Jakarta",
		CountryID:     1}
	err := empGo.CreateLocation(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Location{}).Where("street_address = ?", data.StreetAddress).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateLocation(data)
	empGo.CreateLocation(data)
	empGo.CreateLocation(data)
	db.Model(EmployeesGo.Location{}).Where("street_address = ?", data.StreetAddress).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("street_address = ?", data.StreetAddress).Delete(EmployeesGo.Location{})
}

func TestCreateDepartment(t *testing.T) {

	var table string = "department"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Department{
		DepartmentName: "ICT",
		ManagerID:      1,
		LocationID:     1,
		CompanyID:      1}
	err := empGo.CreateDepartment(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Department{}).Where("department_name = ?", data.DepartmentName).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateDepartment(data)
	empGo.CreateDepartment(data)
	empGo.CreateDepartment(data)
	db.Model(EmployeesGo.Department{}).Where("department_name = ?", data.DepartmentName).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("department_name = ?", data.DepartmentName).Delete(EmployeesGo.Department{})
}

func TestCreateEmployee(t *testing.T) {

	var table string = "employee"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Employee{
		CustomID:     "312312312",
		EmployeeID:   "12312312323",
		FirstName:    "Test",
		LastName:     "Name",
		Email:        "test@gmail.com",
		PhoneNumber:  "12312312312",
		JobID:        1,
		Salary:       90000000,
		Picture:      "",
		ManagerID:    1,
		DepartmentID: 1}
	err := empGo.CreateEmployee(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Employee{}).Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateEmployee(data)
	empGo.CreateEmployee(data)
	empGo.CreateEmployee(data)
	db.Model(EmployeesGo.Employee{}).Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("fist_name = ? and last_name = ?", data.FirstName, data.LastName).Delete(EmployeesGo.Employee{})
}

func TestCreateJobHistory(t *testing.T) {

	var table string = "job_history"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Employee{
		CustomID:     "312312312",
		EmployeeID:   "12312312323",
		FirstName:    "Test",
		LastName:     "Name",
		Email:        "test@gmail.com",
		PhoneNumber:  "12312312312",
		JobID:        1,
		Salary:       90000000,
		Picture:      "",
		ManagerID:    1,
		DepartmentID: 1}
	err := empGo.CreateEmployee(data)
	if err != nil {
		t.Error(fmt.Sprintf("an error was not expected while creating %s ", table), err)
	}

	var c int64
	res := db.Model(EmployeesGo.Employee{}).Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), err)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s has not been stored", table), err)
	}

	// test duplicated entries
	empGo.CreateEmployee(data)
	empGo.CreateEmployee(data)
	empGo.CreateEmployee(data)
	db.Model(EmployeesGo.Employee{}).Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).Count(&c)
	if c > 1 {
		t.Error(fmt.Sprintf("unexpected duplicated entries for %s", table), err)
	}

	// clean up
	db.Where("fist_name = ? and last_name = ?", data.FirstName, data.LastName).Delete(EmployeesGo.Employee{})
}
