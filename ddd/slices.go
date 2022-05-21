package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"aa1", "uwuiwi", "gdwoq"}
	fmt.Println("welcome to the video on the slices", fruitList)
	fruitList = append(fruitList, "dddd", "ffff", "wiwiw")
	fmt.Println(fruitList)

	fruitList = append(fruitList[2:4])
	fmt.Println(fruitList)
	highScores := make([]int, 4)
	highScores[0] = 323
	highScores[1] = 565
	highScores[2] = 321
	highScores[3] = 432
	fmt.Println(highScores)
	highScores = append(highScores, 765, 800, 909)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
}
