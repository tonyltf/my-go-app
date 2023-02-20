package handlers

import (
	"context"
	"fmt"
	"my-go-app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", HealthCheck)
	r.Route("/price", func(r chi.Router) {
		r.Get("/{exchangePair}", GetLastExchangePrice)
		r.Get("/{exchangePair}/average", GetAvgExchangePrice)
	})

	return r
}

//	@Summary	Health checking
//	@Success	200	{string}	string	"ok"
//	@Router		/healthcheck [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

//	@Summary	Get the exchange price
//	@Tag		Exchange
//	@Param		exchange_pair	path		string	true	"Currency Pair"
//	@Param		timestamp		query		string	false	"timestamp"
//	@Success	200				{number}	number	price
//	@Failure	404				{string}	string	"Not found"
//	@Router		/price/{exchange_pair} [get]
func GetLastExchangePrice(w http.ResponseWriter, r *http.Request) {
	// TODO
	exchangePair := chi.URLParam(r, "exchangePair")
	timestamp := r.URL.Query().Get("timestamp")
	fmt.Printf("Price for %s at %s.\n", exchangePair, timestamp)

	rate := service.GetRate(context.Background(), exchangePair, &timestamp)
	if rate != nil {
		w.Write([]byte(*rate))
		return
	}
	http.Error(w, http.StatusText(404), 404)
	return
}

//	@Summary	Get the average exchange price
//	@Tag		Exchange
//	@Param		exchange_pair	path		string	true	"Currency Pair"
//	@Param		from			query		string	false	"From time"
//	@Param		to				query		string	false	"To time"
//	@Success	200				{number}	number	price
//	@Failure	404				{string}	string	"Not found"
//	@Router		/price/{exchange_pair}/average [get]
func GetAvgExchangePrice(w http.ResponseWriter, r *http.Request) {
	// TODO
	exchangePair := chi.URLParam(r, "exchangePair")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	fmt.Printf("Average price for %s from %s to %s.\n", exchangePair, from, to)
	w.Write([]byte(""))
}
