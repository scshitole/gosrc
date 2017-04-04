package main

import "fmt"

func main() {
	var mymap = make(map[string]string)
	mymap["sanjay"] = "good morning"
	mymap["vidya"] = "namaskar"
	fmt.Println(mymap)
}
