package domain

import (
	vo "github.com/ybalcin/another-identity-service/domain/valueObjects"
	p "go.mongodb.org/mongo-driver/bson/primitive"
	t "time"
)

type User struct {
	Id                   p.ObjectID   `bson:"id"`
	FirstName            string       `bson:"first_name"`
	LastName             string       `bson:"last_name"`
	UserName             string       `bson:"user_name"`
	Email                string       `bson:"email"`
	EmailConfirmed       bool         `bson:"email_confirmed"`
	NormalizedEmail      string       `bson:"normalized_email"`
	PasswordHash         string       `bson:"password_hash"`
	BirthDate            t.Time       `bson:"birth_date"`
	PhoneNumber          string       `bson:"phone_number"`
	PhoneNumberConfirmed bool         `bson:"phone_number_confirmed"`
	LastLoginDate        t.Time       `bson:"last_login_date"`
	Gdpr                 bool         `bson:"gdpr"`
	Roles                []Role       `bson:"roles"`
	Addresses            []vo.Address `bson:"addresses"`
	CreatedDate          t.Time       `bson:"created_date"`
	UpdatedDate          t.Time       `bson:"updated_date"`
}
