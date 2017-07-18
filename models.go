package main

import (
	"time"
)

// Note ...
type Note struct {
	ID        int       `storm:"id,increment"`
	Title     string    `storm:"index"`
	CreatedAt time.Time `storm:"index"`
	UpdatedAt time.Time `storm:"index"`
}

func NewNote(title string) *Note {
	return &Note{
		Title:     title,
		CreatedAt: time.Now(),
	}
}

func (t *Note) SetTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}
