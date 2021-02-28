package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	envvariables := os.Environ()
	log.Printf("Printing env variables")
	for x := range envvariables {
		log.Printf("%v", envvariables[x])
	}

	err := readKeyFile("/root/go/bin/certificate.crt")
	if err != nil {
		log.Printf("Error while opening crt %v", err)
	}
	err = readKeyFile("/root/go/bin/certificate.key")
	if err != nil {
		log.Printf("Error while opening key %v", err)
	}
	err = readKeyFile("/root/go/bin/certificate.chain")
	if err != nil {
		log.Printf("Error while opening chain %v", err)
	}
}

func readKeyFile(filename string) error {

	// open private key
	_, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("readKeyFile|File error %w", err)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("GetRSAPublicKey|Error reading public key, path %w", err)

	}
	fmt.Printf("ReadKeyFile|Filename %v\n", filename)
	return nil
}
