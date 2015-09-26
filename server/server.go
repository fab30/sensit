package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joelvim/sensit/timeseries"
	"github.com/joelvim/sensit/timeseries/influxdb"
)

// HTTPConfig is a class carrying the http server configuration.
// It only supports a simple basic authentication with only one login/password
// It is not secured, but as the callback configuration
// displays the credentials in cleartext, implementing a complex and
// as-secure-as-possible auth solution is useless
type HTTPConfig struct {
	ListenPort int    `json:"-"`
	AuthType   string `json:"authType"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
}

/*
API creates an http server that handles the request to the metrics API
It takes 5 arguments

	dbConfig contains the configuration to connect the database
	httpConfig contains the configuration to launch the http server

Note that this server does not support http authentication, the only authentication is made at db level, nor SSL/TLS
*/
func API(dbConfig timeseries.DBConfig, httpConfig HTTPConfig) {
	db, err := influxdb.Connect(dbConfig)

	if err != nil {
		log.Fatalf("Could not connect influxdb : %s", err)
	}

	var authHandler AuthHandler
	switch httpConfig.AuthType {
	case "Basic":
		authHandler = BasicAuthHandler{httpConfig.Login, httpConfig.Password}
	case "Token":
		authHandler = TokenAuthHandler{httpConfig.Login, httpConfig.Password, httpConfig.Salt}
	default:
		log.Fatalf("Unknown authentication method : %s", httpConfig.AuthType)
		return
	}

	router := NewRouter(db, authHandler)

	boundAddress := fmt.Sprintf(":%d", httpConfig.ListenPort)
	log.Printf("Listening on the following address : %s", boundAddress)
	log.Fatal(http.ListenAndServe(boundAddress, router))
}
