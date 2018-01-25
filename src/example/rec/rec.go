package rec

import "fmt"

func protect(g func()){
	defer func() {
		fmt.Println("done")
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v\n", x)
		}
	}()

	fmt.Println("start")
	g()
}
