package utils

// import "github.com/go-playground/validator/v10"

var AppParams *AppConfig

type AppConfig struct {
	HostAddr string
	Port     string
}
