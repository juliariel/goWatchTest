package main

type testStruct struct {
	Field1 int
	Field2 int
}

type testInterface interface {
	Hallo(int) int
	Bye(int) int
}

type MyInterface interface {
   Func1(int)string
}

type iface interface {
   MyInterface
   bla(int) bool
}

func testfunc(test1 int, test2 int) {
	println("HI")
}

func main() {
	println("Hello World")
	testfunc(1, 2)
}
