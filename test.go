package main

type testStruct struct {
	Field1 bool
	Field2 int
}

type testInterface interface {
	Hallo(int) string
	Bye(int) int
}

func main() {
	println("Hello World")
}
