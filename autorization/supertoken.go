package autorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saalcazar/ceadlbk-info/model"
)

// Generar Token
func GenerateSuperToken(data *model.SuperLogin) (string, error) {
	claim := model.SuperClaim{
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
func ValidateSuperToken(t string) (model.SuperClaim, error) {
	token, err := jwt.ParseWithClaims(t, &model.SuperClaim{}, VerifyFunction)
	if err != nil {
		return model.SuperClaim{}, err
	}

	if !token.Valid {
		return model.SuperClaim{}, errors.New("token no válido")
	}
	claim, ok := token.Claims.(*model.SuperClaim)
	if !ok {
		return model.SuperClaim{}, errors.New("no se pudo obtener los claims")
	}
	return *claim, nil
}

// Devuelve la inf del archivo púbico
func VerifySuperFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
