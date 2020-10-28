package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	file, _ := os.Open("/etc/mysql/my.cnf")
	b, _ := ioutil.ReadAll(file)
	fmt.Print(string(b))
}
