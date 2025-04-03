package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler interface {
	Ping(http.ResponseWriter, *http.Request)
	Draw(http.ResponseWriter, *http.Request)
}

func NewRouter(handler Handler) (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/draw", handler.Draw).Methods(http.MethodGet)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})

	return router, nil
}
