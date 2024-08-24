package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func Respond(status int, req *http.Request, w http.ResponseWriter, v any) error {
	if err := req.Context().Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("client disconnected, do not send response")
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, v)
	return nil
}
