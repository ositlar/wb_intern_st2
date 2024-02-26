package main

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse()
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "Dacha":
		return
	case "":
	}
}

func main() {

}
