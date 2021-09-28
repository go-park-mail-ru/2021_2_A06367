package main

import (
	"context"
	"fmt"
	authDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/http"
	authRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	filmsDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/http"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {
	r := mux.NewRouter()
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8080")}
	//os.Setenv("SECRET", "DEBUG")
	conn := ""

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}

	//TODO на этом этапе вытаскивать секретный ключ и класть в конструкторк генератора токена
	//в данный момент это делается внутри метода GetToken
	tokenGenerator := authUsecase.NewTokenator()
	onlineRepo := authRepository.NewOnlineRepo(pool)
	authRepo := authRepository.NewAuthRepo(pool)
	authUse := authUsecase.NewAuthUsecase(authRepo, tokenGenerator)
	authHandler := authDelivery.NewAuthHandler(authUse, onlineRepo)

	filmsHandler := filmsDelivery.FilmsHandler{}

	auth := r.PathPrefix("/user").Subrouter()
	{
		auth.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
		auth.HandleFunc("/logout", authHandler.Logout).Methods(http.MethodDelete)
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost)
		auth.HandleFunc("/auth/{user}", authHandler.AuthStatus).Methods(http.MethodGet)
	}
	film := r.PathPrefix("/films").Subrouter()
	{
		film.HandleFunc("/{genre}", filmsHandler.FilmByGenre).Methods(http.MethodGet)
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
