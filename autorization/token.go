package autorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saalcazar/ceadlbk-info/model"
)

// Generar Token
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		NickUser: data.NickUser,
		Claims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    "CEADL",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Validate TOKEN
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, VerifyFunction)
	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("token no válido")
	}
	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("no se pudo obtener los claims")
	}
	return *claim, nil
}

// Devuelve la inf del archivo púbico
func VerifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
