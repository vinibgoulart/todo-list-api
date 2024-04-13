package user

func HelperUserPasswordMatch(password string, passwordHash string) bool {
	return password == passwordHash
}
