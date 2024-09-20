package executor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anisbhsl/fetch-backend-assignment/api/index"
	"github.com/anisbhsl/fetch-backend-assignment/api/receipts"
	"github.com/anisbhsl/fetch-backend-assignment/routes"
	"github.com/anisbhsl/fetch-backend-assignment/store"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
)

var (
	timeout = 15 * time.Second
)

type Executor struct {
	Config             utils.AppConfig
	IndexAPIService    index.Service
	ReceiptsAPIService receipts.Service
}

// NewExecutor provides an instance of Executor by
// initializing all the services required to run this
// application
func NewExecutor(config *utils.AppConfig) *Executor {
	var inMemStore store.Service = store.NewInMemDB()

	return &Executor{
		Config:             *config,
		IndexAPIService:    index.NewIndexAPIService(),
		ReceiptsAPIService: receipts.NewReceiptsAPIService(inMemStore),
	}
}

// Executor spins up the http server
// at specified port
func (e *Executor) Execute() {
	// spin up the http server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", e.Config.HostAddr, e.Config.Port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
		Handler:      routes.RegisterRoutes(e.IndexAPIService, e.ReceiptsAPIService),
	}

	fmt.Println("server listening....") //todo: change this to logger
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
