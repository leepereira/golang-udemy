package main

import "log"

type car struct {
	Make       string
	Model      string
	Mileage    int
	FrontWheel wheel
	RearWheel  wheel
}

type wheel struct {
	Radius   int
	Material string
}

func main() {
	//defining a struct
	myCar := car{}
	myCar.FrontWheel.Radius = 30
	myCar.Make = "Tesla"
	myCar.Mileage = 100
	myCar.Model = "Model 3"

	log.Println(myCar)

	whosTheDriver(myCar)

}

func whosTheDriver(leonscar car) bool {
	if leonscar.Make == "Tesla" {
		log.Println("Driver will be Leon")
	} else {
		log.Println("Driver is Grinal")
	}
	return true
}
