package main

import "fmt"

func main() {
	for i := 16000; i < 17000; i++ {
		fmt.Println(i, "----", string(i), "-----", []byte(string(i)))
	}
}
