package authorization

import (
	p "go.mongodb.org/mongo-driver/bson/primitive"
)

// role collection name
const role_collection string = "roles"

type roleId p.ObjectID

type role struct {
	RoleId p.ObjectID `bson:"_id"`
	Name   string     `bson:"name"`
}

func NewRole(roleId roleId, name string) *role {
	return &role{
		RoleId: p.ObjectID(roleId),
		Name:   name,
	}
}

func (u *role) Equals(other *role) bool {
	if other == nil {
		return false
	}

	if u == other {
		return true
	}

	return u.RoleId == other.RoleId
}

func NewRoleId() roleId {
	return roleId(p.NewObjectID())
}
