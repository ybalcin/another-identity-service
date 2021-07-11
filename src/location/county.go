package location

import p "go.mongodb.org/mongo-driver/bson/primitive"

type county struct {
	Id        p.ObjectID `bson:"id"`
	CountryId p.ObjectID `bson:"country_id"`
	CityId    p.ObjectID `bson:"city_id"`
	Name      string     `bson:"name"`
}
