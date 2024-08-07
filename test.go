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

type iface interface {
   bla(int) int
   MyInterface
   AnyInterface
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
