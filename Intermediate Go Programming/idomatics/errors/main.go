package main

/* Returning pointer problem:
func doSomething() *MyType {
	// If something goes wrong
	return nil
	// If all goes well
	return &MyType{}
}

// Solution: always return an error
func doSomething() (*MyType, error) {
	// If something goes wrong
	return nil, error.new("Gone wrong")
	// If all goes well
	return &MyType{}, nil
}
*/

func main() {
	//err := doSomething()
}
