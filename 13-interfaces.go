package main

import (
	"fmt"
)

type PereiraCars interface {
	PereiraCarMileage() float64
}

type Honda struct {
	distance     float64
	fuel         float64
	averagespeed string
}

type Tesla struct {
	distance float64
	charge   float64
}

// To implement the PereiraCars interface, we need
// that the Honda and Tesla struct should implement the method Mileage

// these are the methods that are created
// here h is the receiver and it is associated with the type (Honda in this case) that the method Mileage is associated with
// func (receiver) method() return
func (h Honda) PereiraCarMileage() float64 {
	return h.distance / h.fuel
}

func (t Tesla) PereiraCarMileage() float64 {
	return t.distance / t.charge
}

// add both the honda and tesla mileage in a month
// function will take a slice of PereiraCars interface as a parameter

// you should pass a slice of objects
// that implement the PereiraCars interface
// as an argument when calling totalMileage
// this function iterates over the slice : [{310.94 18 80} {310 95}]
func totalMileage(m []PereiraCars) { //[{310.94 18 80} {310 95}]
	tm := 0.0
	// var indexleon int
	for _, v := range m {
		tm = tm + v.PereiraCarMileage()
	}
	fmt.Println("Total Mileage / month ", tm)
}

func main() {
	h1 := Honda{
		distance:     310.94,
		fuel:         18,
		averagespeed: "80",
	}

	t1 := Tesla{
		distance: 310,
		charge:   95,
	}

	//  create a slice person containing both vehicles.
	// person is actually this slice : [{310.94 18 80} {310 95}]
	// which contains instances of Honda and Tesla structs
	// both of which implement the PereiraCars interface by defining the Mileage() method.
	// []PereiraCars declares a slice of the PereiraCars interface type.
	// This means that the person slice can hold any values that implement the PereiraCars interface.
	// This is possible because both Honda and Tesla,
	// implement the PereiraCars interface by defining the Mileage() method.
	person := []PereiraCars{h1, t1}
	fmt.Println(person)
	totalMileage(person)
}
