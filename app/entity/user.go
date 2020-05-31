/*
 *
 */

package entity

type User struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
	Name  string `db:"name"`
}

func NewUser(email, name string) *User {
	return &User{
		Email: email,
		Name:  name,
	}
}
