package router

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"my-go-app/api/v1/handlers"

	_ "my-go-app/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRouter(ctx context.Context) {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/swagger", httpSwagger.WrapHandler)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Running server at port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}
