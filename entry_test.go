package main

import (
	"testing"
	"time"
)

func Test_IsExpired(t *testing.T) {
	entry1 := Entry{Name: "test", Date: time.Now().AddDate(0, 0, -5)}
	if !entry1.IsExpired(4) {
		t.Error("entry1 should be expired")
	}

	entry2 := Entry{Name: "test", Date: time.Now().AddDate(0, 0, -5)}
	if entry2.IsExpired(6) {
		t.Error("entry1 should not be expired")
	}
}
