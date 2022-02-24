package http

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	params "github.com/anikkatiyar99/forecast/pkg/http/params"
	"github.com/tidwall/gjson"
)

// GetOfficeAndGrid returns the office and grid
func GetOfficeAndGrid(latitude string, longitude string) (*params.Grid, error) {
	url := "https://api.weather.gov/points/" + latitude[:len(latitude)-2] + "," + longitude[:len(longitude)-2]
	log.Println("Sending Request to :", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error", err, resp.StatusCode)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("error", ErrNoOfficeFound, resp.StatusCode)
		return nil, ErrNoOfficeFound
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	values := gjson.GetMany(string(body), "properties.gridId", "properties.gridX", "properties.gridY")
	if len(values) != 3 {
		return nil, errors.New("error in fetching values office & grid")
	}
	return &params.Grid{
		GridId: values[0].String(),
		GridX:  values[1].String(),
		GridY:  values[2].String(),
	}, nil
}

// GetForecast returns the forecast for the grid
func GetForecast(grid *params.Grid) (*params.Forecast, error) {
	var period string
	url := "https://api.weather.gov/gridpoints/" + grid.GridId + "/" + grid.GridX + "," + grid.GridY + "/forecast"
	log.Println("Sending Request to :", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("error", err, resp.StatusCode)
		return nil, errors.New("error in sending request to fetch forecast")
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(ErrNoForecastFound.Error())
		log.Println("error", err, resp.StatusCode)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if time.Now().Weekday() == time.Wednesday {
		period = "properties.periods.#(name==" + "\"Tonight\"" + ")"
	} else {
		period = "properties.periods.#(name==" + "\"Wednesday Night\"" + ")"
	}
	values := gjson.Get(string(body), period)
	if values.Exists() {
		return &params.Forecast{
			Temperature:   values.Get("temperature").String() + values.Get("temperatureUnit").String(),
			ShortForecast: values.Get("shortForecast").String(),
			LongForecast:  values.Get("detailedForecast").String(),
			WindSpeed:     values.Get("windSpeed").String(),
		}, nil
	} else {
		return nil, errors.New("error in fetching values forecast")
	}
}

type errReader int

func (errReader) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
}
