  
package main

import (
	"fmt"
)

func main() {
	
	var helloworld = "Hello World"
	
	fmt.Println(helloworld)
	
	
	x := 11
	y := "bom dia"
	z := "olá"
	
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)
	
	x, u := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("u: %v, %T\n\n", u, u)
	
	
	adição := 40 + 150
	
	fmt.Println(adição)
	
	comparação1 := 30 != 50
	comparação2 := 40 == 90 
	comparação3 := 50 > 100
	comparação4 := 60 < 200
	
	fmt.Println(comparação1)
	fmt.Println(comparação2)
	fmt.Println(comparação3)
	fmt.Println(comparação4)

	
}
