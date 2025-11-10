package main

import "fmt"

func main() {
	firstName := "Kurosaki"
	lastName := "Ichigo"

	fmt.Println("Hello '", firstName, lastName, "'")

	// dengan menggunakan print format
	fmt.Printf("Hello '%s %s'", firstName, lastName)
}
