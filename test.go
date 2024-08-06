package main

type testStruct struct {
	Field1 string
	Field2 int
}

type testInterface interface {
	Hallo() int
	Bye(int) int
}

type MyInterface interface {
   Func1(int)int
}

type iface interface {
   MyInterface
   bla(int) int
}

func testfunc(test1 int, test2 int) {
	println("HI")
}

func main() {
	println("Hello World")
	testfunc(1, 2)
}
