package citizen

import (
	"time"
)

type Domain struct {
	Id                 string
	Email              string
	NoHp               int16
	Username           string
	Password           string
	NoKK               uint64
	Nik                uint64
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
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	LoginByEmail(email, password string) (string, error)
	LoginByNIK(nik uint64, password string) (string, error)
	// Update(ctx context.Context, id int) (Domain, error)
	// Delete(ctx context.Context, id int) (Domain, error)
	// 	CreateBooking(ctx context.Context, nik int, nama string) (Domain, error)
}

type Repository interface {
	GetByEmail(email string) (Domain, error)
	Register(data *Domain) (Domain, error)
	GetByNIK(nik uint64) (Domain, error)
	// GetByID(id int) (Domain, error)
	// CreateBooking(nik int, nama string) (Domain, error)
}
