package main

import (
	"errors"
	"fmt"
)

func reverseSlice(slice []int) error {
	if slice == nil {
		return errors.New("the array passed is empty")
	}
	LEN := len(slice)
	temp := make([]int, LEN)
	j := 0
	for i := LEN - 1; i >= 0; i-- {
		temp[j] = slice[i]
		j++
	}
	for i := 0; i < LEN; i++ {
		slice[i] = temp[i]
	}

	return nil
}
func main() {
	slice := []int{0, 1, 2}
	reverseSlice(slice)
	for i := range slice {
		fmt.Println(slice[i])
	}

}
