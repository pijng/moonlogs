package entities

type User struct {
	ID             int    `json:"id" sql:"id"`
	Name           string `json:"name" sql:"name"`
	Email          string `json:"email" sql:"email"`
	Password       string `json:"password" sql:"password"`
	PasswordDigest string `json:"password_digest" sql:"password_digest"`
	Role           Role   `json:"role" sql:"role"`
	Tags           Tags   `json:"tag_ids" sql:"tag_ids"`
	Token          string `json:"token" sql:"token"`
}

type Role string

const (
	MemberRole Role = "Member"
	AdminRole  Role = "Admin"
	TokenRole  Role = "TokenRole"
)
