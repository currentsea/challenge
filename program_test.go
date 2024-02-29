package main

import (
	"testing"
)

func TestProgramCase1(t *testing.T) {
	input := "James Bond 7"
	padding := 3
	want := "James Bond 007"
	result := PadNumbers(input, padding)
	if result != want {
		t.Fatalf("Test failed for input: \":" + input + "\"")
	}
}

func TestProgramCase2(t *testing.T) {
	input := "PI=3.14"
	padding := 2
	want := "PI=03.14"
	result := PadNumbers(input, padding)
	if result != want {
		t.Fatalf("Test failed for input: \":" + input + "\"")
	}
}

func TestProgramCase3(t *testing.T) {
	input := "It's 3:13pm"
	padding := 2
	want := "It's 03:13pm"
	result := PadNumbers(input, padding)
	if result != want {
		t.Fatalf("Test failed for input: \":" + input + "\"")
	}
}

func TestProgramCase4(t *testing.T) {
	input := "It's 12:13pm"
	padding := 2
	want := "It's 12:13pm"
	result := PadNumbers(input, padding)
	if result != want {
		t.Fatalf("Test failed for input: \":" + input + "\"")
	}
}

func TestProgramCase5(t *testing.T) {
	input := "99UR1337"
	padding := 6
	want := "000099UR001337"
	result := PadNumbers(input, padding)
	if result != want {
		t.Fatalf("Test failed for input: \":" + input + "\"")
	}
}
