package account

import (
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

const (
	insertNewUserError string = "An error occured while adding new user!"
)

type (
	errService struct {
		FriendlyMessage string
		InnerException  error
	}
)

type Service interface {
	//	AddNewUser adds new user
	AddNewUser(firstname string, lastname string, username string, email string, password string, birth_date time.Time,
		phone_number string, gdpr bool, address *location.Address) *errService
}

type service struct {
	users UserRepository
}

func (s *service) AddNewUser(firstname string, lastname string, username string, email string, password string, birth_date time.Time,
	phone_number string, gdpr bool, address *location.Address) *errService {

	new_user, _ := CreateNewUser(NewUserId(), firstname, lastname, username, email, password, birth_date, phone_number, gdpr, address)

	if errRepository := s.users.InsertNewUser(new_user); errRepository != nil {
		return &errService{
			FriendlyMessage: insertNewUserError,
			InnerException:  errRepository.InnerException,
		}
	}

	return nil
}

func NewService(userRepository UserRepository) Service {
	return &service{
		users: userRepository,
	}
}
