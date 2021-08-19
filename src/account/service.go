package account

import (
	"time"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/location"
)

const (
	insertNewUserError string = "An error occured while adding new user!"
	userNotFound       string = "User not found!"
)

type (
	errService struct {
		FriendlyMessage string
		InnerException  error
	}

	Service interface {
		//	AddNewUser adds new user
		AddNewUser(firstname string, lastname string, username string, email string, password string, birth_date time.Time,
			phone_number string, gdpr bool, address *location.Address) *errService

		AddRole(roleName string, userId string) *common.FriendlyError
	}
	service struct {
		users UserRepository
	}
)

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

func (s *service) AddRole(roleName string, userId string) *common.FriendlyError {
	user, err := s.users.GetUserById(UserIdFromHex(userId))
	if err != nil {
		return err
	}

	if user == nil {
		return &common.FriendlyError{
			Message: userNotFound,
		}
	}

	if err = user.AddRole(roleName); err != nil {
		return err
	}

	return s.users.UpdateOneByFields(user, []string{"Roles"})
}

func NewService(userRepository UserRepository) Service {
	return &service{
		users: userRepository,
	}
}
