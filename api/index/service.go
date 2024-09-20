package index

import "net/http"

type Service interface {
	Index() http.HandlerFunc
}

type service struct {
}

func NewIndexAPIService() Service {
	return &service{}
}
