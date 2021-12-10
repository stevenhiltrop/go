package main

import "fmt"

type MyString struct {
	str string
}

func NewMyString(s string) MyString {
	return MyString{str: s}
}

func (m MyString) Output() {
	fmt.Println(m.str)
}

func main() {
	hello := NewMyString("Hello, World!")
	hello.Output()
}
