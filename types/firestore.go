package types

import (
	"time"
)

type TagosRecord struct {
	Sender string
	Receiver string
	Year int
	Month int
	Date int
	CreatedAt time.Time
}

func TagosRecordInit(sender, receiver string, year, month, date int) *TagosRecord {
	return &TagosRecord{
		Sender: sender,
		Receiver: receiver,
		Year: year,
		Month: month,
		Date: date,
		CreatedAt: time.Now(),
	}
}