package citizen

import (
	"Reservaksin-BE/businesses/citizen"

	"gorm.io/gorm"
)

type MysqlCitizenRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) citizen.Repository {
	return &MysqlCitizenRepository{
		Conn: conn,
	}
}

func (mysqlrepo *MysqlCitizenRepository) Register(dataCitizen *citizen.Domain) (citizen.Domain, error) {
	recCitizen := fromDomain(*dataCitizen)

	err := mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByEmail(email string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	err := mysqlrepo.Conn.First(&recCitizen, "email = ?", email).Error
	// err:=mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByNIK(nik string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	err := mysqlrepo.Conn.Where("nik = ?", nik).First(&recCitizen).Error
	// err:=mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}
