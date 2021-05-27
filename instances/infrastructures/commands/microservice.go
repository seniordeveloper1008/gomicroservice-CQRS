package instances

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/structures/https"
	"github.com/gorilla/mux"
)

// Microservice doc
type Microservice struct{}

func (c *Microservice) figlet() {
	template := `
+++++++++++++++++++++++++++++++++++++++++
              Microservice
+++++++++++++++++++++++++++++++++++++++++

env     : %s
port    : %s
debug   : %v

+++++++++++++++++++++++++++++++++++++++++
`

	fmt.Printf(template, "development", "8008", true)
}

// Ensure doc
func (c *Microservice) Ensure() {
	fmt.Println("let's check yaml configuration values")
}

// Migrate doc
func (c *Microservice) Migrate() {
	fmt.Println("let's migrate a database and then init some data")
}

// Serve doc
func (c *Microservice) Serve() {
	c.figlet()

	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		data := structures.Response{
			Code:    http.StatusNotFound,
			Message: "URL not found",
		}
		json.NewEncoder(w).Encode(data)
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		data := structures.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		}
		json.NewEncoder(w).Encode(data)
	})

	router.Path("/").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		data := structures.Response{
			Code:    http.StatusOK,
			Message: "Wellcome! trust me, you are a coder",
		}
		json.NewEncoder(res).Encode(data)
	})

	router.Path("/about").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		data := structures.Response{
			Code:    http.StatusOK,
			Message: "Microservice [desc]",
		}
		json.NewEncoder(res).Encode(data)
	})

	handlers := router
	server := &http.Server{
		Addr:    ":8008",
		Handler: handlers,
	}

	serverError := server.ListenAndServe()
	panic(serverError)
}

// Test doc
func (c *Microservice) Test() {
	fmt.Println("Database connection ...")
	fmt.Println("Email connection...")
	fmt.Println("Kafka connection ...")
	fmt.Println("Log write access ...")
}
