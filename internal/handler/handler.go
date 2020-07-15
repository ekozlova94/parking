package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ekozlova94/parking/internal/parking"
	"github.com/ekozlova94/parking/pkg/forms"
)

func Subscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "request not post", http.StatusMethodNotAllowed)
		return
	}
	req, err := parseJson(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := parking.Subscription(r.Context(), req); err != nil {
		if err == parking.ErrConflict {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	createResponse(w, forms.NewResponse(true, false))
}

func Check(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "request not post", http.StatusMethodNotAllowed)
		return
	}
	req, err := parseJson(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := parking.Check(r.Context(), req); err != nil {
		if err == parking.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	createResponse(w, forms.NewResponse(false, true))
}

func parseJson(r *http.Request) (*forms.Request, error) {
	var req forms.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func createResponse(w http.ResponseWriter, result *forms.Response) {
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
