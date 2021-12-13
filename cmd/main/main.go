package main

import (
	"context"
	"fmt"
	_ "github.com/go-park-mail-ru/2021_2_A06367/docs" // docs is generated by Swag CLI, you have to import it.
	actorsDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/delivery/http"
	actorsRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/repo"
	actorsUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/usecase"
	generated2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	authDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/http"
	authRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/config"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/grpc/generated"
	filmsDelivery "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/delivery/http"
	filmsRepository "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/repo"
	filmsUsecase "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/usecase"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/search/delivery"
	generated3 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	http2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/http"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

// @title LimeTV API
// @version v0.1.0
// @description This is a service for LimeTV project "A06367"
// @contact.name Lelya Kochkarova, Ivan Chernov, Slava Rianov
// @contact.email kochkarova.lelya@gmail.com, chernov-ivan.1998@yandex.ru, slavarianov@yandex.ru
// @host http://3.67.182.34:8080
// @BasePath /
func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {
	r := mux.NewRouter()

	//key, err := config.GetCsrfToken()
	//if err != nil {
	//	return err
	//}
	//protect := csrf.Protect(key, csrf.Secure(false))

	r.Use(middleware.CORSMiddleware)
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8000")}

	logger, err := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)

	zapSugar := logger.Sugar()

	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}

	filmsConn, err := grpc.Dial(
		"films:8010",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("cant connect to films grpc")
	}

	filmsClient := generated.NewFilmsServiceClient(filmsConn)

	authConn, err := grpc.Dial(
		"auth:8020",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("cant connect to session grpc")
	}

	authClient := generated2.NewAuthServiceClient(authConn)

	subConn, err := grpc.Dial(
		"subs:8030",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("cant connect to session grpc")
	}

	subsClient := generated3.NewSubsServiceClient(subConn)
	encrypter := authUsecase.NewEncrypter()
	tokenGenerator := authUsecase.NewTokenator()
	authRepo := authRepository.NewAuthRepo(pool, zapSugar)
	authUse := authUsecase.NewAuthUsecase(authRepo, tokenGenerator, encrypter, zapSugar)
	authHandler := authDelivery.NewAuthHandler(authClient, zapSugar)

	filmsRepo := filmsRepository.NewFilmsRepo(pool, zapSugar)
	filmsUse := filmsUsecase.NewFilmsUsecase(filmsRepo, zapSugar)
	filmsHandler := filmsDelivery.NewFilmsHandler(zapSugar, filmsClient, subsClient)

	actorsRepo := actorsRepository.NewActorsRepo(pool, zapSugar)
	actorsUse := actorsUsecase.NewActorsUsecase(actorsRepo, zapSugar)
	actorsHandler := actorsDelivery.NewActorsHandler(actorsUse, zapSugar)

	search := delivery.NewSearchHandler(filmsUse, authUse, actorsUse)

	h := http2.NewSubsHandler(subsClient)

	m := middleware.NewMiddleware(zapSugar)
	m2 := middleware.NewMetricsMiddleware()
	m2.Register(middleware.ServiceMainLabel)

	auth := r.PathPrefix("/api/users").Subrouter()
	//auth.Use(protect)
	{
		auth.HandleFunc("/secure", authHandler.Token).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
		auth.HandleFunc("/logout", authHandler.Logout).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost)
		auth.HandleFunc("/auth", authHandler.AuthStatus).Methods(http.MethodGet)
		auth.HandleFunc("/profile", authHandler.GetProfile).Methods(http.MethodGet)
		auth.HandleFunc("/profile/{id}/follow", authHandler.Follow).Methods(http.MethodPost)
		auth.HandleFunc("/profile/{id}/unfollow", authHandler.Unfollow).Methods(http.MethodDelete)
		auth.HandleFunc("/profile/settings/pic", authHandler.UpdateProfilePic).Methods(http.MethodPost)
		auth.HandleFunc("/profile/settings/pass", authHandler.UpdateProfilePass).Methods(http.MethodPost)
		auth.HandleFunc("/profile/settings/bio", authHandler.UpdateProfileBio).Methods(http.MethodPost)
	}

	film := r.PathPrefix("/api/films").Subrouter()
	{
		film.HandleFunc("/genre/{genre}", filmsHandler.FilmByGenre).Methods(http.MethodGet)
		film.HandleFunc("/selection/{selection}", filmsHandler.FilmBySelection).Methods(http.MethodGet)
		film.HandleFunc("/film/{film_id}", filmsHandler.FilmById).Methods(http.MethodGet)
		film.HandleFunc("/film/{id}/rating", filmsHandler.GetRating).Methods(http.MethodGet)
		film.HandleFunc("/film/{id}/rating", filmsHandler.SetRating).Methods(http.MethodPost)
		film.HandleFunc("/film/{id}/user/rating", filmsHandler.GetRatingByUser).Methods(http.MethodGet)
		film.HandleFunc("/selection/actor/{actor_id}", filmsHandler.FilmByActor).Methods(http.MethodGet)
		film.HandleFunc("/selection/user/personal", filmsHandler.FilmsByUser).Methods(http.MethodGet)
		film.HandleFunc("/selection", filmsHandler.FilmStartSelection).Methods(http.MethodGet)

		film.HandleFunc("/starred", filmsHandler.GetStarred).Methods(http.MethodGet)
		film.HandleFunc("/starred/check/{id}", filmsHandler.IfStarred).Methods(http.MethodGet)
		film.HandleFunc("/starred/{id}", filmsHandler.AddStarred).Methods(http.MethodPost)
		film.HandleFunc("/starred/{id}", filmsHandler.RemoveStarred).Methods(http.MethodDelete, http.MethodOptions)

		film.HandleFunc("/wl", filmsHandler.GetWatchlist).Methods(http.MethodGet)
		film.HandleFunc("/wl/{id}", filmsHandler.AddWatchlist).Methods(http.MethodPost)
		film.HandleFunc("/wl/{id}", filmsHandler.RemoveWatchlist).Methods(http.MethodDelete, http.MethodOptions)
		film.HandleFunc("/wl/check/{id}", filmsHandler.IfWl).Methods(http.MethodGet)
	}

	actors := r.PathPrefix("/api/actors").Subrouter()
	{
		actors.HandleFunc("/actor/{id}", actorsHandler.ActorsById).Methods(http.MethodGet)
		actors.HandleFunc("/film", actorsHandler.FetchActors).Methods(http.MethodPost, http.MethodGet)
	}

	searching := r.PathPrefix("/api/search").Subrouter()
	{
		searching.HandleFunc("/{keyword}", search.Search).Methods(http.MethodGet)
	}

	licensing := r.PathPrefix("/api/licenses").Subrouter()
	{
		licensing.HandleFunc("/licenses", h.SetLicense).Methods(http.MethodPost, http.MethodGet)
	}
	licensing2 := r.PathPrefix("/api/check").Subrouter()
	{
		licensing2.HandleFunc("/check", h.GetLicense).Methods(http.MethodGet)
	}

	// swag init -g ./cmd/main/main.go
	r.PathPrefix("/api-docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/api/metrics").Handler(promhttp.Handler())
	r.Use(m.LogRequest)
	r.Use(m2.LogMetrics)

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)

	return srv.ListenAndServe()
}
