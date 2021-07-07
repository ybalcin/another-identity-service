package domain

import p "go.mongodb.org/mongo-driver/bson/primitive"

type role struct {
	Id   p.ObjectID `bson:"_id"`
	Name string     `bson:"name"`
}
