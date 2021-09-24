package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	sv "gosoft.co.th/mygo-lab/service"
)

func HandlerRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/profile/submit", sv.ProfileSubmitSv).Methods("POST")
	r.HandleFunc("/profile/detail", sv.ProfileDetailSv).Methods("GET")
	r.HandleFunc("/profile/list", sv.ProfileListSv).Methods("POST")
	r.HandleFunc("/profile/delete", sv.ProfileDeleteSv).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	return handler
}
