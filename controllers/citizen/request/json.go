package request

import "Reservaksin-BE/businesses/citizen"

type Citizen struct {
	Email              string `json:"email"`
	NoHp               string `json:"nohp"`
	Username           string `json:"username"`
	Password           string `json:"password"`
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
}

type CitizenLoginEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CitizenLoginNIK struct {
	Nik      uint64 `json:"email"`
	Password string `json:"password"`
}

func (req *Citizen) ToDomain() *citizen.Domain {
	return &citizen.Domain{
		Email:              req.Email,
		NoHp:               req.NoHp,
		Username:           req.Username,
		Password:           req.Password,
		NoKK:               req.NoKK,
		Nik:                req.Nik,
		DateofBirth:        req.DateofBirth,
		FamilyRelationship: req.FamilyRelationship,
		Gender:             req.Gender,
		MarriageStatus:     req.MarriageStatus,
		Role:               req.Role,
		Address:            req.Address,
		Desa:               req.Desa,
		Kota:               req.Kota,
		Kecamatan:          req.Kecamatan,
		Provinsi:           req.Provinsi,
	}
}
