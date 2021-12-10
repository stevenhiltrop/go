package main

var availble = make(chan Buffer, 10)
var toController = make(chan Buffer)

type Buffer struct{}

func (b Buffer) process() {}
func (b Buffer) get()     {}

func worker() {
	for {
		var b Buffer
		select {
		case b = <-availble:
		default:
			b = Buffer{}
		}
		b.get()
		toController <- b
	}
}

func controller() {
	b := <-toController
	b.process()

	select {
	case availble <- b:
	default:
	}
}

func main() {
	go controller()
	go worker()
}
