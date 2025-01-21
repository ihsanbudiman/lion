package main

import (
	"flag"
	"fmt"
	"lion/config"
	"lion/driver/db"
	app_http "lion/internal/handler/http"
	"lion/internal/repository"
	"lion/internal/seeder"
	"lion/internal/usecase"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conf := config.NewConfig()
	handleArgs(conf)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// init db connection
	dbSource := db.NewDBConn(conf.DBSource)
	dbDestination := db.NewDBConn(conf.DBDestination)

	// init repository
	sourceProductRepo := repository.NewSourceProductRepo(dbSource)
	destinationProductRepo := repository.NewDestinationProductRepo(dbDestination)

	// init usecase
	productUsecase := usecase.NewProductUsecase(sourceProductRepo, destinationProductRepo)

	// init handler
	productHandler := app_http.NewProductHandler(productUsecase)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/sync-product", productHandler.SyncProduct)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)
}

func handleArgs(conf config.Config) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			dbSource := db.NewDBConn(conf.DBSource)
			dbDestination := db.NewDBConn(conf.DBDestination)

			seeder.Seed(dbSource, dbDestination)
			os.Exit(0)
		}
	}
}
