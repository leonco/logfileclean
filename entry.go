package main

import (
	"time"
)

type Entry struct {
	Name string
	Date time.Time
}

func (e Entry) String() string {
	return e.Name + " " + e.Date.Format("2006-01-02")
}
func (e Entry) IsExpired(d int) bool {
	keep := time.Now().AddDate(0, 0, -d)
	return e.Date.Before(keep)
}
