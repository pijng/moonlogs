package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Role string

const (
	RoleMember Role = "Member"
	RoleAdmin  Role = "Admin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password_digest").NotEmpty(),
		field.String("role").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
