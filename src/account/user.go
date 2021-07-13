package account

import (
	"strings"
	t "time"

	l "github.com/ybalcin/another-identity-service/location"
	"github.com/ybalcin/another-identity-service/utils"
	p "go.mongodb.org/mongo-driver/bson/primitive"
)

//	Unique identifier for user
type UserId p.ObjectID

type user struct {
	UserId               UserId      `bson:"_id"`
	Firstname            string      `bson:"first_name"`
	Lastname             string      `bson:"last_name"`
	Username             string      `bson:"user_name"`
	NormalizedUsername   string      `bson:"normalized_user_name"`
	Email                string      `bson:"email"`
	NormalizedEmail      string      `bson:"normalized_email"`
	EmailConfirmed       bool        `bson:"email_confirmed"`
	PasswordHash         string      `bson:"password_hash"`
	BirthDate            t.Time      `bson:"birth_date"`
	PhoneNumber          string      `bson:"phone_number"`
	PhoneNumberConfirmed bool        `bson:"phone_number_confirmed"`
	LastLoginDate        t.Time      `bson:"last_login_date"`
	Gdpr                 bool        `bson:"gdpr"`
	Roles                []string    `bson:"roles"`
	Addresses            []l.Address `bson:"addresses"`
	CreatedDate          t.Time      `bson:"created_date"`
	UpdatedDate          t.Time      `bson:"updated_date"`
}

// NewUser creates new user
func NewUser(id UserId, firstname string, lastname string, username string, email string, password_hash string, birth_date t.Time,
	phone_number string, gdpr bool, address *l.Address) *user {

	addresses := []l.Address{
		*address,
	}

	return &user{
		UserId:               id,
		Firstname:            firstname,
		Lastname:             lastname,
		Username:             username,
		NormalizedUsername:   strings.ToUpper(*utils.RemoveDiacritics(&username)),
		EmailConfirmed:       false,
		PasswordHash:         password_hash,
		BirthDate:            birth_date,
		PhoneNumber:          phone_number,
		PhoneNumberConfirmed: false,
		Email:                email,
		NormalizedEmail:      strings.ToUpper(*utils.RemoveDiacritics(&email)),
		Gdpr:                 gdpr,
		Addresses:            addresses,
		Roles:                []string{},
		LastLoginDate:        t.Now().UTC(),
		CreatedDate:          t.Now().UTC(),
		UpdatedDate:          t.Now().UTC(),
	}
}

// NewUserId generates uniqueue user Id
func NewUserId() UserId {
	return UserId(p.NewObjectID())
}
