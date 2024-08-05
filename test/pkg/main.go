package main

import "fmt"

func chessboard(size int) {
	for i := 0; i < size; i++ {
		chetStr := i%2 == 0
		for j := 0; j < size; j++ {
			chetCol := j%2 == 0
			if (chetStr && chetCol) || (!chetStr && !chetCol) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}

func main() {
	var size int
	fmt.Scanf("%d*%d", &size)
	chessboard(size)
}
