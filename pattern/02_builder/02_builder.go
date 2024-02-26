package main

import (
	"fmt"

	"ositlar.com/pattern/02_builder/dacha"
	"ositlar.com/pattern/02_builder/model"
	"ositlar.com/pattern/02_builder/shop"
)

func getBuilder(builderType string) model.Builder {
	switch builderType {
	case "Dacha":
		return dacha.NewDachaBuilder()
	case "Shop":
		return shop.NewShopBuilder()
	}
	return nil
}

func main() {
	dachaBuilder := getBuilder("Dacha")
	shopBuilder := getBuilder("Shop")

	director := NewDirector(dachaBuilder)
	dacha := director.BuildHouse(10, "Omsk", 2)

	fmt.Printf("Dacha's square: %f\n", dacha.GetSquare())
	fmt.Printf("Dacha's town: %s\n", dacha.GetTown())
	fmt.Printf("Dacha's floors: %d\n", dacha.GetFloors())

	director.SetBuilder(shopBuilder)
	shop := director.BuildHouse(1, "Omsk", 1)
	fmt.Printf("Shop's square: %f\n", shop.GetSquare())
	fmt.Printf("Shop's town: %s\n", shop.GetTown())
	fmt.Printf("Shop's floors: %d\n", shop.GetFloors())
}
