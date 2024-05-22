package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CurrentDate struct {
	Currency []Currency `xml:"Currency"`
}

type Currency struct {
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
}

func main() {
	url := "https://www.tcmb.gov.tr/kurlar/today.xml"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get XML: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var currentDate CurrentDate
	err = xml.Unmarshal(body, &currentDate)
	if err != nil {
		log.Fatalf("Failed to unmarshal XML: %v", err)
	}

	for _, currency := range currentDate.Currency {
		fmt.Printf("Currency: %s\n", currency.Isim)
		fmt.Printf("Code: %s\n", currency.CurrencyCode)
		fmt.Printf("Unit: %s\n", currency.Unit)
		fmt.Printf("Forex Buying: %s\n", currency.ForexBuying)
		fmt.Printf("Forex Selling: %s\n", currency.ForexSelling)
		fmt.Printf("Banknote Buying: %s\n", currency.BanknoteBuying)
		fmt.Printf("Banknote Selling: %s\n\n", currency.BanknoteSelling)
	}

}
