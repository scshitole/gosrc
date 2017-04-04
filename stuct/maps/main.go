package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 100
	m["k2"] = 200
	fmt.Println("map", m)
	delete(m, "k2")
	v, ok := m["k1"]
	fmt.Println("ok?: ", ok, v)
}
