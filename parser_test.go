package main

import (
	"testing"
	"time"
)

func Test_ParseString(t *testing.T) {
	parser := NewParser("localhost.#date#.log")
	entry, _ := parser.ParseString("localhost.2023-01-01.log")
	if entry.Date.Equal(time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)) {
		t.Error("Failed to parse string")
	}
	entry, _ = parser.ParseString("localhost.20230101.log")
	if entry.Date.Equal(time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)) {
		t.Error("Failed to parse string")
	}
}
