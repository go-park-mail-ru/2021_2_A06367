package main

import (
	"fmt"
	authDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/http"
	filmsDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	r := mux.NewRouter()
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8080")}

	authHandler := authDelivery.AuthHandler{}
	filmsHandler := filmsDelivery.FilmsHandler{}

	auth := r.PathPrefix("/user").Subrouter()
	{
		auth.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
		auth.HandleFunc("/logout", authHandler.Logout).Methods(http.MethodDelete)
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost)
	}
	film := r.PathPrefix("/films").Subrouter()
	{
		film.HandleFunc("/{genre}", filmsHandler.FilmByGenre).Methods(http.MethodGet)

	}

	http.Handle("/", r)
	log.Print("store running on: ", srv.Addr)
	return srv.ListenAndServe()
}
