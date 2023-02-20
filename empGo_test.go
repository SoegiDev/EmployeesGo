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

func TestCreateRole(t *testing.T) {

	empGo := EmployeesGo.New(EmployeesGo.EmpOption{
		TablesPrefix: prefix_test,
		DB:           db,
	})

	// test create role
	region := empGo.Re
	err := empGo.CreateRegion("role-a")
	if err != nil {
		t.Error("an error was not expected while creating role ", err)
	}

	var c int64
	res := db.Model(AuthorizationGo.Role{}).Where("name = ?", "role-a").Count(&c)
	if res.Error != nil {
		t.Error("unexpected error while storing role: ", err)
	}
	if c == 0 {
		t.Error("role has not been stored")
	}

	// test duplicated entries
	auth.CreateRole("role-a")
	auth.CreateRole("role-a")
	auth.CreateRole("role-a")
	db.Model(AuthorizationGo.Role{}).Where("name = ?", "role-a").Count(&c)
	if c > 1 {
		t.Error("unexpected duplicated entries for role")
	}

	// clean up
	db.Where("name = ?", "role-a").Delete(AuthorizationGo.Role{})
}
