package citizen

import (
	"Reservaksin-BE/app/middlewares"
	"Reservaksin-BE/businesses"
	"Reservaksin-BE/helpers/encrypt"
	"context"
	"strings"
	"time"
)

type citizenService struct {
	citizenRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middlewares.ConfigJWT
}

func NewService(repoCitizen Repository, jwtauth *middlewares.ConfigJWT, timeout time.Duration) Service {
	return &citizenService{
		citizenRepository: repoCitizen,
		contextTimeout:    timeout,
		jwtAuth:           jwtauth,
	}
}

func (repo *citizenService) Register(ctx context.Context, citizenDomain *Domain) (Domain, error) {
	existedCitizen, err := repo.citizenRepository.GetByID(citizenDomain.Id)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}

	if existedCitizen != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	citizenDomain.Password = encrypt.HashAndSalt([]byte(citizenDomain.Password))
	result, err := repo.citizenRepository.Register(citizenDomain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (repo *citizenService) Login(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.Id)
	return token, nil

}

func (repo *citizenService) GetByID(ctx context.Context, id int) (Domain, error) {
	resp, err := repo.citizenRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
