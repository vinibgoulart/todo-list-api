package user

import "golang.org/x/crypto/bcrypt"

func HelperUserApiToModelMapping(userStruct *UserStruct) (*User, error) {
	passwordHash, err := HelperUserPasswordHash(userStruct.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     userStruct.Name,
		Email:    userStruct.Email,
		Password: passwordHash,
	}, nil
}

func HelperUserPasswordMatch(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func HelperUserPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
