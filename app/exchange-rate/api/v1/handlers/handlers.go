package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", HealthCheck)
	r.Route("/price", func(r chi.Router) {
		r.Get("/{exchangePair}", GetLastExchangePrice)
	})

	return r
}

//	@Summary	Health checking
//	@Success	200	{string}	string	"ok"
//	@Router		/healthcheck [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

//	@Summary	Get the last price
//	@Tag		Exchange
//	@Param		exchange_pair	path		string	true	"Currency Pair"
//	@Success	200				{number}		number		price
//	@Failure	404				{string}	string	"Not found"
//	@Router		/price/{exchange_pair} [get]
func GetLastExchangePrice(w http.ResponseWriter, r *http.Request) {
	exchangePair := chi.URLParam(r, "exchangePair")
	if exchangePair == "BTCUSD" {
		w.Write([]byte("24000"))
	} else {
		http.Error(w, http.StatusText(404), 404)
		return
	}
}
