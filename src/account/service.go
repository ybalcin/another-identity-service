package account

import (
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

type Service interface {
	//	AddNewUser adds new user
	AddNewUser(firstname string, lastname string, username string, email string, password string, birth_date time.Time,
		phone_number string, gdpr bool, address *location.Address)
}

type service struct {
	users UserRepository
}

func (s *service) AddNewUser(firstname string, lastname string, username string, email string, password string, birth_date time.Time,
	phone_number string, gdpr bool, address *location.Address) {
	new_user := CreateNewUser(NewUserId(), firstname, lastname, username, email, password, birth_date, phone_number, gdpr, address)
	s.users.InsertNewUser(new_user)
}

func NewService(userRepository UserRepository) Service {
	return &service{
		users: userRepository,
	}
}
