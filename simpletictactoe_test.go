package main

import (
	"testing"
)

func TestCheckForWin1(t *testing.T) {
	board := [][]string{
		{"X", "_", "_"},
		{"X", "_", "_"},
		{"X", "_", "_"},
	}
	isWinning := CheckForWin("X", board)
	if !isWinning {
		t.Fatalf(`isWinning should be true`)
	}
}

func TestCheckForWin2(t *testing.T) {
	board := [][]string{
		{"X", "_", "_"},
		{"X", "_", "_"},
		{"X", "_", "_"},
	}
	isWinning := CheckForWin("X", board)
	if !isWinning {
		t.Fatalf(`isWinning should be true`)
	}
}

func TestCheckForWin3(t *testing.T) {
	board := [][]string{
		{"O", "_", "_"},
		{"X", "_", "_"},
		{"X", "_", "_"},
	}
	isWinning := CheckForWin("X", board)
	if isWinning {
		t.Fatalf(`isWinning should be false`)
	}
}

func TestCheckForWin4(t *testing.T) {
	board := [][]string{
		{"O", "O", "X"},
		{"O", "X", "X"},
		{"X", "X", "O"},
	}
	isWinning := CheckForWin("X", board)
	if !isWinning {
		t.Fatalf(`isWinning should be true`)
	}
}

func TestCheckForWin5(t *testing.T) {
	board := [][]string{
		{"O", "O", "X"},
		{"O", "X", "X"},
		{"X", "X", "O"},
	}
	isWinning := CheckForWin("O", board)
	if isWinning {
		t.Fatalf(`isWinning should be false`)
	}
}

func TestCheckForWin6(t *testing.T) {
	board := [][]string{
		{"O", "O", "X"},
		{"O", "O", "X"},
		{"X", "X", "O"},
	}
	isWinning := CheckForWin("O", board)
	if !isWinning {
		t.Fatalf(`isWinning should be true`)
	}
}
