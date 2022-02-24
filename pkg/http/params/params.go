package http

// Error is object for application errors response
type Error struct {
	Error ErrorDetails `json:"error"`
}

// ErrorDetails is object for application errors details
type ErrorDetails struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Status      string `json:"httpStatus"`
}

// Grid is a struct for the office coordinates
type Grid struct {
	GridId string `json:"gridId"`
	GridX  string `json:"gridX"`
	GridY  string `json:"gridY"`
}

// Forecast is a struct for the forecast
type Forecast struct {
	Temperature   string `json:"temperature"`
	ShortForecast string `json:"-"`
	LongForecast  string `json:"-"`
	WindSpeed     string `json:"-"`
}
