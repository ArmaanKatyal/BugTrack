package testing

import (
	"Backend/config"
	"Backend/controllers"
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Initialize initializes the database by creating all the collections that are required by tbe application
func Initialize() {
	client := config.ClientConnection()
	db := client.Database(config.ViperEnvVariable("dbName"))
	command := bson.D{{"create", "auth"}}
	var result bson.M
	if err := db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command2 := bson.D{{"create", "users"}}
	if err := db.RunCommand(context.TODO(), command2).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command3 := bson.D{{"create", "projects"}}
	if err := db.RunCommand(context.TODO(), command3).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command4 := bson.D{{"create", "logs"}}
	if err := db.RunCommand(context.TODO(), command4).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command5 := bson.D{{"create", "company"}}
	if err := db.RunCommand(context.TODO(), command5).Decode(&result); err != nil {
		log.Fatal(err)
	}
	command6 := bson.D{{"create", "tickets"}}
	if err := db.RunCommand(context.TODO(), command6).Decode(&result); err != nil {
		log.Fatal(err)
	}
}

//func init() {
//	Initialize()
//}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/auth/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/api/v1/project", controllers.AllProjects).Methods("GET")
	router.Path("/api/v1/user").Queries("role", "{role}").HandlerFunc(controllers.AllUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/validUsername/{username:[A-Za-z][A-Za-z0-9_]{7,29}}", controllers.CheckUsernameExists).Methods("GET")
	return router
}

// authToken is used to store the token that we get from login request in TestLogin
var authToken string

// Signup if working as expected that has been already checked
// TestSignupWithAlreadyRegisteredUser checks when a user is already registered
func TestSignupWithAlreadyRegisteredUser(t *testing.T) {
	body := strings.NewReader(`{"first_name":"tester","last_name":"test","username":"tester123", "email":"test@test.com"
		,"password":"test@123", "company_name": "TESTING","company_code":"TEST"}`)
	request, _ := http.NewRequest("POST", "/api/v1/auth/signup", body)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusConflict, response.Code, "Conflict response is expected")
}

// TestLogin tests the login functionality
func TestLogin(t *testing.T) {
	body := strings.NewReader(`{"username":"tester123","password":"test@123", "company_code":"TEST"}`)
	request, _ := http.NewRequest("POST", "/api/v1/auth/login", body)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	authToken = response.Result().Cookies()[0].Value
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

// TestProject test if the function returns all the projects when the user is admin
func TestProject(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/project", nil)
	request.Header.Set("token", authToken)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

// TestProjectWithInvalidToken tests if the function returns all the projects when the user is admin but the token is invalid
func TestProjectWithInvalidToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/project", nil)
	request.Header.Set("token", "invalidToken")
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "Unauthorized response is expected")
}

// TestProjectWithEmptyToken tests if the function returns all the projects when the user is admin but the token is empty
func TestProjectWithEmptyToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/project", nil)
	request.Header.Set("token", "")
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "Unauthorized response is expected")
}

// TestUsers test if the function returns all the users when the user is admin
func TestUsers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user?role=", nil)
	request.Header.Set("token", authToken)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

// TestUsersWithInvalidToken tests if the function returns all the users when the user is admin but the token is invalid
func TestUsersWithInvalidToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user?role=", nil)
	request.Header.Set("token", "invalidToken")
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "Unauthorized response is expected")
}

// TestUserByRole test if the function returns all the users when the user is admin
func TestUserByRole(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user?role=admin", nil)
	request.Header.Set("token", authToken)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	expected := "[{\"id\":\"6289d580ac49f637051dc0c8\",\"first_name\":\"tester\",\"last_name\":\"test\",\"username\":\"tester123\",\"email\":\"test@test.com\",\"role\":\"admin\",\"created_on\":\"2022-05-22T00:17:36.198-06:00\",\"company_code\":\"TEST\",\"locked\":false}]\n"
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, expected, response.Body.String(), "expected and actual response should be same")
}

// TestUserByRoleWithInvalidToken tests if the function returns all the users when the user is admin but the token is invalid
func TestUserByRoleWithInvalidToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user?role=admin", nil)
	request.Header.Set("token", "invalidToken")
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "Unauthorized response is expected")
}

// TestValidUsername tests if the function returns Conflict when the username is Invalid
func TestValidUsername(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user/validUsername/tester123", nil)
	request.Header.Set("token", authToken)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	expected := "{\"exists\":true}\n"
	assert.Equal(t, http.StatusConflict, response.Code, "Status Conflict response is expected")
	assert.Equal(t, expected, response.Body.String(), "expected and actual response should be same")
}

// TestValidUsername2 tests if the function returns StatusOK when the username is Valid
func TestValidUsername2(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/user/validUsername/tester1234", nil)
	request.Header.Set("token", authToken)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	expected := "{\"exists\":false}\n"
	assert.Equal(t, http.StatusOK, response.Code, "Status OK response is expected")
	assert.Equal(t, expected, response.Body.String(), "expected and actual response should be same")
}
