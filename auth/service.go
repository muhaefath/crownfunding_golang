package auth

import (
	"github.com/dgrijalva/jwt-go"
	// "github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(user_id int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("GOLANG_CROWDFUNDING_SECRET_KEY")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user_id int) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = user_id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// err := godotenv.Load(".env")
	// os.Getenv("SECRET_KEY")

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
