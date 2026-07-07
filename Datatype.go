package main

import "fmt"

func main(){
	var price float64 = 19.99
	var quantity int = 5
	var total float64 = price * float64(quantity)
	
	fmt.Println("Price:", price)
	fmt.Println("Quantity:", quantity)
	fmt.Println("Total:", total)
}
