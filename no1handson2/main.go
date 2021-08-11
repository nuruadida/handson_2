package main

import "fmt"

func main()  {
	// N0.1
	var a [2] string
	a[0] = "Go"
	a[1] = "C"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var primes = [6] int {2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//no.2
	var aprimes = primes[3:6]
	fmt.Println(aprimes)

}
