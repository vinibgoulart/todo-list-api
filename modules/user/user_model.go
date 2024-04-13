package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserStorage interface {
	Get(*sql.DB, string) (*User, error)
	GetByEmail(*sql.DB, string) (*User, error)
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
	sql := `SELECT * FROM users WHERE id=$1`

	row := db.QueryRow(sql, id)

	user := &User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}

func (u *UserStore) GetByEmail(db *sql.DB, email string) (*User, error) {
	sql := `SELECT * FROM users WHERE email=$1`

	row := db.QueryRow(sql, email)

	user := &User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}

func (u *UserStore) Insert(db *sql.DB, user User) (*User, error) {
	sql := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *`

	row := db.QueryRow(sql, user.Name, user.Email, user.Password)

	userCreated := &User{}

	err := row.Scan(&userCreated.ID, &userCreated.Name, &userCreated.Email, &userCreated.Password)

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
