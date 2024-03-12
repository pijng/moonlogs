package entities

type User struct {
	ID             int       `json:"id" sql:"id" bson:"id"`
	Name           string    `json:"name" sql:"name" bson:"name"`
	Email          string    `json:"email" sql:"email" bson:"email"`
	Password       string    `json:"password" sql:"password" bson:"password"`
	PasswordDigest string    `json:"password_digest" sql:"password_digest" bson:"password_digest"`
	Role           Role      `json:"role" sql:"role" bson:"role"`
	Tags           Tags      `json:"tag_ids" sql:"tag_ids" bson:"tag_ids"`
	Token          string    `json:"token" sql:"token" bson:"token"`
	IsRevoked      BoolAsInt `json:"is_revoked" sql:"is_revoked" bson:"is_revoked"`
}

type Role string

const (
	MemberRole Role = "Member"
	AdminRole  Role = "Admin"
	TokenRole  Role = "TokenRole"
)
