package main

import (
	"fmt"
	"log"

	"github.com/mishrasunny174/Corona-Stats/api"
)

func main() {
	countries, err := api.GetAllCountryStats()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%V\n", countries)
}
