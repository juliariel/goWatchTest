package main

type testStruct struct {
	Field1 int
}

type testInterface interface {
	Hallo(string) bool
	Bye(string) int
}

func main() {
	println("Hello World")
}
