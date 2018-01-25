package rec

func ExampleProtect() {
	var s []byte
	protect(func(){s[0]=0})
	protect(func(){panic(42)})
	// Output:
	// start
	// done
	// run time panic: runtime error: index out of range
	// start
	// done
	// run time panic: 42
}

