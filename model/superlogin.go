package model

import "github.com/golang-jwt/jwt/v5"

type SuperLogin struct {
	NickUser string `json:"nick"`
	Password string `json:"password"`
}

// Claim
type SuperClaim struct {
	NickUser string `json:"nickuser"`
	jwt.Claims
}
