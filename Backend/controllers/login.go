package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// UserLogin handles the login of a user
func UserLogin(w http.ResponseWriter, r *http.Request) {
	// if the request is not POST method, return 405
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get the username and password from the request body
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create the client connection
	client := config.ClientConnection()
	defer func() {
		// close the client connection
		if err := client.Disconnect(context.TODO()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	// get the collection
	coll := client.Database("bugTrack").Collection("users")
	// find the user with the given username
	var user models.User
	err = coll.FindOne(context.TODO(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert the password to bytes
	databasePassword := []byte(user.Password)
	userPassword := []byte(credentials.Password)

	// compare the password
	err = bcrypt.CompareHashAndPassword(databasePassword, userPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create the expiration time
	expirationTime := time.Now().Add(1 * time.Hour)
	// create the claims that contain the information carried by the token
	claims := &models.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// encrypt the claims with the secret key and HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenKEY := []byte(config.ViperEnvVariable("JWT_KEY"))
	tokenString, err := token.SignedString(tokenKEY)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set the cookie with the token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

// authenticate checks if the jwt token/User is valid or not
func authenticate(r *http.Request) (string, error) {
	// get the token from the header of the request
	authToken := r.Header.Get("token")
	// create the claims
	claims := &models.Claims{}

	// parse the token
	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ViperEnvVariable("JWT_KEY")), nil
	})
	if err != nil {
		// if the signature is invalid return an error
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	// if the token is invalid return an error
	if !token.Valid {
		return "", err
	}
	// if the token is valid return the username
	return claims.Username, nil
}

// UserLogout handles the logout of a user
func UserLogout(w http.ResponseWriter, r *http.Request) {
	// if the request is not GET method, return 405
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// set the cookie with the token to blank
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}
