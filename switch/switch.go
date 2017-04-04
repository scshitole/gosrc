package main

import "fmt"

func main() {
	switch "sanjay" {
	case "Vidya":
		fmt.Println("whats up Vidya how are you ? ")
	case "sanjay":
		fmt.Println("whats up Sanjay how are you ? ")
		fallthrough
	case "Ria":
		fmt.Println("whats up Ria how are you ? ")
	default:
		fmt.Println("No friends ")

	}

}
