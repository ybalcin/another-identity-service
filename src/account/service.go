package account

import (
	"time"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/location"
)

const (
	err_insert_new_user string = "An error occured while adding new user!"
	err_user_not_found  string = "User not found!"
)

type (
	errService struct {
		FriendlyMessage string
		InnerException  error
	}

	IAccountService interface {
		//	AddNewUser adds new user
		AddNewUser(firstname string, lastname string, username string, email string, password string, birthDate time.Time,
			phoneNumber string, gdpr bool, address *location.Address) *errService
		//	AddRoleToUser adds role to user
		AddRoleToUser(roleName string, userId string) *common.FriendlyError
		//	GetUserList gets users
		GetUserList() ([]*user, *common.FriendlyError)
	}
	accountService struct {
		users IUserRepository
	}
)

func (s *accountService) AddNewUser(firstname string, lastname string, username string, email string, password string, birthDate time.Time,
	phoneNumber string, gdpr bool, address *location.Address) *errService {

	newUser, _ := NewUser(NewUserId(), firstname, lastname, username, email, password, birthDate, phoneNumber, gdpr, address)

	if errRepository := s.users.InsertNewUser(newUser); errRepository != nil {
		return &errService{
			FriendlyMessage: err_insert_new_user,
			InnerException:  errRepository.InnerException,
		}
	}

	return nil
}

func (s *accountService) AddRoleToUser(roleName string, userId string) *common.FriendlyError {
	user, err := s.users.GetUserById(UserIdFromHex(userId))
	if err != nil {
		return err
	}

	if user == nil {
		return &common.FriendlyError{
			Message: err_user_not_found,
		}
	}

	if err = user.AddRole(roleName); err != nil {
		return err
	}

	return s.users.UpdateOneByFields(user, []string{"Roles"})
}

func (s *accountService) GetUserList() ([]*user, *common.FriendlyError) {
	users, err := s.users.GetUserList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// NewAccountService accountService initializing constructor
func NewAccountService(userRepository IUserRepository) IAccountService {
	return &accountService{
		users: userRepository,
	}
}
