package main

import (
	"context"
	"fmt"
	authDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/http"
	authRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/config"
	filmsDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/http"
	filmsRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/repo"
	filmsUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
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

	r.Use(middleware.CORSMiddleware)
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8000")}

	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}

	encrypter := authUsecase.NewEncrypter(os.Getenv("SECRET"))
	tokenGenerator := authUsecase.NewTokenator()
	onlineRepo := authRepository.NewOnlineRepo(pool)
	onlineUsecase := authUsecase.NewOnlineUsecase(onlineRepo)
	authRepo := authRepository.NewAuthRepo(pool)
	authUse := authUsecase.NewAuthUsecase(authRepo, tokenGenerator, encrypter)
	authHandler := authDelivery.NewAuthHandler(authUse, onlineUsecase)

	filmsRepo := filmsRepository.NewFilmsRepo(pool)
	filmsUse := filmsUsecase.NewFilmsUsecase(filmsRepo)
	filmsHandler := filmsDelivery.NewFilmsHandler(filmsUse)

	auth := r.PathPrefix("/user").Subrouter()
	{
		auth.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
		auth.HandleFunc("/logout", authHandler.Logout).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost)
		auth.HandleFunc("/auth", authHandler.AuthStatus).Methods(http.MethodGet)
	}
	film := r.PathPrefix("/films").Subrouter()
	{
		film.HandleFunc("/genre/{genre}", filmsHandler.FilmByGenre).Methods(http.MethodGet)
		film.HandleFunc("/selection/{selection}", filmsHandler.FilmBySelection).Methods(http.MethodGet,
			http.MethodOptions)
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
