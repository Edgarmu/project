package main

import (
	"log"
	todo "project"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error")
	}

}
