package citizen

import (
	"context"
	"time"
)

type Domain struct {
	Id                 int
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
	Register(ctx context.Context, data *Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	// 	CreateBooking(ctx context.Context, nik int, nama string) (Domain, error)
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(data *Domain) (Domain, error)
	GetByID(id int) (Domain, error)
	// CreateBooking(nik int, nama string) (Domain, error)
}
