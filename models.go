package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
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

// SetTitle ...
func (n *Note) SetTitle(title string) {
	n.Title = title
	n.UpdatedAt = time.Now()
}

// SaveBody ...
func (n *Note) SaveBody(root string, body []byte) error {
	filename := path.Join(
		root, fmt.Sprintf("%s.md", n.CreatedAt.Format("2006-01-02-150405")),
	)

	return ioutil.WriteFile(filename, body, 0600)
}

// LoadBody ...
func (n *Note) LoadBody(root string) ([]byte, error) {
	filename := path.Join(
		root, fmt.Sprintf("%s.md", n.CreatedAt.Format("2006-01-02-150405")),
	)

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DeleteBody ...
func (n *Note) DeleteBody(root string) error {
	filename := path.Join(
		root, fmt.Sprintf("%s.md", n.CreatedAt.Format("2006-01-02-150405")),
	)

	return os.Remove(filename)
}
