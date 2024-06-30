package main

import "fmt"

func main() {
	i := 10
	p := &i
	fmt.Printf("Value of i = %d\n", i)
	fmt.Printf("Addres of &i = %p\n", &i)
	fmt.Printf("Pointer (address) of p = %p\n", p)
	fmt.Printf("Value Pointer of p %d\n", *p)
}
