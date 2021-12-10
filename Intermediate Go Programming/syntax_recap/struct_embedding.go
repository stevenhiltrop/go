package main

import (
	"fmt"
	"strings"
	"sync"
)

type MyString struct {
	sync.Mutex
	str string
}

type Shouting struct {
	MyString
}

func NewMyString(s string) MyString {
	return MyString{str: s}
}

func (m MyString) Output() {
	fmt.Println(m.str)
}

func NewShoutingString(s string) Shouting {
	load := Shouting{}
	load.str = strings.ToUpper(s)
	return load
}

// Overriding the parent Output
func (m Shouting) Output() {
	fmt.Printf("Really loud: %s\n", m.str)
}

func main() {
	hello := NewShoutingString("Hello, World!")
	hello.Lock()
	hello.Output()
	hello.Unlock()
}
