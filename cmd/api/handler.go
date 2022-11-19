package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/danh996/go-school/data/models"
	"github.com/golang-jwt/jwt/v4"
)

func (app *Config) GetListSchools(w http.ResponseWriter, r *http.Request) {

	// validate the user against the database
	schools, err := app.Models.School.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Get list school success"),
		Data:    schools,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) InsertSchool(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	insertSchool := models.School{
		Name: requestPayload.Name,
	}

	schoolId, err := app.Models.School.Insert(insertSchool)
	if err != nil {
		app.errorJSON(w, errors.New("Error when insert school"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Insert school success"),
		Data:    schoolId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateSchool(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	updateSchool := models.School{
		Name: requestPayload.Name,
		ID:   requestPayload.Id,
	}

	// validate the user against the database
	err = updateSchool.Update()
	if err != nil {
		app.errorJSON(w, errors.New("Error when update school"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Updated school info success"),
		Data:    "",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Id int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.School.DeleteByID(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, errors.New("Error when delete school"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Delete school success"),
		Data:    "",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New(fmt.Sprintf("Can not find user with provided email %s", err)), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	generatedToken := GenerateToken(strconv.FormatUint((uint64(user.ID)), 10))
	fmt.Println("1111111111111111111111111", generatedToken)

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    generatedToken,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func GenerateToken(UserID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	var sampleSecretKey = []byte("UserID")

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		panic(err)
	}
	return tokenString
}

func (app *Config) CreateUser(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Active   int    `json:"active,omitempty"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	insertUser := models.User{
		Name:     requestPayload.Name,
		Password: requestPayload.Password,
		Email:    requestPayload.Email,
		Active:   requestPayload.Active,
	}

	schoolId, err := app.Models.User.Insert(insertUser)
	if err != nil {
		app.errorJSON(w, errors.New(fmt.Sprintf("Error when insert user %s", err)), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Insert user success"),
		Data:    schoolId,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
