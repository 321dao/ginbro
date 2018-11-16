// Package classification User API.
//
// The purpose of this service is to provide an application that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta
package main

import (
	"github.com/dejavuzhou/ginbro/boilerplate/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
