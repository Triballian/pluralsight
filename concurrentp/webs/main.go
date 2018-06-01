package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=GOOGL")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteReponse)
	xml.Unmarshal(body, &quote)
	fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

type QuoteReponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MarketCap        int
	Volume           int
	changeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
