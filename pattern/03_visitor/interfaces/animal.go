package interfaces

type Animal interface {
	GetType() string
	Accept(Visitor)
}
