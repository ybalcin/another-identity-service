package domain

import p "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	Id   p.ObjectID `bson:"id"`
	Name string     `bson:"name"`
}
