package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	url = "https://api.hgbrasil.com/finance?key=ce1e33e9"
)

func getApiData(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var stock Stock
	err = json.NewDecoder(response.Body).Decode(&stock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stock)
}
