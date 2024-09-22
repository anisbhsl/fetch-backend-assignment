package index

import "net/http"

type Service interface {
	Index() http.HandlerFunc
}

// service implements Index Service interface
type service struct {
}

func NewIndexAPIService() Service {
	return &service{}
}
