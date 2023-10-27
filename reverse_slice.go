package reverseslice

import "errors"

func ReverseSlice(slice []int) error {
	if slice == nil {
		return errors.New("the array passed is empty")
	}
	startingIndex := len(slice) - 1
	for i := startingIndex; i >= 0; i-- {
		for j := 0; i < len(slice); j++ {
			slice[j] = slice[i]
		}
	}
	return nil
}
