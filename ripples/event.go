package main

import "fmt"

type Event interface {
	Do(*Board) error
}

type EventClick struct {
	x, y uint
}

func (e *EventClick) Do(b *Board) error {
	if !b.Click(e.x, e.y) {
		return fmt.Errorf("click failed: (%d, %d)", e.x, e.y)
	}
	return nil
}

type EventResize struct {
	height, width uint
}

func (e *EventResize) Do(b *Board) error {
	return b.Resize(e.height, e.width)
}
