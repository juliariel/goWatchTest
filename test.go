package main

type testStruct struct {
	Field1 int
	Field2 int
}

type testInterface interface {
	Hallo(int) string
	Bye(int) intg
}

type MyInterface interface {
   Func1(int)string
}

type iface interface {
   MyInterface
   bla(int) bool
}

func main() {
	println("Hello World")
}
