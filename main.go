package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ExchangeRate struct {
	Usd struct {
		Code        string  `json:"code"`
		Rate 		float64 `json:"inverseRate"`
	} `json:"usd"`
}

func main() {
	r, err := http.Get("http://www.floatrates.com/daily/try.json")

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	// status control check
	if r.StatusCode != http.StatusOK {
		fmt.Println("Status not successful")
	}

	body, err := ioutil.ReadAll(r.Body)

	var exchangeRate ExchangeRate
	json.Unmarshal(body, &exchangeRate)

	if err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(exchangeRate)
}
