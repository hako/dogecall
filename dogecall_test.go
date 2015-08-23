package main

import (
	"testing"
)

func TestCheckPhoneNumber(t *testing.T) {
	numbers := []string{
		"07700900390",
		"+447700900497",
		"202-555-0188",
		"+1-202-555-0188",
	}

	for _, number := range numbers {
		check := CheckNumber(number)
		if check != true {
			t.Errorf("CheckNumber(%s) == %s, want %s", number, check, nil)
		}
	}
}
