package main

import "fmt"

func chessboard(size1 int, size2 int) {
	for i := 0; i < size1; i++ {
		chetstr := i%2 == 0
		for j := 0; j < size2; j++ {
			chetstlb := j%2 == 0
			if chetstr {
				if chetstlb {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			} else {
				if chetstlb {
					fmt.Print(" ")
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Println()
	}

}

func main() {

	var size1, size2 int
	n, err := fmt.Scanf("%d*%d", &size1, &size2)
	if err != nil {
		fmt.Println("Incorect", n, err)
	}
	chessboard(size1, size2)
}
