package user

import "github.com/vinibgoulart/todo-list/db"

type UserStorage interface {
	Get(int64) (*User, error)
	Insert(User) (int64, error)
}

type UserStore struct {
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserStore) Get(id int64) (*User, error) {
	conn, err := db.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

	user := &User{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}

func (u *UserStore) Insert(user User) (int64, error) {
	conn, err := db.ConnectDatabase()

	if err != nil {
		return -1, err
	}

	defer conn.Close()

	sql := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err = conn.QueryRow(sql, user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, err
}
