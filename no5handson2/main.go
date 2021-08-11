package main

import "fmt"

func main() {
	m := map[string]int{
		"answer" : 48,
	}
	fmt.Println("The value:", m["answer"])

	var elem, ok = m["answer"]

	//if ok {
		fmt.Println("The value:", elem, "Present?", ok)
	//}
}