package user

func UserPasswordMatch(password string, passwordHash string) bool {
	return password == passwordHash
}
