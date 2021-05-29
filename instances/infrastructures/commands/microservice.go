package instances

import (
	"fmt"
	"net/http"

	"bitbucket.org/burhanmubarok/microservice/instances/routers"
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

	handlers := new(instances.Router).Handler()
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
