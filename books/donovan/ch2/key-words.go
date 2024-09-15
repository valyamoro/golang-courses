package main

import "fmt"

func Break() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}

func Case() {
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week!")
	}
}

func Chan() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}

func Const() {
	const Pi = 3.14
	fmt.Println(Pi)
}

func Default() {
	defer fmt.Println("Done!")
	fmt.Println("Doing work...")
}

func Else() {
	x := 5
	if x > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func FallThrough() {
	switch x := 1; x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
}

func For() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func Func() {
	s := func(a int, b int) int {
		return a + b
	}

	fmt.Println(s(2, 3))
}

func Go() {
	go fmt.Println("This runs concurrently")
}

func Goto() {
	i := 0
	goto Label
	fmt.Println("Это мы пропустим...")
Label:
	fmt.Println(i)
}

func If() {
	if x := 5; x > 3 {
		fmt.Println("Greater")
	}
}

// import "fmt" - Импортируем пакет.

type Stringer interface {
	String() string
}

func Map() {
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m["key"])
}

// package main - Инициализирует текущий пакет

func Range() {
	arr := []int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Return() {
	s := func() int {
		return 42
	}

	fmt.Println(s())
}

func Select() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() { ch1 <- 1 }()
	go func() { ch2 <- 1 }()

	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	}
}

type Point struct {
	X, Y int
}

func Struct() {
	p := Point{X: 10, Y: 20}
	fmt.Println(p)
}

func Switch() {
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}
}

func Type() {
	type MyInt int
	var x MyInt = 10
	fmt.Println(x)
}

func Var() {
	var x int = 5
	fmt.Println(x)
}

//------------------------------------------------

func predefinedWords() {
	var isTrue bool = true
	var isFalse bool = false
	var b byte = 'A'
	var c64 complex64 = complex(1, 2)
	var c128 complex128 = complex(3, 4)
	var err error = fmt.Errorf("an error occurred")
	var f float64 = 3.14
	var f32 float32 = 3.14
	var i int = 42
	var i8 int8 = 42
	var i16 int16 = 42
	var i32 int32 = 42
	var i64 int64 = 42
	var r rune = '☺'
	var s string = "Hello, Go!"
	var u uint = 10
	var u8 uint8 = 10
	var u16 uint16 = 10
	var u32 uint32 = 10
	var u64 uint64 = 10
	var p uintptr = 0x1234
	const (
		First = iota
		Second
		Third
	)
	var s *int
	if s == nil {
		fmt.Println("S == nil")
	}

	m := make(map[string]int)
	length := len([]int{1, 2, 3})
	cap := cap(make([]int, 3, 4))
	p := new(int)
	arr := append([]int{1, 2}, 3, 4)
	{
		src := []int{1, 2, 3}
		dst := make([]int, len(src))
		copy(dst, src)
	}
	{
		ch := make(chan int)
		go func() {
			ch <- 42
			close(ch)
		}()
	}
	{
		m := map[string]int{"k1": 1, "k2": 2}
		delete(m, "k1")
	}
	imag := imag(complex(2, 3))
	{
		s := func(a, b int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered", r)
				}
			}()
			fmt.Println(a / b)
		}

		s(10, 0)
	}
}
