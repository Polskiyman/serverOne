package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func readBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func writeResponse(w http.ResponseWriter, statusCode int, resp []byte) {
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}

func response500(w http.ResponseWriter, err error) {
	writeResponse(w, http.StatusInternalServerError, []byte(err.Error()))
}

func response400(w http.ResponseWriter, err error) {
	writeResponse(w, http.StatusBadRequest, []byte(err.Error()))
}

func response(w http.ResponseWriter, status int, resp interface{}) {
	data, err := json.Marshal(resp)
	if err != nil {
		response500(w, err)
		return
	}
	writeResponse(w, status, data)
}
