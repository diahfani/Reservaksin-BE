package response

import (
	"Reservaksin-BE/businesses/citizen"
	"time"
)

type CitizenResponse struct {
	Id                 string `json:"id"`
	Email              string `json:"email"`
	NoHp               string `json:"nohp"`
	Username           string `json:"username"`
	NoKK               uint64 `json:"nokk"`
	Nik                string `json:"nik"`
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
	CreatedAt          time.Time
	// UpdatedAt          time.Time
	Token string `json:"token"`
}

func FromDomain(domain citizen.Domain) *CitizenResponse {
	return &CitizenResponse{
		Id:                 domain.Id,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		Username:           domain.Username,
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
		CreatedAt:          domain.CreatedAt,
	}
}
