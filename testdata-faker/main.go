package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	// step 1
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Email())
	fmt.Println(gofakeit.Phone())
	fmt.Println(gofakeit.BS())
	fmt.Println(gofakeit.BeerName())
	fmt.Println(gofakeit.Color())
	fmt.Println(gofakeit.Company())
	fmt.Println(gofakeit.HackerPhrase())
	fmt.Println(gofakeit.JobTitle())
	fmt.Println(gofakeit.CurrencyShort())

	// step 2
	fmt.Println(gofakeit.CreditCardType())
	fmt.Println(gofakeit.CreditCardNumber(nil))

	// step 3
	gofakeit.Seed(0)
	ccInfo := gofakeit.CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
}
