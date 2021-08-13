package authorization

import p "go.mongodb.org/mongo-driver/bson/primitive"

// role collection name
const role_collection string = "roles"

type role struct {
	Id   p.ObjectID `bson:"_id"`
	Name string     `bson:"name"`
}
