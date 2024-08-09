package main

type testStruct2 struct {
	Field3 int
	Field4 string
}

type testStruct struct {
	Field1 int
	Field2 int
	testStruct2
}

type testInterface interface {
	Hallo() string
	Bye(int) bool
}

type MyInterface interface {
   Func1(int) string
}

type newPerson person {
}

type iface interface {
   bla(int) int
   MyInterface
   AnyInterface
}

type animal interface {
	rabbit(string)
	cat() string
	dog()
	TigerInterface
	ghost
}

type person interface {
	firstperson(int, string, bool)
	lastperson() (int, string, bool)
	specialperson(int, string) (bool, int)
	oldperson(int) bool
	PersonInterface
	animal
}

type ghost interface {
}

func testfunc(test1 int, test2 int) {
	println("HI")
}

func main() {
	println("Hello World")
	testfunc(1, 2)
}
