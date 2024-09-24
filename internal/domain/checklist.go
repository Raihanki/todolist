package domain

import "time"

type Checklist struct {
	ID        int
	UserId    int
	Title     string
	CreatedAt time.Time
}
