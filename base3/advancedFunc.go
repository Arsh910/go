package main

import "fmt"

// Higher order function
// func agg(a int, b int, add func(int, int) int) int {
// 	return add(a, b)
// }

// func addd(a int, b int) int {
// 	return a + b
// }

// currying
// func mul(x int, y int) int {
// 	return x * y
// }

// func square(mul func(int, int) int) func(int) int {
// 	return func(x int) int {
// 		return mul(x, x)
// 	}
// }

// defer

// func copyFile(src string, to string) (int, error) {
// 	file, err := os.Open(src)

// 	if err != nil {
// 		return 0, err
// 	}

// 	// called automatically at end of function
// 	defer file.Close()

// 	return 1, nil
// }

// closures , func that refer var outside its scope
// func checkCount() func() int {
// 	count := 1
// 	return func() int {
// 		return count
// 	}
// }

func main() {
	// HOF
	// fmt.Println(agg(10, 10, addd))

	// Currying
	// sq := square((mul))
	// fmt.Println(sq(5))

	// defer
	// fmt.Println(copyFile("1", "2"))

	// closure
	// s := checkCount()
	// fmt.Println(s())

	// anonymus func
	s := func() int {
		return 1
	}
	fmt.Println(s())
}
