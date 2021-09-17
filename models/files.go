package models

import "time"

type Files struct {
	Location string
	Size     int64
	Ext      string
	Name     string
	Time     time.Time
}
