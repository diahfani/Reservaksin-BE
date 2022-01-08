package citizen

import (
	"Reservaksin-BE/app/middlewares"
	"Reservaksin-BE/businesses"
	"Reservaksin-BE/helpers/encrypt"
	"strings"
)

type citizenService struct {
	citizenRepository Repository
	// contextTimeout    time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewService(repoCitizen Repository, jwtauth *middlewares.ConfigJWT) Service {
	return &citizenService{
		citizenRepository: repoCitizen,
		// contextTimeout:    timeout,
		jwtAuth: jwtauth,
	}
}

func (repo *citizenService) Register(citizenDomain *Domain) (Domain, error) {
	existedCitizen, err := repo.citizenRepository.GetByNIK(uint64(citizenDomain.Nik))
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

func (repo *citizenService) LoginByEmail(email, password string) (string, error) {
	// ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout)
	// defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByEmail(email)
	// citizenNIK, err := repo.citizenRepository.GetByNIK(ctx, nik)

	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.Id)
	return token, nil

}

func (repo *citizenService) LoginByNIK(nik uint64, password string) (string, error) {
	// ctx, cancel := context.WithTimeout(ctx, repo.contextTimeout)
	// defer cancel()

	if nik == 0 && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByNIK(uint64(nik))
	// citizenNIK, err := repo.citizenRepository.GetByNIK(ctx, nik)

	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.Id)
	return token, nil

}

// func (repo *citizenService) GetByNIK(ctx context.Context, id int) (Domain, error) {
// 	resp, err := repo.citizenRepository.GetByNIK(id)
// 	if err != nil {
// 		return Domain{}, err
// 	}
// 	return resp, nil
// }
