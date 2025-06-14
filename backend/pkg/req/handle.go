package req

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/vnkot/piklnk/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}

func HandleQueryParams[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var result T

	decoder := schema.NewDecoder()
	err := decoder.Decode(&result, r.URL.Query())

	if err != nil {
		return nil, err
	}

	return &result, nil
}
