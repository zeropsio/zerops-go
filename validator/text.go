package validator

import (
	"fmt"
	"unicode/utf8"
)

func Length(s string, min, max int) error {
	length := utf8.RuneCountInString(s)

	checkMin := min != -1
	checkMax := max != -1

	if (checkMin && length < min) || (checkMax && length > max) {
		if checkMax && checkMin {
			if max == min {
				return fmt.Errorf("value should have %d characters", min)
			} else {
				return fmt.Errorf("value should have between %d and %d characters", min, max)
			}
		}
		if checkMax {
			return fmt.Errorf("value should have %d characters or less", max)
		}
		if checkMin {
			return fmt.Errorf("value should have %d characters or more", min)
		}
	}

	return nil
}

func LengthBytes(s []byte, min, max int) error {
	length := len(s)

	checkMin := min != -1
	checkMax := max != -1

	if (checkMin && length < min) || (checkMax && length > max) {
		if checkMax && checkMin {
			if max == min {
				return fmt.Errorf("value should have %d bytes", min)
			} else {
				return fmt.Errorf("value should have between %d and %d bytes", min, max)
			}
		}
		if checkMax {
			return fmt.Errorf("value should have %d bytes or less", max)
		}
		if checkMin {
			return fmt.Errorf("value should have %d bytes or more", min)
		}
	}

	return nil
}
