package models

import "github.com/golang-jwt/jwt/v4"

// Jwks struct holds the JSON Web Key Set
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys represents a public key.
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// Auth holds the JWT token
type Auth struct {
	Token string `json:"token"`
}

// Claims is the struct that contains the claims of the jwt token
type Claims struct {
	Username    string `json:"username"`
	CompanyCode string `json:"company_code"`
	Role        string `json:"role"`
	jwt.RegisteredClaims
}

type Signup struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	CompanyCode string `json:"company_code"`
}
