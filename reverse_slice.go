package reverseslice

import (
	"errors"
)

func ReverseSlice(slice []int) error {
	if slice == nil {
		return errors.New("the slice passed nil")
	}
	LEN := len(slice)
	if LEN == 0 {
		return errors.New("the slice passed is empty")
	}
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
