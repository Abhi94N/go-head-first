package rest

import (
	"github.com/Abhi94N/go-head-first/candyshop/pkg/adding"
	"github.com/Abhi94N/go-head-first/candyshop/pkg/reading"
	"github.com/gorilla/mux"
)

func InitHandlers(rs reading.Service, as adding.Service) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies/", getAllCandiesHandler(rs)).Methods("GET")
	router.HandleFunc("/api/candy/", addCandy(as)).Methods("POST")
	return router
}
