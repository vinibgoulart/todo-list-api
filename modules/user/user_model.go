package user

type UserStorage interface {
	ListUsers() []*User
	GetUser(string) *User
	StoreUser(user User)
}

type UserStore struct {
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []*User{
	{
		ID:       "1",
		Name:     "Vinicius Goulart",
		Email:    "viblaziusgoulart@gmail.com",
		Password: "123456",
	},
}

func (u *UserStore) ListUsers() []*User {
	return users
}

func (u *UserStore) GetUser(id string) *User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}

	return nil
}

func (u *UserStore) StoreUser(user User) {
	users = append(users, &user)
}
