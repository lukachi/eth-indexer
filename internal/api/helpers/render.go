package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func RenderErr(w http.ResponseWriter, statusCode int, respBody any) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/vnd.api+json")

	errResp := respBody

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		panic(errors.Wrap(err, "failed to encode error response"))
	}
}

func Render(w http.ResponseWriter, statusCode int, respBody any) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/vnd.api+json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		panic(errors.Wrap(err, "failed to encode response body"))
	}
}
