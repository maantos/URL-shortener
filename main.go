package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maantos/urlShortener/api"
	"github.com/maantos/urlShortener/config"
	mr "github.com/maantos/urlShortener/repository/mongodb"
	rr "github.com/maantos/urlShortener/repository/redis"
	"github.com/maantos/urlShortener/shortener"
	"github.com/spf13/viper"
)

// storage <- service -> serializer -> http

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	port := viper.GetString("port")
	repo := chooseRepo()
	log.Println(repo)
	service := shortener.NewRedirectService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on port %s...", port)
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func chooseRepo() shortener.RedirectRepository {
	log.Println(viper.GetString("URL_DB"))
	switch viper.GetString("URL_DB") {
	case "redis":
		redisURL := viper.GetString("redis.uri") //os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo

	case "mongo":
		mongoURL := viper.GetString("mongo.uri") // os.Getenv("MONGO_URL")
		mongodb := viper.GetString("mongo.name") //os.Getenv("MONGO_DB")
		//mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, 20)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
