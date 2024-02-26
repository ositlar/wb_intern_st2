package interfaces

type Visitor interface {
	VisitForLion(Animal)
	VisitForTiger(Animal)
	VisitForDolphin(Animal)
}
