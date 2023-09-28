package main

// 1 gb data
type BigData struct {
}

func doSomething(data BigData) error {
	// ..
	// ..
	// do something
	return nil
}

func doSomethingOptimal(data *BigData) error {
	// ..
	// ..
	// do something
	return nil
}

func main() {
	userData := BigData{}
	for i := 0; i > 5; i++ {
		// each time the 1 gb has to copied
		doSomething(userData)
	}

	cartData := &BigData{}
	for i := 0; i > 5; i++ {
		// each only 8 byte long pointer is copied
    // this is one of the reasons why we use pinter
    // do avoid large duplication
		doSomethingOptimal(cartData)
	}
}
