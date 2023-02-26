package EmployeesGo_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/SoegiDev/EmployeesGo"
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

// TEST REGION
func TestCreateRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	data := EmployeesGo.Region{
		RegionName: "Asia Pasific",
	}
	err := empGo.CreateRegion(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Region is Oke")
	}
	// test duplicated entries
	err_duplicate := empGo.CreateRegion(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}
func TestUpdateRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := map[string]interface{}{"region_name": "Asia Jakarta"}
	err := empGo.UpdateRegion(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Region Ok")
	}
}
func TestGetRegionById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.Region
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetRegionId(params)
	if err != nil {
		if err.Error() == "Region Not Found" {
			fmt.Println("Region Not Found is Oke")
		}
	} else {
		fmt.Println("Get Region By Id is Oke")
		fmt.Println(empData)
	}
}
func TestGetRegionAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Region
	empData, err := empGo.GetRegion()
	if err != nil {
		if err.Error() == "Region Not Found" {
			fmt.Println("Region Not Found is Oke")
		}
	} else {
		fmt.Println("Get Region All is Oke")
		fmt.Println(empData)
	}
}

// TEST COUNTRY
func TestCreateCountry(t *testing.T) {

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
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Country is Oke")
	}
	// test duplicated entries
	err_duplicate := empGo.CreateCountry(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}
func TestGetCountryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.CountryResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetCountryId(params)
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Get Country By Id is Oke")
		fmt.Println(empData)
	}
}
func TestGetCountryByAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Country
	empData, err := empGo.GetCountry()
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Country Not Found is Oke")
		}
	} else {
		fmt.Println("Get Country By Id is Oke")
		fmt.Println(empData)
	}
}
func TestUpdateCountry(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := map[string]interface{}{"country_name": "Indonesia", "region_id": 1}
	err := empGo.UpdateCountry(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Country Ok")
	}
	// TEST ERROR
	data_1 := map[string]interface{}{"country_name": "Indonesia", "region_id": 1}
	err_1 := empGo.UpdateCountry(1, data_1)
	if err_1 != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Country Ok")
	}

}

// TEST LOCATION
func TestCreateLocation(t *testing.T) {
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
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Location is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateLocation(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}
func TestGetLocationById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.LocationResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetLocationId(params)
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Get Location Not Found is Oke")
		}
	} else {
		fmt.Println("Get Location By Id is Oke")
		fmt.Println(empData)
	}
}
func TestGetLocationAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Location
	empData, err := empGo.GetLocation()
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Get Location Not Found is Oke")
		}
	} else {
		fmt.Println("Get Location By Id is Oke")
		fmt.Println(empData)
	}
}

func TestUpdateLocation(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := map[string]interface{}{"street_address": "Indonesia", "postal_code": "21321", "city": "Jakarta", "state_province": "Jakarta", "country_id": 1}
	err := empGo.UpdateLocation(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Location Ok")
	}
}

// TEST DEPARTMENT
func TestCreateDepartment(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Department{
		DepartmentName: "ICT",
		ManagerID:      0,
		LocationID:     1,
		CompanyID:      1}
	err := empGo.CreateDepartment(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Department is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateDepartment(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}

func TestGetDepartmentById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.DepartmentResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetDepartmentId(params)
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Get Department Not Found is Oke")
		}
	} else {
		fmt.Println("Get Department By Id is Oke")
		fmt.Println(empData)
	}
}

func TestGetDepartmentAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Department
	empData, err := empGo.GetDepartment()
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Get Department Not Found is Oke")
		}
	} else {
		fmt.Println("Get Department By Id is Oke")
		fmt.Println(empData)
	}
}

func TestUpdateDepartment(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := map[string]interface{}{"department_name": "ITC", "manager_id": 0, "location_id": 1, "company_id": 1}
	err := empGo.UpdateDepartment(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Department Ok")
	}
}

// TEST JOB
func TestCreateJob(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := EmployeesGo.Job{
		JobTitle:  "Software Engineer",
		MinSalary: 10000000,
		MaxSalary: 19000000}
	err := empGo.CreateJob(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Job is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateJob(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}
func TestGetJobById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.JobResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetJobId(params)
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Get Job Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job By Id is Oke")
		fmt.Println(empData)
	}
}

func TestGetJobAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Job
	empData, err := empGo.GetJob()
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Get Job Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job By Id is Oke")
		fmt.Println(empData)
	}
}

func TestUpdateJob(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := map[string]interface{}{"job_title": "Programmer", "job_description": "IT Programmer", "min_salary": 900000, "max_salary": 19900000}
	err := empGo.UpdateJob(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Job Ok")
	}
}

// TEST EMPLOYEE
func TestCreateEmployee(t *testing.T) {
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
		ManagerID:    0,
		DepartmentID: 1}
	err := empGo.CreateEmployee(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Employee is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateEmployee(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}

func TestGetEmployeeById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.EmployeeResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetEmployeeId(params)
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Get Employee Not Found is Oke")
		}
	} else {
		fmt.Println("Get Employee By Id is Oke")
		fmt.Println(empData)
	}
}

func TestGetEmployeeAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.Employee
	empData, err := empGo.GetEmployee()
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Get Employee Not Found is Oke")
		}
	} else {
		fmt.Println("Get Employee By Id is Oke")
		fmt.Println(empData)
	}
}

func TestUpdateEmployee(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := map[string]interface{}{"employee_id": "12312321", "first_name": "Test", "last_name": "user"}
	err := empGo.UpdateEmployee(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Employee Ok")
	}
}

// TEST JOB HISTORY
func TestCreateJobHistory(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.JobHistory{
		EmployeeID:   1,
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		JobID:        1,
		DepartmentID: 1}
	err := empGo.CreateJobHistory(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Create Job History is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateJobHistory(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist" {
			fmt.Println("Duplicate is Oke")
		}

	}
}

func TestGetJobHistoryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.JobHistoryResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetJobHistoryId(params)
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Get Job History Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job History By Id is Oke")
		fmt.Println(empData)
	}
}

func TestGetJobHistoryAll(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var empData []EmployeesGo.JobHistory
	empData, err := empGo.GetJobHistory()
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Get Job History Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job History By Id is Oke")
		fmt.Println(empData)
	}
}

func TestUpdateJobHistory(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	data := map[string]interface{}{"employee_id": 1, "job_id": 1, "department_id": 1}
	err := empGo.UpdateJobHistory(1, data)
	if err != nil {
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Update Job History Ok")
	}
}

// TEST DELETE

func TestDeleteRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteRegion(1)
	if err != nil {
		if err.Error() == "Region Not Found" {
			fmt.Println("Region Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Region Ok")
	}
}
func TestDeleteCountry(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteCountry(1)
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Country Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Country Ok")
	}
}

func TestDeleteLocation(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteLocation(1)
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Location Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Location Ok")
	}
}
func TestDeleteDepartment(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteDepartment(1)
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Department Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Department Ok")
	}
}
func TestDeleteJob(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteJob(1)
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Job Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Job Ok")
	}
}
func TestDeleteJobHistory(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteJobHistory(1)
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Job History Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Job History Ok")
	}
}
func TestDeleteEmployee(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteEmployee(1)
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Employee Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Employee Ok")
	}
}

// TEST NOT FOUND BY ID

func TestNotFoundEmployee(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.EmployeeResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetEmployeeId(params)
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Get Employee Not Found is Oke")
		}
	} else {
		fmt.Println("Get Employee By Id is Oke")
		fmt.Println(empData)
	}
}
func TestNotFoundJobById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.JobResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetJobId(params)
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Get Job Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundDepartmentById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.DepartmentResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetDepartmentId(params)
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Get Department Not Found is Oke")
		}
	} else {
		fmt.Println("Get Department By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundRegionById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.Region
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetRegionId(params)
	if err != nil {
		if err.Error() == "Region Not Found" {
			fmt.Println("Region Not Found is Oke")
		}
	} else {
		fmt.Println("Get Region By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundLocationById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.LocationResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetLocationId(params)
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Get Location Not Found is Oke")
		}
	} else {
		fmt.Println("Get Location By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundJobHistoryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.JobHistoryResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetJobHistoryId(params)
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Get Job History Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job History By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundCountryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.CountryResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetCountryId(params)
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Get Country By Id is Oke")
		fmt.Println(empData)
	}
}

// TEST NOT FOUND DELETE
func TestNotFoundDeleteRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteRegion(1)
	if err != nil {
		if err.Error() == "Region Not Found" {
			fmt.Println("Region Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Region Ok")
	}
}
func TestNotFoundDeleteCountry(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteCountry(1)
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Country Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Country Ok")
	}
}
func TestNotFoundDeleteDepartment(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteDepartment(1)
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Department Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Department Ok")
	}
}
func TestNotFoundDeleteJob(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteJob(1)
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Job Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Job Ok")
	}
}
func TestNotFoundDeleteJobHistory(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteJobHistory(1)
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Job History Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Job History Ok")
	}
}
func TestNotFoundDeleteEmployee(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteEmployee(1)
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Employee Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Employee Ok")
	}
}
func TestNotFoundDeleteLocation(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteLocation(1)
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Location Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found" {
			fmt.Println("invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Delete Location Ok")
	}
}

// TEST NOT FOUND PARAMETER ID
func TestNotFoundGetJobById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetJobId("")
	if err != nil {
		if err.Error() == "Job Not Found" {
			fmt.Println("Get Job Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job By Id is Oke")
		fmt.Println(empData)
	}
}
func TestNotFoundGetJobHistoryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetJobHistoryId("")
	if err != nil {
		if err.Error() == "Job History Not Found" {
			fmt.Println("Get Job History Not Found is Oke")
		}
	} else {
		fmt.Println("Get Job History By Id is Oke")
		fmt.Println(empData)
	}
}
func TestNotFoundGetCountryById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetCountryId("")
	if err != nil {
		if err.Error() == "Country Not Found" {
			fmt.Println("Invalid Transaction is Oke")
		}
	} else {
		fmt.Println("Get Country By Id is Oke")
		fmt.Println(empData)
	}
}
func TestNotFoundGetEmployeeById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetEmployeeId("")
	if err != nil {
		if err.Error() == "Employee Not Found" {
			fmt.Println("Get Employee Not Found is Oke")
		}
	} else {
		fmt.Println("Get Employee By Id is Oke")
		fmt.Println(empData)
	}
}
func TestNotFoundGetDepartmentById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetDepartmentId("")
	if err != nil {
		if err.Error() == "Department Not Found" {
			fmt.Println("Get Department Not Found is Oke")
		}
	} else {
		fmt.Println("Get Department By Id is Oke")
		fmt.Println(empData)
	}
}

func TestNotFoundGetLocationById(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	empData, err := empGo.GetLocationId("")
	if err != nil {
		if err.Error() == "Location Not Found" {
			fmt.Println("Get Location Not Found is Oke")
		}
	} else {
		fmt.Println("Get Location By Id is Oke")
		fmt.Println(empData)
	}
}
func TestCleanUp(t *testing.T) {
	tbl_regions := `"employeeGo_regions"`
	tbl_countries := `"employeeGo_countries"`
	tbl_locations := `"employeeGo_locations"`
	tbl_departments := `"employeeGo_departments"`
	tbl_jobs := `"employeeGo_jobs"`
	tbl_employees := `"employeeGo_employees"`
	tbl_job_history := `"employeeGo_job_history"`
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_regions))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_countries))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_locations))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_departments))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_jobs))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_employees))
	db.Exec(fmt.Sprintf("truncate table %s restart identity;", tbl_job_history))
}
