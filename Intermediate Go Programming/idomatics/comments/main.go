// Package dinosaur provides functions for creating and manipulating dinosaurs
package dinosaur

import "io"

// A Dinosaur is single dinosaur with name, weight and whether it is a herbivore
type Dinosaur struct {
	Name      string
	Weight    int
	Herbivore bool
}

// WriteName writes the name of the dinosaur to the provided io.Writer
func (d *Dinosaur) WriteName(w io.Writer) error {
	return nil
}

func main() {}
