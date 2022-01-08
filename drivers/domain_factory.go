package drivers

import (
	// adminDomain "ca-reservaksin/businesses/admin"
	// currentAddressDomain "ca-reservaksin/businesses/currentAddress"
	// healthFacilitiesDomain "ca-reservaksin/businesses/healthFacilities"
	// sessionDomain "ca-reservaksin/businesses/session"
	// vaccineDomain "ca-reservaksin/businesses/vaccine"
	// adminDB "ca-reservaksin/drivers/database/admin"
	// currentAddressDB "ca-reservaksin/drivers/database/currentAddress"
	// healthFacilitiesDB "ca-reservaksin/drivers/database/healthFacilities"
	// sessionDB "ca-reservaksin/drivers/database/session"
	// vaccineDB "ca-reservaksin/drivers/database/vaccine"
	citizenDomain "Reservaksin-BE/businesses/citizen"
	citizenDB "Reservaksin-BE/drivers/database/citizen"

	"gorm.io/gorm"
)

func NewCitizenRepository(conn *gorm.DB) citizenDomain.Repository {
	return citizenDB.NewMysqlRepository(conn)
}

// func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
// 	return adminDB.NewMysqlRepository(conn)
// }

// func NewVaccineRepository(conn *gorm.DB) vaccineDomain.Repository {
// 	return vaccineDB.NewMysqlRepository(conn)
// }

// func NewCurrentAddressRepository(conn *gorm.DB) currentAddressDomain.Repository {
// 	return currentAddressDB.NewMysqlRepository(conn)
// }

// func NewHealthFacilitiesRepository(conn *gorm.DB) healthFacilitiesDomain.Repository {
// 	return healthFacilitiesDB.NewMysqlRepository(conn)
// }

// func NewSessionRepository(conn *gorm.DB) sessionDomain.Repository {
// 	return sessionDB.NewMysqlRepository(conn)
// }
