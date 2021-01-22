package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}

	fmt.Println(s)

	//exercise 1
	// fmt.Println(strings.Join(os.Args[0:], " "))

	//exercise 2
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println("index: ", i, " value: ", os.Args[i])
	// }

	//exercise 3
	// start1 := time.Now()
	// fmt.Println(strings.Join(os.Args[0:], " "))
	// secs1 := time.Since(start1).Seconds()
	// fmt.Println("secs1: ", secs1)

	// start2 := time.Now()
	// for i := 1; i < len(os.Args); i++ {
	// 	// fmt.Println("index: ", i, " value: ", os.Args[i])
	// }
	// secs2 := time.Since(start2).Seconds()
	// fmt.Println("secs2: ", secs2)

	// start3 := time.Now()
	// for i := 1; i < len(os.Args); i++ {
	// 	// fmt.Println("index: ", i, " value: ", os.Args[i])
	// }
	// secs3 := time.Since(start3).Seconds()
	// fmt.Println("secs3: ", secs3)

}

/**
In each iterat ion of the loop, range produces a pair of values: the index and the value of the
element at that index. In this example, we don’t need the index, but the syntax of a range lo op
re quires that if we deal wit h the element, we must deal wit h the index too. One ide a would be
to assig n the index to an obv iou sly temporar y var iable like temp and ignore its value, but Go
do es not per mit unu sed local variables, so this wou ld result in a compi lat ion error.
The solut ion is to use the bl ank ident ifier, whose name is _ (t hat is, an underscore).
**/

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 	s += sep + arg
// 	sep = " "
// 	}
// 	fmt.Println(s)
// }
