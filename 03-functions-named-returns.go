package main

import (
	"errors"
	"log"
)

func main() {
	mult_out, div_out, _ := calculator(3222, 0)
	log.Println(mult_out, div_out)

}

func calculator(a int, b int) (mul int, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by Zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}
