package main

import "fmt"


func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i := 0; i < len(pow); i++ {
		//if pow[i]^2 == 0 {
		//	fmt.Println("2**", pow[i],i)
		//}
		fmt.Printf("2**%v = %v\n",i, pow[i])
	}
}