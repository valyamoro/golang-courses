package main 

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	j := 1
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Print(j)

	for {
		fmt.Println("loop")
		break 
	}

	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue 
		}

		fmt.Println(n)
	}

	for ; i < 10; i++ {
		fmt.Println(i * i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(j)
		}

		fmt.Println()
	}
}
