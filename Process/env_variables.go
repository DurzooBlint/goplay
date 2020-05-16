package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := os.Setenv("Gopherino", "doodino")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(os.Getenv("Gopherino"))
	err = os.Unsetenv("Gopherino")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("Gopherino") + "dsdsd")
}
