package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhi94N/go-head-first/candyshop/pkg/http/rest"
	"github.com/Abhi94N/go-head-first/candyshop/pkg/storage"
	"github.com/Abhi94N/go-head-first/candyshop/pkg/reading"
	"github.com/Abhi94N/go-head-first/candyshop/pkg/adding"
)

func main() {

	repo, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage:", err)
	}

	//create readingService which returns a service which is passed a repo object that represents the cassandra instance session
	rs := reading.NewService(repo)
	as := adding.NewService(repo)

	// candy, err := repo.GetAllCandyNames()
	// if err != nil {
	// 	log.Fatalln("error while getting candies in storage", err)
	// }
	// log.Println(candy)

	fmt.Println("Starting server on port 8080....")
	
	//passs reading reading service to init handler so that those routes can be set up
	router := rest.InitHandlers(rs, as)
	log.Fatal(http.ListenAndServe(":8080", router))
}
