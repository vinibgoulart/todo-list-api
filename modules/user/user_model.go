package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserStorage interface {
	Get(*sql.DB, string) (*User, error)
	Insert(*sql.DB, User) (*User, error)
}

type UserStore struct {
}

type User struct {
	ID       uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *UserStore) Get(db *sql.DB, id string) (*User, error) {
	row := db.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

	user := &User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}

func (u *UserStore) Insert(db *sql.DB, user User) (*User, error) {
	sql := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email`

	userCreated := &User{}

	err := db.QueryRow(sql, user.Name, user.Email, user.Password).Scan(&userCreated.ID, &userCreated.Name, &userCreated.Email)

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
