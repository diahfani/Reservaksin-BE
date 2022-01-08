package citizen

import (
	"time"
)

type Domain struct {
	Id                 string
	Email              string
	NoHp               string
	Username           string
	Password           string
	NoKK               uint64
	Nik                string
	DateofBirth        string
	FamilyRelationship string
	Gender             string
	MarriageStatus     string
	Role               string
	Address            string
	Desa               string
	Kota               string
	Kecamatan          string
	Provinsi           string
	ImageURI           string
	// CurrentAddressID   string
	// CurrentAddress     currentAddress.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	LoginByEmail(email, password string) (string, error)
	LoginByNIK(nik string, password string) (string, error)
	// Update(id int) (Domain, error)
	// Delete(id int) (Domain, error)

}

type Repository interface {
	GetByEmail(email string) (Domain, error)
	Register(data *Domain) (Domain, error)
	GetByNIK(nik string) (Domain, error)
	// GetByEmail(email string) (Domain, error)
	// Register(data *Domain) (Domain, error)
	// GetByNIK(nik string) (Domain, error)
	// GetByNoKK(noKK string) ([]Domain, error)
	// Update(id string, data *Domain) (Domain, error)
	// Delete(id string) (string, error)
}
