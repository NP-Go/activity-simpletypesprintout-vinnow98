package main

import "fmt"

func main() {
	text := "the following is the account information"
	firstName := "Luke"
	lastName := "Skywalker"
	age := 20
	weight := 73.0
	height := 1.72
	remainingCredit := 123.55
	accountName := "admin"
	accountPassword := "password"
	subscribedUser := true

	fmt.Printf("%v \n%v \n%v \n%v \n%v \n%v \n%v \n%v \n%v \n%v", text, firstName, lastName, age, weight, height, remainingCredit, accountName,
		accountPassword, subscribedUser)
}
