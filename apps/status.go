package main

import (
	"fmt"
	"log"

	"github.com/ecc1/medtronic"
)

func main() {
	pump, err := medtronic.Open()
	if err != nil {
		log.Fatal(err)
	}
	result, err := pump.PumpStatus()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}