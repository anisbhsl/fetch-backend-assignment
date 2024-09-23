package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	chain "github.com/justinas/alice"

	"github.com/anisbhsl/fetch-backend-assignment/api/index"
	"github.com/anisbhsl/fetch-backend-assignment/api/receipts"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
)

var httpMethods = struct {
	GET    string
	POST   string
	PUT    string
	PATCH  string
	DELETE string
}{
	"GET",
	"POST",
	"DELETE",
	"PUT",
	"PATCH",
}

type routeConfig struct {
	Handler     http.Handler
	Methods     []string
	Middlewares []chain.Constructor
}

var routes = func(
	indexApiService index.Service,
	receiptsApiService receipts.Service) map[string]routeConfig {
	GENERAL := map[string]routeConfig{
		"/": {
			Handler:     indexApiService.Index(),
			Methods:     []string{httpMethods.GET},
			Middlewares: nil,
		},
	}

	RECEIPTS := map[string]routeConfig{
		"/receipts/process": {
			Handler:     receiptsApiService.ProcessReceipts(),
			Methods:     []string{httpMethods.POST},
			Middlewares: nil,
		},
		"/receipts/{id}/points": {
			Handler:     receiptsApiService.ProcessReceiptPoints(),
			Methods:     []string{httpMethods.GET},
			Middlewares: nil,
		},
	}

	return func(routeMaps ...map[string]routeConfig) map[string]routeConfig {
		routeDefs := routeMaps[0]
		for i := 1; i < len(routeMaps); i++ {
			for path, config := range routeMaps[i] {
				routeDefs[path] = config
			}
		}
		return routeDefs
	}(GENERAL, RECEIPTS)
}

// RegisterRoutes registers all API routes by providing a wrapper
func RegisterRoutes(indexApiService index.Service,
	receiptsApiService receipts.Service) http.Handler {

	utils.GetLogger().Info("registering API routes.....")

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	for path, config := range routes(indexApiService, receiptsApiService) {
		r.Handle(path, chain.New(config.Middlewares...).Then(config.Handler)).Methods(config.Methods...)
	}

	return r
}
