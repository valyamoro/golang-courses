package main 

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a  bool")
		case int:
			fmt.Println("Im an int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("heu")

	switch num := 6; num % 2 == 0 {
	case true:
		fmt.Print("true")
	case false:
		fmt.Print("false")
	}
}
