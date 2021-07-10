package location

// Address acts as value object
type Address struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	County  string `bson:"county"`
}
