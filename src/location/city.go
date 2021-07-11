package location

import p "go.mongodb.org/mongo-driver/bson/primitive"

type city struct {
	Id        p.ObjectID `bson:"id"`
	CountryId p.ObjectID `bson:"country_id"`
	Name      string     `bson:"name"`
}
