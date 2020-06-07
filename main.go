package main

import (
	"fmt"
	"log"
)

func main() {
	if err := generate(); err != nil {
		log.Fatalln("failed to generate file", err)
	}

	fmt.Println("Current Time:", Now())
}
