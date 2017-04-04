package main

import "fmt"

func main() {

	greetings := map[int]string{

		0: "Good morning ",
		1: "su prabhat ",
		2: " Bonjour ",
		3: "Buenous Diase ",
		4: " Bongiorno ",
	}

	for key, val := range greetings {
		fmt.Println(key, "-----", val)
	}
}
