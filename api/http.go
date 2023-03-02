package api

import (
	"log"
	"net/http"

	js "github.com/maantos/urlShortener/serializer/json"
	ms "github.com/maantos/urlShortener/serializer/msgpck"
	"github.com/maantos/urlShortener/shortener"
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

func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

func (h *handler) Get(http.ResponseWriter, *http.Request) {

}
func (h *handler) Post(http.ResponseWriter, *http.Request) {

}
