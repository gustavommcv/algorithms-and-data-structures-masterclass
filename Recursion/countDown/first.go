package main

import "fmt"

func main() {
	countDownRecursive(5)
}

func countDownRecursive(n int) {
	if n == 0 {
		fmt.Println("All done!")
		return
	}

	fmt.Println(n)
	countDownRecursive(n - 1)
}
