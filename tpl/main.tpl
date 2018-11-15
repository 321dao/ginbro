package main

import (
	"{{.ProjectPackage}}/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
