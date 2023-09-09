package main

import "fmt"

func main() {
	i := 42
	fmt.Println(&i) //0xc000012028
	fmt.Println(i)  // 42

	var p *int //Declare p as a pointer to an int
	p = &i
	fmt.Println(p) //0xc000012028

	fmt.Println(*p) // read i through the pointer p : 42
	*p = 21         // set i through the pointer p
	fmt.Println(p)  // 0xc000012028
	fmt.Println(i)  // 21

}
