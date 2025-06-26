// go run 7_reverse_integer/main.go
package main

import "fmt"


func main(){
	fmt.Println(reverse(153423646))
}

func reverse(x int) int {

	var max int = 2147483647
	var min int = -2147483648
		
	var num int
	for x != 0 {
		var modulus int = x % 10
		x = x / 10

		if num >max/10 || (num == max/ 10 && modulus > 7) {
			return 0
		}
		if num < min/10 || (num == min/10 && modulus < -8){
			return 0
		}
		num = (num * 10) + modulus
	} 

	return num
}