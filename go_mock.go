package Go_demo

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
}
