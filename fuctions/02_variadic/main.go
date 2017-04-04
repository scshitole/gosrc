package main

import "fmt"

func main() {
	s := average(1, 1, 1, 1, 1, 1)
	fmt.Println("Average is ", s)
}

func average(sf ...float64) float64 {

	fmt.Println(sf)
	fmt.Printf("%T  \n ", sf)
	var total float64
	for _, v := range sf {
		total += v
		fmt.Println("value of total", total)
	}
	return total / float64(len(sf))
}
