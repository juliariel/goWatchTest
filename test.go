package main

type testStruct2 struct {
	Field3 string
	Field4 int
}

type testStruct struct {
	Field1 string
	Field2 int
	testStruct2
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
