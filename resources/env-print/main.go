package main

import (
	"log"
	"os"
)

func main() {
	envvariables := os.Environ()
	log.Printf("Printing env variables")
	for x := range envvariables {
		log.Printf("%v", envvariables[x])
	}
}
