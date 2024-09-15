package main

import "fmt"

func TaskFifteen() []byte {
	justString := `ko qqq dqwdqw dwqqw dqwD 0913310391 DIQWDJ OQWdqw ldpa s`

	hugeStringByte := []byte(justString)
	justStringByte := hugeStringByte[:100]
	newSlice := make([]byte, len(justStringByte))
	copy(newSlice, hugeStringByte)
	fmt.Println(newSlice)
	
	return newSlice
}
