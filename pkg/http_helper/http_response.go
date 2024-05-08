package http_helper

import (
	"encoding/json"
	"net/http"
)

type simpleErr struct {
	Err string `json:"error"`
}

const (
	headerContentType = "Content-Type"
	jsonMIME          = "application/json;charset=UTF-8"
)

func JSON(w http.ResponseWriter, code int, i interface{}) {
	w.Header().Set(headerContentType, jsonMIME)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	_ = enc.Encode(i)
}

func JSONErr(w http.ResponseWriter, code int, err string) {
	w.Header().Set(headerContentType, jsonMIME)
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	_ = enc.Encode(simpleErr{Err: err})
}
