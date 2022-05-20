package controllers

import (
	"Backend/config"
	"Backend/models"
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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
	coll := client.Database("bugTrack").Collection("auth")
	// find the user with the given username
	var user models.DatabaseCreds
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

// ChangePassword handles the change of password of a user
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // if the request is not POST method, return 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	Author, err := authenticate(r) // check if the user is authenticated
	if err != nil {                // if not return 401
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body) // get the new password from the request body
	var credentials models.Credentials
	err = decoder.Decode(&credentials)
	if err != nil { // if the request body is not valid return 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if Author != credentials.Username { // if the user is not the same as the authenticated user return 401
		w.WriteHeader(http.StatusUnauthorized)
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

	cost, _ := strconv.Atoi(config.ViperEnvVariable("BCRYPT_COST")) // get the cost of the bcrypt algorithm
	// hash the password
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), cost) // hash the password
	if err != nil {                                                                       // if the hashing failed return 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the collection
	coll := client.Database("bugTrack").Collection("auth")
	// update the user password with the new one
	_, err = coll.UpdateOne(context.TODO(), bson.M{"username": credentials.Username}, bson.M{"$set": bson.M{"password": string(hasedPassword)}})
	if err != nil { // if the update failed return 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create a connection to log collection
	logColl := client.Database("bugTrack").Collection("logs")
	// create a new log
	log := models.Log{
		Type:        "Password Change",
		Author:      Author,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
		Description: "Password changed by " + credentials.Username,
		Table:       "auth",
	}
	// insert the log
	_, err = logColl.InsertOne(context.TODO(), log)
	// if there is an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.WriteHeader(http.StatusOK)                       // return 200
	output := struct {                                 // create the output
		Message string `json:"message"`
	}{
		Message: "success",
	}
	err = json.NewEncoder(w).Encode(output) // encode the output
	if err != nil {                         // if the encoding failed return 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
