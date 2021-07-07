package value_objects

import "go.mongodb.org/mongo-driver/bson/primitive"

type UniqueId struct {
	Value interface{}
}

//	SetNewId sets new object id as value
func (u *UniqueId) SetNewId() {
	u.Value = primitive.NewObjectID()
}
