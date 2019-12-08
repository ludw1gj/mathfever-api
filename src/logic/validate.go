package logic

import (
	"errors"
	"fmt"
	"strconv"
)

const max = 999999999999

func validateBinary(binary string) error {
	b, err := strconv.ParseInt(binary, 2, 0)
	if err != nil || b < 1 || len(binary) > 64 {
		return errors.New("Please enter a valid binary number, greater than 0 and is not over 64 " +
			"characters in length.")
	}
	return nil
}

func validateInt(positiveOnly bool, numbers ...int) error {
	const msg = "Please enter a valid whole number, "
	for _, n := range numbers {
		if positiveOnly && n < 1 {
			return fmt.Errorf("%vmust be greater than 0.", msg)
		}
		if n > max {
			return fmt.Errorf("%vmust not be greater than 999,999,999,999.", msg)
		}
	}
	return nil
}

func validateFloat(positiveOnly bool, floats ...float64) error {
	const msg = "Please enter a valid decimal number, "
	for _, n := range floats {
		if positiveOnly && n < 1 {
			return fmt.Errorf("%vmust be greater than 0.", msg)
		}
		if n > max {
			return fmt.Errorf("%vmust not be greater than 999,999,999,999.", msg)
		}
	}
	return nil
}

func validateHexadecimal(hexadecimal string) error {
	h, err := strconv.ParseInt(hexadecimal, 16, 0)
	if err != nil || h < 1 || len(hexadecimal) > 64 {
		return errors.New("Please enter a valid hexadecimal number, greater than 0 and is not over " +
			"999,999,999,999.")
	}
	return nil
}
