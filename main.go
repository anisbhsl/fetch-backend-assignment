package main

import (
	"flag"

	"github.com/anisbhsl/fetch-backend-assignment/executor"
	"github.com/anisbhsl/fetch-backend-assignment/utils"
	"github.com/go-playground/validator/v10"
)

func init() {
	utils.Validate = validator.New(validator.WithRequiredStructEnabled())
}

func main() {
	host := flag.String("host", "127.0.0.1", "host address")
	port := flag.String("port", "3000", "port to listen to")
	flag.Parse()

	utils.AppParams = &utils.AppConfig{
		HostAddr: *host,
		Port:     *port,
	}

	executor.NewExecutor(utils.AppParams).Execute()
}
