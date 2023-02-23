package EmployeesGo

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

var tablePrefix string

type EmployeeSet struct {
	DB *gorm.DB
}

type EmpOption struct {
	TablesPrefix string
	DB           *gorm.DB
}

var empSet *EmployeeSet

func New(em EmpOption) *EmployeeSet {
	tablePrefix = em.TablesPrefix
	empSet = &EmployeeSet{DB: em.DB}
	migrateTables(em.DB)
	return empSet
}

var (
	ErrDepartmentNotFound = errors.New("Department Not Found")
	ErrCountryNotFound    = errors.New("Country Not Found")
	ErrLocationNotFound   = errors.New("Location Not Found")
	ErrEmployeeNotFound   = errors.New("Employee Not Found")
	ErrParameterNotFound  = errors.New("parameter Not Found")
	ErrJobNotFound        = errors.New("Job Not Found")
	ErrJobHistoryNotFound = errors.New("Job History Not Found")
	ErrRegionNotFound     = errors.New("Region Not Found")
	ErrDataIsExisted          = errors.New("data Already Exist")
)

func migrateTables(db *gorm.DB) {
	db.AutoMigrate(&Region{})
	db.AutoMigrate(&Country{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Department{})
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&JobHistory{})
	db.AutoMigrate(&Job{})
}

func Resolve() *EmployeeSet {
	return empSet
}

// DEPARTMENT //
func (a *EmployeeSet) CreateDepartment(data Department) error {
	var dbDepartment Department
	//res := a.DB.Where("department_name = ?", data.DepartmentName).First(&dbDepartment)
	res := a.DB.FirstOrCreate(&dbDepartment, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted

}

func (a *EmployeeSet) UpdateDepartment(departmentId uint, data map[string]interface{}) error {
	var dbDepartment Department
	res := a.DB.Model(&dbDepartment).Where("id = ?", departmentId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteDepartment(departmentId uint) error {
	var dbDepartment Department
	res := a.DB.Where("id = ?", departmentId).First(&dbDepartment)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrDepartmentNotFound
		}

	}
	// Delete the Department
	a.DB.Where("id = ?", departmentId).Delete(Department{})

	return nil
}

func (a *EmployeeSet) GetDepartmentId(id string) (DepartmentResponse, error) {
	var dbDepartment Department
	var departmentResponse DepartmentResponse
	var dbLocation Location
	var dbEmp Employee
	if NullString(id) == nil {
		return departmentResponse, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	departmentId := uint(conv)

	res := a.DB.Where("id = ?", departmentId).First(&dbDepartment)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return departmentResponse, ErrDepartmentNotFound
		}
	}
	if dbDepartment.LocationID > 0 {
		_ = a.DB.Where("id = ?", dbDepartment.LocationID).First(&dbLocation)
	}
	if dbDepartment.ManagerID > 0 {
		_ = a.DB.Where("id = ?", dbDepartment.ManagerID).First(&dbEmp)
	}
	departmentResponse = DepartmentResponse{
		ID:             dbDepartment.ID,
		DepartmentName: dbDepartment.DepartmentName,
		Manager:        dbEmp,
		Location:       dbLocation,
		CompanyID:      dbDepartment.CompanyID}
	// Get the Department Single Data
	return departmentResponse, nil
}

func (a *EmployeeSet) GetDepartment() ([]Department, error) {
	var dbDepartment []Department
	res := a.DB.Find(&dbDepartment)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbDepartment, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbDepartment, nil
}

// COUNTRY //
func (a *EmployeeSet) CreateCountry(data Country) error {
	var dbCountry Country
	//res := a.DB.Where("country_name = ?", data.CountryName).First(&dbCountry)
	res := a.DB.FirstOrCreate(&dbCountry, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateCountry(countryId uint, data map[string]interface{}) error {
	var dbCountry Country
	res := a.DB.Model(&dbCountry).Where("id = ?", countryId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteCountry(countryId uint) error {
	var dbCountry Country
	res := a.DB.Where("id = ?", countryId).First(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrCountryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", countryId).Delete(Country{})
	return nil
}

func (a *EmployeeSet) GetCountryId(id string) (CountryResponse, error) {
	var dbCountry Country
	var countryResponse CountryResponse
	var dbRegion Region
	if NullString(id) == nil {
		return countryResponse, ErrParameterNotFound
	}
	conv, _ := strconv.Atoi(id)
	countryId := uint(conv)

	res := a.DB.Where("id = ?", countryId).First(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return countryResponse, ErrCountryNotFound
		}

	}
	if dbCountry.RegionID > 0 {
		_ = a.DB.Where("id = ?", dbCountry.RegionID).First(&dbRegion)
	}

	countryResponse = CountryResponse{
		ID:          dbCountry.ID,
		CountryName: dbCountry.CountryName,
		Region:      dbRegion}
	// Get the Country Single Data
	return countryResponse, nil
}

func (a *EmployeeSet) GetCountry() ([]Country, error) {
	var dbCountry []Country
	res := a.DB.Find(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbCountry, ErrCountryNotFound
		}

	}
	// Get the Country Single Data
	return dbCountry, nil
}

// LOCATION //
func (a *EmployeeSet) CreateLocation(data Location) error {
	var dbLocation Location
	//res := a.DB.Where("street_address = ?", data.StreetAddress).First(&dbLocation)
	res := a.DB.FirstOrCreate(&dbLocation, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateLocation(locationId uint, data map[string]interface{}) error {
	var dbLocation Location
	res := a.DB.Model(&dbLocation).Where("id = ?", locationId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteLocation(locationId uint) error {
	var dbLocation Location
	res := a.DB.Where("id = ?", locationId).First(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrLocationNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", locationId).Delete(Location{})
	return nil
}

func (a *EmployeeSet) GetLocationId(id string) (LocationResponse, error) {
	var dbLocation Location
	var locationResponse LocationResponse
	var dbCountry Country
	if NullString(id) == nil {
		return locationResponse, ErrLocationNotFound
	}
	conv, _ := strconv.Atoi(id)
	locationId := uint(conv)

	res := a.DB.Where("id = ?", locationId).First(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return locationResponse, ErrLocationNotFound
		}
	}
	if dbLocation.CountryID > 0 {
		_ = a.DB.Where("id = ?", dbLocation.CountryID).First(&dbCountry)
	}

	locationResponse = LocationResponse{
		ID:            dbLocation.ID,
		StreetAddress: dbLocation.StreetAddress,
		PostalCode:    dbLocation.PostalCode,
		City:          dbLocation.City,
		StateProvince: dbLocation.StateProvince,
		Country:       dbCountry}
	// Get the Location Single Data
	return locationResponse, nil
}

func (a *EmployeeSet) GetLocation() ([]Location, error) {
	var dbLocation []Location
	res := a.DB.Find(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbLocation, ErrLocationNotFound
		}

	}
	// Get the Location Single Data
	return dbLocation, nil
}

// REGION //
func (a *EmployeeSet) CreateRegion(data Region) error {
	var dbRegion Region
	//res := a.DB.Where("region_name = ?", data.RegionName).First(&dbRegion)
	res := a.DB.FirstOrCreate(&dbRegion, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateRegion(regionId uint, data map[string]interface{}) error {
	var dbRegion Region
	res := a.DB.Model(&dbRegion).Where("id = ?", regionId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteRegion(regionId uint) error {
	var dbRegion Region
	res := a.DB.Where("id = ?", regionId).First(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrRegionNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", regionId).Delete(Region{})
	return nil
}

func (a *EmployeeSet) GetRegionId(id string) (Region, error) {
	var dbRegion Region
	if NullString(id) == nil {
		return dbRegion, ErrRegionNotFound
	}
	conv, _ := strconv.Atoi(id)
	regionId := uint(conv)

	res := a.DB.Where("id = ?", regionId).First(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbRegion, ErrRegionNotFound
		}

	}
	// Get the Region Single Data
	return dbRegion, nil
}

func (a *EmployeeSet) GetRegion() ([]Region, error) {
	var dbRegion []Region
	res := a.DB.Find(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbRegion, ErrRegionNotFound
		}

	}
	// Get the Region List
	return dbRegion, nil
}

// JOB //
func (a *EmployeeSet) CreateJob(data Job) error {
	var dbJob Job
	res := a.DB.Where("job_title = ?", data.JobTitle).First(&dbJob)
	res := a.DB.FirstOrCreate(&dbJob, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateJob(jobId uint, data map[string]interface{}) error {
	var dbJob Job
	res := a.DB.Model(&dbJob).Where("id = ?", jobId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteJob(jobId uint) error {
	var dbJob Job
	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrJobNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", jobId).Delete(Job{})
	return nil
}

func (a *EmployeeSet) GetJobId(id string) (Job, error) {
	var dbJob Job
	if NullString(id) == nil {
		return dbJob, ErrJobNotFound
	}
	conv, _ := strconv.Atoi(id)
	jobId := uint(conv)

	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrJobNotFound
		}

	}
	// Get the Employee Single Data
	return dbJob, nil
}

func (a *EmployeeSet) GetJob() ([]Job, error) {
	var dbJob []Job
	res := a.DB.Find(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrJobNotFound
		}

	}
	// Get the Employee Single Data
	return dbJob, nil
}

// JOB HISTORY //
func (a *EmployeeSet) CreateJobHistory(data JobHistory) error {
	var dbJobHistory JobHistory
	//res := a.DB.Where("job_id = ?", data.JobID).First(&dbJob)
	res := a.DB.FirstOrCreate(&dbJobHistory, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateJobHistory(jobId uint, data map[string]interface{}) error {
	var dbJob JobHistory
	res := a.DB.Model(&dbJob).Where("id = ?", jobId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteJobHistory(jobId uint) error {
	var dbJob JobHistory
	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrJobHistoryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", jobId).Delete(JobHistory{})
	return nil
}

func (a *EmployeeSet) GetJobHistoryId(id string) (JobHistoryResponse, error) {
	var dbJob JobHistory
	var returnJob JobHistoryResponse
	var dbJ Job
	var dbDepart Department
	if NullString(id) == nil {
		return returnJob, ErrJobHistoryNotFound
	}
	conv, _ := strconv.Atoi(id)
	jobId := uint(conv)

	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return returnJob, ErrJobHistoryNotFound
		}

	}
	if dbJob.JobID > 0 {
		_ = a.DB.Where("id = ?", dbJob.JobID).First(&dbJ)
	}
	if dbJob.JobID > 0 {
		_ = a.DB.Where("id = ?", dbJob.DepartmentID).First(&dbDepart)
	}
	returnJob = JobHistoryResponse{
		ID:         dbJob.ID,
		EmployeeID: dbJob.EmployeeID,
		StartDate:  dbJob.StartDate,
		EndDate:    dbJob.EndDate,
		Job:        dbJ,
		Department: dbDepart}
	// Get the Job History Single Data
	return returnJob, nil
}

func (a *EmployeeSet) GetJobHistory() ([]JobHistory, error) {
	var dbJob []JobHistory
	res := a.DB.Find(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrJobHistoryNotFound
		}

	}
	// Get the Job History Single Data
	return dbJob, nil
}

// EMPLOYEE //
func (a *EmployeeSet) CreateEmployee(data Employee) error {
	var dbEmployee Employee
	//res := a.DB.Where("first_name = ? and last_name = ?", data.FirstName, data.LastName).First(&dbEmployee)
	res := a.DB.FirstOrCreate(&dbEmployee, data)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidTransaction) {
			//a.DB.Create(&data)
			return res.Error
		}
	}
	if res.RowsAffected == 1 {
		return nil
	}
	//a.DB.Create(&data)
	return ErrDataIsExisted
}

func (a *EmployeeSet) UpdateEmployee(employeeId uint, data map[string]interface{}) error {
	var dbEmployee Employee
	res := a.DB.Model(&dbEmployee).Where("id = ?", employeeId).Updates(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (a *EmployeeSet) DeleteEmployee(employeeId uint) error {
	var dbJob JobHistory
	res := a.DB.Where("id = ?", employeeId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrEmployeeNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", employeeId).Delete(Employee{})
	return nil
}

func (a *EmployeeSet) GetEmployeeId(id string) (EmployeeResponse, error) {
	var dbEmployee Employee
	var dbDP Department
	var dbJob Job
	var returnMeta EmployeeResponse
	if NullString(id) == nil {
		return returnMeta, ErrParameterNotFound
	}
	conv, _ := strconv.Atoi(id)
	employeeId := uint(conv)

	res := a.DB.Where("id = ?", employeeId).First(&dbEmployee)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return returnMeta, ErrEmployeeNotFound
		}
	}
	if dbEmployee.DepartmentID > 0 {
		_ = a.DB.Where("id = ?", dbEmployee.DepartmentID).First(&dbDP)
	}
	if dbEmployee.JobID > 0 {
		_ = a.DB.Where("id = ?", dbEmployee.JobID).First(&dbJob)
	}
	returnMeta = EmployeeResponse{
		ID:          dbEmployee.ID,
		CustomID:    dbEmployee.CustomID,
		EmployeeID:  dbEmployee.EmployeeID,
		FirstName:   dbEmployee.FirstName,
		LastName:    dbEmployee.LastName,
		Email:       dbEmployee.Email,
		PhoneNumber: dbEmployee.PhoneNumber,
		JobID:       dbEmployee.JobID,
		Salary:      dbEmployee.Salary,
		Picture:     dbEmployee.Picture,
		ManagerID:   dbEmployee.ManagerID,
		Department:  dbDP,
		Job:         dbJob}

	// Get the Employee Single Data
	return returnMeta, nil
}

func (a *EmployeeSet) GetEmployee() ([]Employee, error) {
	var dbEmployee []Employee
	res := a.DB.Find(&dbEmployee)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbEmployee, ErrEmployeeNotFound
		}

	}
	// Get the Employee List
	return dbEmployee, nil
}

func NullString(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}
