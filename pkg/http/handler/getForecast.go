package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	params "github.com/anikkatiyar99/forecast/pkg/http/params"
)

var (
	ErrNoOfficeFound   = errors.New("404")
	ErrNoForecastFound = errors.New("500")
)

// GetForecastHandler is the handler for the /forecast endpoint
func GetForecastHandler(w http.ResponseWriter, r *http.Request) {
	var jerr params.Error
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	coOrdinates := r.Header.Get("Co-Ordinates")
	if coOrdinates == "" {
		jerr.Error.Status = strconv.Itoa(http.StatusBadRequest)
		jerr.Error.Code = "400"
		jerr.Error.Description = "Invalid request, missing co-ordinates header"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jerr)
		return
	}

	// Splitting the co-ordinates from 38.2527° N, 85.7585° W to 38.2527° N and 85.7585° W
	location := strings.Split(coOrdinates, ", ")
	if len(location) != 2 {
		jerr.Error.Status = strconv.Itoa(http.StatusBadRequest)
		jerr.Error.Code = "400"
		jerr.Error.Description = "Invalid request, incorrect co-ordinates header format"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jerr)
		return
	}

	// Splitting the co-ordinates from 38.2527° N to 38.2527° and N
	latitude := strings.Split(location[0], " ")
	longitude := strings.Split(location[1], " ")
	if (len(longitude) != 2 || len(latitude) != 2) || (latitude[1] != "N" && latitude[1] != "S") || (longitude[1] != "E" && longitude[1] != "W") {
		jerr.Error.Status = strconv.Itoa(http.StatusBadRequest)
		jerr.Error.Code = "400"
		jerr.Error.Description = "Invalid request, incorrect longitude/longitude co-ordinates header values"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jerr)
		return
	}

	if latitude[0][len(latitude[0])-1] != '°' || longitude[0][len(longitude[0])-1] != '°' {
		jerr.Error.Status = strconv.Itoa(http.StatusBadRequest)
		jerr.Error.Code = "400"
		jerr.Error.Description = "Invalid request, invalid longitude/longitude minutes values in header"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jerr)
		return
	}

	if latitude[1] == "S" {
		latitude[0] = "-" + latitude[0]
	}
	if longitude[1] == "W" {
		longitude[0] = "-" + longitude[0]
	}

	gridParams, err := GetOfficeAndGrid(latitude[0], longitude[0])
	if err == ErrNoOfficeFound {
		jerr.Error.Status = strconv.Itoa(http.StatusNotFound)
		jerr.Error.Code = "404"
		jerr.Error.Description = "Invalid request, no office found for the given co-ordinates"

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(jerr)
		return
	}
	if err != nil {
		jerr.Error.Status = strconv.Itoa(http.StatusInternalServerError)
		jerr.Error.Code = "500"
		jerr.Error.Description = "Error in fetching office and grid from api"

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jerr)
		return
	}

	forecastResp, err := GetForecast(gridParams)
	if err != nil {
		jerr.Error.Status = strconv.Itoa(http.StatusInternalServerError)
		jerr.Error.Code = "500"
		jerr.Error.Description = "Error in fetching forecast from api"

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jerr)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forecastResp)
}
