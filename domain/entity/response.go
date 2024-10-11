package entity

// Response is a struct that represents the response of the API
type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
