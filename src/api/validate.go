package api

import (
	"fmt"
	"strconv"
)

const max = 999999999999

func validateBinary(binary string) error {
	b, err := strconv.ParseInt(binary, 2, 0)
	if err != nil || b < 1 || len(binary) > 64 {
		return fmt.Errorf("please enter a valid binary number, greater than 0 and is not over 64 "+
			"characters in length. Incorrect input: %s", binary)
	}
	return nil
}

func validatePositiveInt(numbers ...int) error {
	for _, n := range numbers {
		if n < 1 || n > max {
			return fmt.Errorf("please enter a valid decimal number, greater than 0 and less than "+
				"1'000'000'000'000. Incorrect input: %d", n)
		}
	}
	return nil
}

func validateFloat(positiveOnly bool, floats ...float64) error {
	for _, n := range floats {
		if positiveOnly && n < 1 {
			return fmt.Errorf("please enter a valid decimal number, greater than 0")
		}
		if n > max {
			return fmt.Errorf("please enter a valid decimal number, must be less than "+
				"1'000'000'000'000: %f", n)
		}
	}
	return nil
}

func validateHexadecimal(hexadecimal string) error {
	h, err := strconv.ParseInt(hexadecimal, 16, 0)
	if err != nil || h < 1 || len(hexadecimal) > 64 {
		return fmt.Errorf("please enter a valid hexadecimal number, greater than 0 and less than "+
			"1'000'000'000'000: %s", hexadecimal)
	}
	return nil
}
