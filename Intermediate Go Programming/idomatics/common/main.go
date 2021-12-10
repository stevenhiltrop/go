package main

func main() {
	/* Indentation problem:
	if f, err := os.Open("/tmp/myfile"); err != nil {
		return
	} else {
		if f.Chdir()
			if
				if
					...
	}

	Solution: declare far left so usuable everywhere
	var f *os.File
	var err error
	*/

	/* Long If-else problem
	var n int
	if n < 0 {
		//
	} else if n > 0 {
	} else if n == 0 {
	}

	Solution: switch case
	switch {
	case n < 0:
	case n > 0:
	case n == 0:
	}
	*/
}
