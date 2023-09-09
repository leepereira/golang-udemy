package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i : memory location of i in assigned to p : 0x14000112018
	fmt.Println(p)  // 0x14000112018
	fmt.Println(*p) // read i through the pointer : 42
	*p = 21         // set i through the pointer to be now 21
	fmt.Println(i)  // see the new value of i : now 21

	p = &j         // point to j : memory location of j
	*p = *p / 37   // divide j through the pointer : means 2701 / 37 and assign to *p
	fmt.Println(j) // see the new value of j
}
