package main

import "fmt"

func sum(nums ...int) {
	for i := 0; ; {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func main() {

	// for loop
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// while loop
	// for i := 0; ; {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// arrays
	// arr := [6]int{2, 3, 4, 5, 6, 7}
	// for i := 0; i < 6; i++ {
	// 	fmt.Println(arr[i])
	// }

	// slices
	// sli := []int{2, 3}
	// fmt.Println(sli)

	// sli := make([]int, 3, 5)
	// fmt.Println((sli))

	// sli := make([]int, 3, 3)
	// sli2 := append(sli, 4)
	// fmt.Println(sli2)

	// sli := []int{1, 2, 3}
	// sli2 := sli
	// sli2[0] = 10
	// fmt.Println(sli)

	// sli := []int{1, 2, 3}
	// sli2 := make([]int, len(sli))
	// copy(sli2, sli)
	// fmt.Print(sli2)

	// slices are reference to array
	// arr := [6]int{1, 2, 3, 4, 5, 6}
	// sli := arr[0:3]
	// sli2 := arr[3:5]
	// sli[0] = 10
	// sli2[1] = 10
	// fmt.Println((arr))

	// variadic function
	// sum(1, 2, 3)

	// spread
	// sli := []int{1, 2, 3}
	// sum(sli...)

	// maps
	// mp := map[string]int{
	// 	"h": 1,
	// 	"j": 2,
	// }
	// fmt.Println(mp["h"])
	// elem, ok := mp["h"]
	// fmt.Println(elem, ok)

	// delete(mp, "h")

	// elem2, ok2 := mp["h"]
	// fmt.Println(elem2, ok2)

	// not every thing can be used as map .
	// anything comparable can be key.

	// mp := map[string]map[string]int{
	// 	"first":  {"sec": 2},
	// 	"second": {"sec": 3},
	// }
	// fmt.Println(mp["first"]["sec"])

	fmt.Println("hello")

}
