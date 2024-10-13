package main

import (
	"reflect"
	"testing"
)

func TestNormalizePhoneNumbers(t *testing.T) {
	got := NormalizePhoneNumbers([]string{
		"1234567890",
		"123 456 7891",
		"(123) 456 782",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	})
	want := []string{
		"1234567890",
		"1234567891",
		"123456782",
		"1234567893",
		"1234567894",
		"1234567890",
		"1234567892",
		"1234567892",
	}
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
