package api

import (
	"net/http"
)

func HandleBody[T any](writer *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		Json(*writer, err.Error(), 402)
		return &body, err
	}
	err = IsValid(body)
	if err != nil {
		Json(*writer, err.Error(), 402)
		return &body, err
	}
	return &body, nil
}
