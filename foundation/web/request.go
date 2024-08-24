package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Decode unmarshals the request body to given struct
func Decode(r *http.Request, v any) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("request: unable to read payload: %w", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("request: decode: %w", err)
	}
	return nil
}

// Param returns the web call parameters from the request.
func Param(r *http.Request, key string) string {
	vars := mux.Vars(r)
	val, ok := vars[key]
	if ok {
		return val
	}
	return ""
}

// Query returns the query parameter of gien key from request
func QueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
