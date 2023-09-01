package main

import (
	"log"
	"sort"
)

func main() {

	var pereiraSlice []string
	pereiraSlice = append(pereiraSlice, "Leon")
	pereiraSlice = append(pereiraSlice, "Grinal")
	pereiraSlice = append(pereiraSlice, "Chelsea")
	pereiraSlice = append(pereiraSlice, "Hayden")

	log.Println(pereiraSlice) // 2023/08/31 22:53:07 [Leon Grinal Chelsea Hayden]

	var pereiraAgeSlice []int
	pereiraAgeSlice = append(pereiraAgeSlice, 13)
	pereiraAgeSlice = append(pereiraAgeSlice, 47)
	pereiraAgeSlice = append(pereiraAgeSlice, 45)
	pereiraAgeSlice = append(pereiraAgeSlice, 16)

	log.Println(pereiraAgeSlice) //2023/08/31 22:55:30 [47 45 16 13]

	sort.Ints(pereiraAgeSlice)
	log.Println(pereiraAgeSlice)

	// alternative approach to defining a slice

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)      //2023/08/31 23:00:38 [1 2 3 4 5 6 7 8 9 10]
	log.Println(numbers[0:2]) //2023/08/31 23:00:38 [1 2]

	names := []string{"Leon", "Grinal", "Chelsea", "Hayden"}
	log.Println(names) //2023/08/31 23:01:58 [Leon Grinal Chelsea Hayden]
}
