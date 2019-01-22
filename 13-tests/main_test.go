package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Add(42, 10)
	if result != 52 {
		t.Error("Result of 42 + 10 must be 52")
	}
}

// // Mostrando Error
// func TestAddFail(t *testing.T) {
// 	calculator := &Calculator{}
// 	result := calculator.Add(42, 10)
// 	if result != 50 {
// 		t.Error("Result of 42 + 10 must be 52")
// 	}
// }

// Desafio
func TestSubSuccess(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Sub(42, 10)
	if result != 32 {
		t.Error("Result of 42 - 10 must be 32")
	}
}

func TestMultSuccess(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Mult(42, 10)
	if result != 420 {
		t.Error("Result of 42 * 10 must be 420")
	}
}

func TestDivSuccess(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Div(42, 10)
	if result != 4 {
		t.Error("Result of 42 / 10 must be 4")
	}
}
