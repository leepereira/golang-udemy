package main

import "log"

func main() {
	pereiraPets := make(map[string]string) //[string]string  ---> [key]value

	pereiraPetsAge := make(map[string]int) //[string]int  ---> [key]value

	pereiraPets["bird1"] = "Luna"
	pereiraPets["bird2"] = "Rio"

	log.Println("Pets in the house are :", pereiraPets["bird1"], pereiraPets["bird2"])

	pereiraPets["bird2"] = "Pickles"
	log.Println("Pets in the house are :", pereiraPets["bird1"], pereiraPets["bird2"])

	pereiraPetsAge["bird1"] = 2
	pereiraPetsAge["bird2"] = 1
	log.Println("Age of Pets in the house are :", pereiraPetsAge["bird1"], pereiraPetsAge["bird2"])

}
