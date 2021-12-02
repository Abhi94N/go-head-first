package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Abhi94N/go-head-first/candyshop/pkg/adding"
)

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCandy adding.Candy
		if err := json.NewDecoder(r.Body).Decode(&newCandy); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		id, err := as.AddCandy(newCandy)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		//set id from storage
		newCandy.Id = id
		json.NewEncoder(w).Encode(newCandy)
	}
}
