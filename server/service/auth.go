package service

type AuthService struct{}

func (auth *AuthService) Login(username string, password string) (string, error) {
	token, err := auth.createToken(username)
	return token, err
}

func (auth *AuthService) JwtValid(jwt string) (err error) {
	return nil
}

func (auth *AuthService) createToken(username string) (string, error) {
	return username, nil
}
