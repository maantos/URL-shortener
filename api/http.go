package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maantos/urlShortener/shortener"
	"github.com/pkg/errors"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortener.RedirectService
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{
		redirectService: redirectService,
	}
}

func setupResponse(rw http.ResponseWriter, contentType string, body []byte, statusCode int) {
	rw.Header().Set("Content-Type", contentType)
	rw.WriteHeader(statusCode)
	_, err := rw.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect, err := h.redirectService.Find(code)

	if err != nil {
		if errors.Cause(err) == shortener.ErrorRedirectNotFound {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, redirect.URL, http.StatusMovedPermanently)

}
func (h *handler) Post(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(requestBody, redirect); err != nil {
		http.Error(rw, "error decoding request body", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.redirectService.Store(redirect)

	if err != nil {
		if errors.Cause(err) == shortener.ErrorRedirectInvalid {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(redirect)
	if err != nil {
		http.Error(rw, "error encoding response body", http.StatusInternalServerError)
	}
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(rw, contentType, responseBody, http.StatusCreated)
}
