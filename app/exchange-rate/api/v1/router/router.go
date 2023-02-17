package router

import (
	"fmt"
	"net/http"
	"os"

	"my-go-app/app/exchange-rate/api/v1/handlers"

	_ "my-go-app/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RunRouter() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router.Mount("/swagger", httpSwagger.WrapHandler)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})

	fmt.Println("Running server at port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}
