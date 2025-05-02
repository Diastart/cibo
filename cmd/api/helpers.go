package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) ReadDishNameParam(request *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(request.Context())
	dishName := params.ByName("dishName")
	
	if dishName == "" {
		return "", errors.New("dish name parameter is missing")
	}
	
	return dishName, nil
}

func (app *application) ReadUserNationality(request *http.Request) (string, error) {
	nationality := request.Header.Get("X-User-Nationality")
	
	if nationality == "" {
		return "", errors.New("X-User-Nationality header is missing")
	}
	
	return nationality, nil
}

func (app *application) ReadDishFeedback(request *http.Request) (string, error) {
	var input struct {
		Feedback string `json:"feedback"`
	}
	
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)
	if err != nil {
		return "", errors.New("invalid request body")
	}
	
	if input.Feedback == "" {
		return "", errors.New("missing feedback value")
	}
	
	if input.Feedback != "like" && input.Feedback != "dislike" {
		return "", errors.New("feedback must be either 'like' or 'dislike'")
	}
	
	return input.Feedback, nil
}