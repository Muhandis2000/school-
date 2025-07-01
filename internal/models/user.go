// internal/models/user.go
package models

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Verified bool   `db:"verified"`
}
