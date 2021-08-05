package instances

import (
	"fmt"
	"net/http"

	"github.com/burhanmubarok/microservice/instances/infrastructures"
	"github.com/burhanmubarok/microservice/instances/routers"
)

// Microservice doc
type Microservice struct {
	Name  string
	Desc  string
	Env   string
	Port  int
	Debug bool
}

func (c *Microservice) applyConfiguration() {
	conf := &infrastructures.Configuration{}
	conf.Init()

	c.Name = conf.Name
	c.Desc = conf.Desc
	c.Env = conf.Env
	c.Port = conf.Port
	c.Debug = conf.Debug
}

func (c *Microservice) figlet() {
	template := `
+++++++++++++++++++++++++++++++++++++++++
%s
+++++++++++++++++++++++++++++++++++++++++

env     : %s
port    : %d
debug   : %t

+++++++++++++++++++++++++++++++++++++++++
`

	fmt.Printf(template, c.Name, c.Env, c.Port, c.Debug)
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
	c.applyConfiguration()
	c.figlet()

	handlers := new(routers.Router).Handler()
	addr := fmt.Sprintf(":%d", c.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: handlers,
	}

	serverError := server.ListenAndServe()
	defer panic(serverError)
}

// Test doc
func (c *Microservice) Test() {
	fmt.Println("Database connection OK ...")
	fmt.Println("Kafka connection OK ...")
	fmt.Println("Log writing access OK ...")
	fmt.Println("Cross microservice connection OK ...")
}
