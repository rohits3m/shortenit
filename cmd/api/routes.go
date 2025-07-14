package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (app *Application) FailureResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"message": message,
	})
}

func (app *Application) SuccessResponse(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileserver)

	mux.HandleFunc("GET /{linkId}", func(w http.ResponseWriter, r *http.Request) {
		linkId := r.PathValue("linkId")

		link, err := app.links.GetByLinkId(linkId)
		if err != nil {
			w.Write([]byte("<h1>Page not found</h1>"))
			return
		}

		http.Redirect(w, r, link.OriginalUrl, http.StatusPermanentRedirect)
	})

	mux.HandleFunc("POST /v1/link/create", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			app.FailureResponse(w, err.Error())
			return
		}

		originalUrl, exists := data["original_url"]
		if exists {
			// Checking if the given originalUrl is valid
			_, err := url.ParseRequestURI(originalUrl.(string))
			if err != nil {
				app.FailureResponse(w, "The given url is not valid")
				return
			}

			linkId, err := app.links.Create(originalUrl.(string))
			if err != nil {
				app.FailureResponse(w, err.Error())
				return
			}

			app.SuccessResponse(w, fmt.Sprintf("%s/%s", app.config.baseUrl, linkId), "")
		} else {
			app.FailureResponse(w, "original_url is required")
		}
	})

	return mux
}
