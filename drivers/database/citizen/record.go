package citizen

import (
	"Reservaksin-BE/businesses/citizen"

	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Id                 string `json:"id" gorm:"Primarykey; Not Null"`
	Email              string `json:"email"`
	NoHp               string `json:"nohp"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	NoKK               uint64 `json:"nokk"`
	Nik                string `json:"nik" gorm:"unique"`
	DateofBirth        string `json:"dob"`
	FamilyRelationship string `json:"relationship"`
	Gender             string `json:"gender"`
	MarriageStatus     string `json:"status"`
	Role               string `json:"role"`
	Address            string `json:"alamat"`
	Desa               string `json:"desa"`
	Kota               string `json:"kota"`
	Kecamatan          string `json:"kecamatan"`
	Provinsi           string `json:"provinsi"`
	ImageURI           string `json:"imageURI"`
}

func (rec *Citizen) toDomain() citizen.Domain {
	return citizen.Domain{
		Id:                 rec.Id,
		Email:              rec.Email,
		NoHp:               rec.NoHp,
		Username:           rec.Username,
		Password:           rec.Password,
		NoKK:               rec.NoKK,
		Nik:                rec.Nik,
		DateofBirth:        rec.DateofBirth,
		FamilyRelationship: rec.FamilyRelationship,
		Gender:             rec.Gender,
		MarriageStatus:     rec.MarriageStatus,
		Role:               rec.Role,
		Address:            rec.Address,
		Desa:               rec.Desa,
		Kota:               rec.Kota,
		Kecamatan:          rec.Kecamatan,
		Provinsi:           rec.Provinsi,
		ImageURI:           rec.ImageURI,
	}
}

func fromDomain(domain citizen.Domain) *Citizen {
	return &Citizen{
		Id:                 domain.Id,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		Username:           domain.Username,
		Password:           domain.Password,
		NoKK:               domain.NoKK,
		Nik:                domain.Nik,
		DateofBirth:        domain.DateofBirth,
		FamilyRelationship: domain.FamilyRelationship,
		Gender:             domain.Gender,
		MarriageStatus:     domain.MarriageStatus,
		Role:               domain.Role,
		Address:            domain.Address,
		Desa:               domain.Desa,
		Kota:               domain.Kota,
		Kecamatan:          domain.Kecamatan,
		Provinsi:           domain.Provinsi,
		ImageURI:           domain.ImageURI,
	}
}
