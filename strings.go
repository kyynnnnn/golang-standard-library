package main

import (
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains("Raditya Adi Santoso", "Adi")
	split := strings.Split("Raditya Adi Santoso", " ")
	toLower := strings.ToLower("Raditya Adi Santoso")
	toUpper := strings.ToUpper("Raditya Adi Santoso")
	trim := strings.Trim("           Raditya Adi Santoso            ", " ")
	replaceAll := strings.ReplaceAll("Raditya Adi Santoso Adi Doremi", "Adi", "Naruto")

	fmt.Println(contains)
	fmt.Println(split)
	fmt.Println(toLower)
	fmt.Println(toUpper)
	fmt.Println(trim)
	fmt.Println(replaceAll)
}
