package topic

import (
	"log"
	"net/http"
)

type Topic struct {
	logger log.Logger
}

func (t *Topic)SetupRouter(mux http.ServeMux) {
	mux.HandleFunc()
}
