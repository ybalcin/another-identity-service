package account

import (
	"fmt"
	"reflect"
	t "time"

	"github.com/ybalcin/another-identity-service/common"
	l "github.com/ybalcin/another-identity-service/location"
	"github.com/ybalcin/another-identity-service/utils"
	p "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

const (
	//	user colleciton name
	user_collection string = "users"

	ERR_EMPTY_ROLE_NAME = "Role name is empty!"
)

//	Unique identifier for user
type userId p.ObjectID

type user struct {
	UserId               p.ObjectID  `bson:"_id"`
	Firstname            string      `bson:"firstname" validate:"ne=''"`
	Lastname             string      `bson:"lastname" validate:"ne=''"`
	Username             string      `bson:"username" validate:"ne=''"`
	NormalizedUsername   string      `bson:"normalized_username"`
	Email                string      `bson:"email" validate:"email"`
	NormalizedEmail      string      `bson:"normalized_email"`
	EmailConfirmed       bool        `bson:"email_confirmed"`
	PasswordHash         string      `bson:"password_hash"`
	BirthDate            t.Time      `bson:"birthdate"`
	PhoneNumber          string      `bson:"phonenumber" validate:"ne=''"`
	PhoneNumberConfirmed bool        `bson:"phonenumber_confirmed"`
	LastLoginDate        t.Time      `bson:"last_login_date"`
	Gdpr                 bool        `bson:"gdpr"`
	Roles                []string    `bson:"roles"`
	Addresses            []l.Address `bson:"addresses"`
	CreatedDate          t.Time      `bson:"created_date"`
	UpdatedDate          t.Time      `bson:"updated_date"`
}

// NewUser creates new user
func CreateNewUser(id userId, firstname string, lastname string, username string, email string, passsword string, birth_date t.Time,
	phone_number string, gdpr bool, address *l.Address) (*user, []*common.ValidationError) {

	addresses := []l.Address{
		*address,
	}

	user := user{
		UserId:               p.ObjectID(id),
		Firstname:            firstname,
		Lastname:             lastname,
		Username:             username,
		NormalizedUsername:   utils.Normalize(username),
		EmailConfirmed:       false,
		PasswordHash:         utils.HashPassword(passsword),
		BirthDate:            birth_date,
		PhoneNumber:          phone_number,
		PhoneNumberConfirmed: false,
		Email:                email,
		NormalizedEmail:      utils.Normalize(email),
		Gdpr:                 gdpr,
		Addresses:            addresses,
		Roles:                []string{},
		LastLoginDate:        t.Now().UTC(),
		CreatedDate:          t.Now().UTC(),
		UpdatedDate:          t.Now().UTC(),
	}

	//	validate
	if validation_errors := validate(&user); validation_errors != nil {
		return nil, validation_errors
	}

	return &user, nil
}

func (u *user) Equals(other *user) bool {
	if other == nil {
		return false
	}

	if u == other {
		return true
	}

	return u.UserId == other.UserId
}

func (u *user) AddRole(roleName string) *common.FriendlyError {
	if roleName == "" {
		return &common.FriendlyError{
			Message: ERR_EMPTY_ROLE_NAME,
		}
	}

	u.Roles = append(u.Roles, roleName)
	return nil
}

func (u *user) GetFieldValue(field string) interface{} {
	e := reflect.ValueOf(u).Elem()

	reflectedField := e.FieldByName(field)
	return reflectedField.Interface()
}

// NewUserId generates uniqueue user Id
func NewUserId() userId {
	return userId(p.NewObjectID())
}

func UserIdFromHex(id string) userId {
	objId, _ := p.ObjectIDFromHex(id)

	return userId(objId)
}

func validate(u *user) []*common.ValidationError {
	validate := validator.New()
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	validaton_errors := []*common.ValidationError{}
	for _, e := range err.(validator.ValidationErrors) {
		validaton_errors = append(validaton_errors, &common.ValidationError{
			Message: fmt.Sprintf("Error for: %s", e.Field()),
			Field:   e.Field(),
		})
	}

	return validaton_errors
}
