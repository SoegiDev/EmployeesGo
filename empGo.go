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
	ErrDepartmentNotFound = errors.New("cannot delete Department cause Not Found")
	ErrCountryNotFound    = errors.New("cannot delete Country cause Not Found")
	ErrLocationNotFound   = errors.New("cannot delete Location cause Not Found")
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
	res := a.DB.Where("department_name = ?", data.DepartmentName).First(&dbDepartment)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(
				&Department{
					DepartmentName: data.DepartmentName,
					ManagerID:      data.ManagerID,
					LocationID:     data.LocationID,
					CompanyID:      data.CompanyID})
			return nil
		}
	}
	return res.Error
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

func (a *EmployeeSet) GetDepartmentId(id string) (Department, error) {
	var dbDepartment Department
	if NullString(id) != nil {
		return dbDepartment, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	departmentId := uint(conv)

	res := a.DB.Where("id = ?", departmentId).First(&dbDepartment)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbDepartment, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbDepartment, nil
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
	res := a.DB.Where("country_name = ?", data.CountryName).First(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(
				&Country{
					CountryName: data.CountryName,
					RegionID:    data.RegionID})
		}
	}
	return res.Error
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

func (a *EmployeeSet) GetCountryId(id string) (Country, error) {
	var dbCountry Country
	if NullString(id) != nil {
		return dbCountry, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	countryId := uint(conv)

	res := a.DB.Where("id = ?", countryId).First(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbCountry, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbCountry, nil
}

func (a *EmployeeSet) GetCountry() ([]Country, error) {
	var dbCountry []Country
	res := a.DB.Find(&dbCountry)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbCountry, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbCountry, nil
}

// LOCATION //
func (a *EmployeeSet) CreateLocation(data Location) error {
	var dbLocation Location
	res := a.DB.Where("street_address = ?", data.StreetAddress).First(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(data)
		}
	}
	return res.Error
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
			return ErrCountryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", locationId).Delete(Location{})
	return nil
}

func (a *EmployeeSet) GetLocationId(id string) (Location, error) {
	var dbLocation Location
	if NullString(id) != nil {
		return dbLocation, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	locationId := uint(conv)

	res := a.DB.Where("id = ?", locationId).First(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbLocation, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbLocation, nil
}

func (a *EmployeeSet) GetLocation() ([]Location, error) {
	var dbLocation []Location
	res := a.DB.Find(&dbLocation)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbLocation, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbLocation, nil
}

// REGION //
func (a *EmployeeSet) CreateRegion(data Region) error {
	var dbRegion Region
	res := a.DB.Where("region_name = ?", data.RegionName).First(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(data)
		}
	}
	return res.Error
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
			return ErrCountryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", regionId).Delete(Region{})
	return nil
}

func (a *EmployeeSet) GetRegionId(id string) (Region, error) {
	var dbRegion Region
	if NullString(id) != nil {
		return dbRegion, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	regionId := uint(conv)

	res := a.DB.Where("id = ?", regionId).First(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbRegion, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbRegion, nil
}

func (a *EmployeeSet) GetRegion() ([]Region, error) {
	var dbRegion []Region
	res := a.DB.Find(&dbRegion)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbRegion, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbRegion, nil
}

// JOB //
func (a *EmployeeSet) CreateJob(data Job) error {
	var dbJob Job
	res := a.DB.Where("job_title = ?", data.JobTitle).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(data)
		}
	}
	return res.Error
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
			return ErrCountryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", jobId).Delete(Job{})
	return nil
}

func (a *EmployeeSet) GetJobId(id string) (Job, error) {
	var dbJob Job
	if NullString(id) != nil {
		return dbJob, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	jobId := uint(conv)

	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbJob, nil
}

func (a *EmployeeSet) GetJob() ([]Job, error) {
	var dbJob []Job
	res := a.DB.Find(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbJob, nil
}

// JOB HISTORY //
func (a *EmployeeSet) CreateJobHistory(data JobHistory) error {
	var dbJob JobHistory
	res := a.DB.Where("job_id = ?", data.JobID).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			a.DB.Create(data)
		}
	}
	return res.Error
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
			return ErrCountryNotFound
		}

	}
	// Delete the Country
	a.DB.Where("id = ?", jobId).Delete(JobHistory{})
	return nil
}

func (a *EmployeeSet) GetJobHistoryId(id string) (JobHistory, error) {
	var dbJob JobHistory
	if NullString(id) != nil {
		return dbJob, ErrDepartmentNotFound
	}
	conv, _ := strconv.Atoi(id)
	jobId := uint(conv)

	res := a.DB.Where("id = ?", jobId).First(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbJob, nil
}

func (a *EmployeeSet) GetJobHistory() ([]JobHistory, error) {
	var dbJob []JobHistory
	res := a.DB.Find(&dbJob)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return dbJob, ErrDepartmentNotFound
		}

	}
	// Get the Department Single Data
	return dbJob, nil
}

func NullString(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}
