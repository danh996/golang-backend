package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/danh996/go-school/data/models"
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

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
