package main

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}
