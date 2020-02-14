package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoint Endpoints) http.Handler {

	m := mux.NewRouter()

	m.Methods("POST").Path("/dostuff/").Handler(httptransport.NewServer(
		endpoint.DoStuffEndpoint,
		decodeHTTPDoStuffRequest,
		encodeHTTPGenericResponse,
	))

	return m
}

func decodeHTTPDoStuffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req = DoStuffRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
