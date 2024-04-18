package controllers

import (
	"net/http"
)

func Routes() http.Handler{
	mux := http.NewServeMux()
	mux.Handle("/", NewIndexControlleur())
	mux.Handle("/tuto", NewTutoControlleur())
	return mux
}
