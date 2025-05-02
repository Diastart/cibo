package main

import (
	"net/http"
)

func (app *application) dishdetailsHandler(response http.ResponseWriter, request *http.Request) {
	dishName, err := app.ReadDishNameParam(request)
	if err != nil {
		app.notFoundResponse(response)
		return
	}
	
	nationality, err := app.ReadUserNationality(request)
	if err != nil {
		app.notFoundResponse(response)
		return
	}
	
}

func (app *application) dishfeedbackHandler(response http.ResponseWriter, request *http.Request) {
	// Use the helper functions to parse and validate inputs
	dishName, err := app.ReadDishNameParam(request)
	if err != nil {
		app.notFoundResponse(response)
		return
	}
	
	nationality, err := app.ReadUserNationality(request)
	if err != nil {
		app.notFoundResponse(response)
		return
	}
	
	feedback, err := app.ReadDishFeedback(request)
	if err != nil {
		app.badRequestResponse(response, err.Error())
		return
	}
	
}