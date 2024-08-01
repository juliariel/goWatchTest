package main

type testStruct struct {
	Field1 int
	Field2 string
}

type testInterface interface {
	Hallo(int) int
	Bye(int) int
}

func main() {
	println("Hello World")
}
