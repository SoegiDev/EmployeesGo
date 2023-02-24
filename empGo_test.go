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

func TestCreateRegion(t *testing.T) {
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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Create Region is Oke")
		}
	// test duplicated entries
	err_duplicate := empGo.CreateRegion(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}
	// clean up
	//db.Where("region_name = ?", data.RegionName).Delete(EmployeesGo.Region{})
}
func TestUpdateRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := map[string]interface{}{"region_name":"Asia Jakarta"}
	err := empGo.UpdateRegion(1,data)
	if err != nil {
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Update Region Ok")
	}

}

func TestDeleteRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test Delete Region
	err := empGo.DeleteRegion(1)
	if err != nil {
		if err.Error() == "Region Not Found"{
			fmt.Println("Region Not Found is Oke")
		}
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Delete Region Ok")
	}

}

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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Create Country is Oke")
	}
	// test duplicated entries
	err_duplicate := empGo.CreateCountry(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}
	// clean up
	//db.Where("country_name = ?", data.CountryName).Delete(EmployeesGo.Country{})
}

func TestUpdateCountry(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := map[string]interface{}{"country_name":"Indonesia","region_id":1}
	err := empGo.UpdateCountry(1,data)
	if err != nil {
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Update Country Ok")
	}
	// TEST ERROR
	data_1 := map[string]interface{}{"country_name":"Indonesia","region_id":1}
	err_1 := empGo.UpdateCountry(1,data_1)
	if err_1 != nil {
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Update Country Ok")
	}

}

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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
		fmt.Println("Create Location is Oke")
	}

	// test duplicated entries
	err_duplicate := empGo.CreateLocation(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}

	// clean up
	//db.Where("street_address = ?", data.StreetAddress).Delete(EmployeesGo.Location{})
}

func TestCreateDepartment(t *testing.T) {
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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Create Department is Oke")
		}

	// test duplicated entries
	err_duplicate := empGo.CreateDepartment(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}

	// clean up
	//db.Where("department_name = ?", data.DepartmentName).Delete(EmployeesGo.Department{})
}

func TestCreateJob(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	data := EmployeesGo.Job{
		JobTitle:  "Software Engineer",
		MinSalary: 10000000,
		MaxSalary: 19000000}
	err := empGo.CreateJob(data)
	if err != nil {
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Create Job is Oke")
		}

	// test duplicated entries
	err_duplicate := empGo.CreateJob(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}

	// clean up
	//db.Where("job_title = ?", data.JobTitle).Delete(EmployeesGo.Job{})
}

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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Create Employee is Oke")
		}

	// test duplicated entries
	err_duplicate := empGo.CreateEmployee(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}

	// clean up
	//db.Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).Delete(EmployeesGo.Employee{})
}

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
		if err.Error() == "invalid Transaction Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Create Job History is Oke")
		}

	// test duplicated entries
	err_duplicate:=empGo.CreateJobHistory(data)
	if err_duplicate != nil {
		if err_duplicate.Error() == "data Already Exist"{
			fmt.Println("Duplicate is Oke")
		}
		
	}

	// clean up
	//db.Where("job_id = ?", data.JobID).Delete(EmployeesGo.JobHistory{})
}

// RETRIEVE
// EMPLOYEE //
func TestGetEmployee(t *testing.T) {
	var table string = "get_employee"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var c int64

	var data uint = 1
	res := db.Model(EmployeesGo.Employee{}).Where("id = ?", data).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), res.Error)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s is Empty", table), res.Error)
	} else {
		var empData EmployeesGo.EmployeeResponse
		params := strconv.FormatUint(uint64(data), 10)
		empData, err := empGo.GetEmployeeId(params)
		if err != nil {
			t.Error(fmt.Sprintf("Error Get Data %s ", table), err)
		}
		fmt.Println(empData)
	}
}

func TestGetCountry(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 1
	var empData EmployeesGo.CountryResponse
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetCountryId(params)
	if err != nil {
		if err.Error() == "Country Not Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Get Country By Id is Oke")
			fmt.Println(empData)
		}
}

func TestGetRegion(t *testing.T) {
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var data uint = 2
	var empData EmployeesGo.Region
	params := strconv.FormatUint(uint64(data), 10)
	empData, err := empGo.GetRegionId(params)
	if err != nil {
		if err.Error() == "Region Not Found"{
			fmt.Println("Invalid Transaction is Oke")
		}
	}else{
			fmt.Println("Get Region By Id is Oke")
			fmt.Println(empData)
		}
}

func TestGetDepartment(t *testing.T) {
	var table string = "get Department"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var c int64

	var data uint = 1
	res := db.Model(EmployeesGo.Department{}).Where("id = ?", data).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), res.Error)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s is Empty", table), res.Error)
	} else {
		var empData EmployeesGo.DepartmentResponse
		params := strconv.FormatUint(uint64(data), 10)
		empData, err := empGo.GetDepartmentId(params)
		if err != nil {
			t.Error(fmt.Sprintf("Error Get Data %s ", table), err)
		}
		fmt.Println(empData)
	}
}

func TestGetLocation(t *testing.T) {
	var table string = "get location"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var c int64

	var data uint = 1
	res := db.Model(EmployeesGo.Location{}).Where("id = ?", data).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), res.Error)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s is Empty", table), res.Error)
	} else {
		var empData EmployeesGo.LocationResponse
		params := strconv.FormatUint(uint64(data), 10)
		empData, err := empGo.GetLocationId(params)
		if err != nil {
			t.Error(fmt.Sprintf("Error Get Data %s ", table), err)
		}
		fmt.Println(empData)
	}
}

func TestGetJob(t *testing.T) {
	var table string = "get country"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var c int64

	var data uint = 1
	res := db.Model(EmployeesGo.Job{}).Where("id = ?", data).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), res.Error)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s is Empty", table), res.Error)
	} else {
		var empData EmployeesGo.Job
		params := strconv.FormatUint(uint64(data), 10)
		empData, err := empGo.GetJobId(params)
		if err != nil {
			t.Error(fmt.Sprintf("Error Get Data %s ", table), err)
		}
		fmt.Println(empData)
	}
}

func TestGetJobHistory(t *testing.T) {
	var table string = "get country"
	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})
	var c int64

	var data uint = 1
	res := db.Model(EmployeesGo.JobHistory{}).Where("id = ?", data).Count(&c)
	if res.Error != nil {
		t.Error(fmt.Sprintf("unexpected error while storing %s: ", table), res.Error)
	}
	if c == 0 {
		t.Error(fmt.Sprintf("%s is Empty", table), res.Error)
	} else {
		var empData EmployeesGo.JobHistoryResponse
		params := strconv.FormatUint(uint64(data), 10)
		empData, err := empGo.GetJobHistoryId(params)
		if err != nil {
			t.Error(fmt.Sprintf("Error Get Data %s ", table), err)
		}
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
