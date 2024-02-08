package model

import "github.com/golang-jwt/jwt/v5"

type Login struct {
	NickUser string `json:"nickuser"`
	Password string `json:"password"`
}

// Claim
type Claim struct {
	NickUser string `json:"nickuser"`
	jwt.Claims
}
