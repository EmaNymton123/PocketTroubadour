package controllers

import (
	"net/http"
)

func Routes(public_files http.FileSystem) http.Handler{
	var index=NewIndexControlleur()
	assets_controller := http.FileServer(public_files)

	mux := http.NewServeMux()
	mux.Handle("/tuto", NewTutoControlleur())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.URL.Path == "/" {
			index.ServeHTTP(w, r)
			return
		}

		assets_controller.ServeHTTP(w, r)
	})
	return mux
}
