package validator

import (
	"testing"
)

func TestStringMaxLength(t *testing.T) {
	dataProvider := []struct {
		text         string
		min          int
		max          int
		errorMessage string
	}{
		{text: "", min: 1, max: 3, errorMessage: "value should have between 1 and 3 characters"},
		{text: "aaaa", min: -1, max: 3, errorMessage: "value should have 3 characters or less"},
		{text: "", min: 1, max: -1, errorMessage: "value should have 1 characters or more"},
		{text: "", min: 2, max: 2, errorMessage: "value should have 2 characters"},
	}

	for _, test := range dataProvider {
		test := test // scope lint
		t.Run(test.errorMessage, func(t *testing.T) {
			err := Length(test.text, test.min, test.max)
			if test.errorMessage != "" {
				if err == nil {
					t.Error("should not be nil")
				}
				if err.Error() != test.errorMessage {
					t.Error("error message diff")
				}
			} else {
				if err != nil {
					t.Error("should be nil")
				}
			}
		})
	}

}
