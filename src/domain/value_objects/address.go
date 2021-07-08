package value_objects

type Address struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	County  string `bson:"county"`
}