package main

import "net/http"

// healthcheckHandler
//
//	@Summary		Health Check
//	@Description	Returns the health status of the API
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"available"}`))
}
