package main

import (
	"math/rand/v2"
)

type Direction string

const (
	L Direction = "L" // Swipe Left
	R           = "R" // Swipe Right
	U           = "U" // Swipe Up
	D           = "D" // Swipe Down
)

type Board [16]uint16

// Spawn randomly fills a random empty spot (denoted by '0').
// Probability: 2 (90% chance), 4 (10% chance).
// This is a no-op when there are no empty cells to fill.
func (b Board) Spawn() Board {
	// Determine all the empty spots
	var holes []int
	for idx, val := range b {
		if val == 0 {
			holes = append(holes, idx)
		}
	}

	if len(holes) == 0 {
		return b
	}

	// 90% likely to have 2 populated, 10% chance to have 4
	val := uint16(2)
	if rand.Float64() >= 0.9 {
		val = 4
	}

	idx := holes[rand.IntN(len(holes))]
	b[idx] = val

	return b
}

// Slide performs a swipe in the provided direction and returns
// the final board post merge.
//
// It also computes the score delta and returns a boolean if a
// swipe in the provided direction is not possible.
func (b Board) Slide(d Direction) (merged Board, score int, ok bool) {
	// If nothing has moved, it is an invalid move
	merged = b.merge(d)
	if merged == b {
		return merged, 0, false
	}

	// Score is the sum total of newly created cells
	counter := make(map[uint16]int)
	for i := range 16 {
		counter[merged[i]]++
		counter[b[i]]--
	}
	for cell, count := range counter {
		if cell != 0 && count > 0 {
			score += int(cell) * count
		}
	}

	return merged, score, true
}

// merge returns the board config after swiping in the given direction.
// To simplify things, we rotate the board based on input direction and
// always solve for swipe RTL (right to left).
func (b Board) merge(d Direction) Board {
	// Determine clockwise rotation count
	rotCW := 0
	switch d {
	case R:
		rotCW = 2
	case U:
		rotCW = 3
	case D:
		rotCW = 1
	}

	b = rotateCW(b, rotCW)

	// Solve one line at a time
	for i := range 4 {
		var line [4]uint16
		copy(line[:], b[i*4:(i+1)*4])
		mergedLine := mergeLine(line)
		copy(b[i*4:], mergedLine[:])
	}

	// Rotate back to how things were initially
	return rotateCW(b, 4-rotCW)
}

func mergeLine(line [4]uint16) [4]uint16 {
	for p1, p2 := 0, 1; p2 < 4; {
		switch {
		// If both are caught up, incr p2
		case p1 == p2:
			p2++

		// [0, 0, _, _], [2, 0, _, _] => incr right only
		case line[p2] == 0:
			p2++

		// [0, 2, _, _] => fill the hole, incr right only
		case line[p1] == 0:
			line[p1], line[p2] = line[p2], line[p1]
			p2++

		// [2, 2, _, _] => merge into left, incr both ptrs
		case line[p1] == line[p2]:
			line[p1], line[p2] = line[p1]*2, 0
			p1++
			p2++

		// [2, 4, _, _], [2, 0, _, _] => incr left until right
		default:
			p1++
		}
	}

	return line
}

// Rotate the given board clockwise for specified no of times.
func rotateCW(b Board, count int) Board {
	// helper to return 1D indexing from 2D based indexing.
	ij2i := func(i, j int) int { return i*4 + j }

	// Rotating beyond 3 times we get back to where we started
	count %= 4
	if count == 0 {
		return b
	}

	// Can simplify with something hardcoded, but this just reads better
	for range count {
		// Transpose
		for i := range 4 {
			for j := i; j < 4; j++ {
				p1, p2 := ij2i(i, j), ij2i(j, i)
				b[p1], b[p2] = b[p2], b[p1]
			}
		}

		// Reverse rows
		for i := range 4 {
			for start, end := 0, 3; start < end; start, end = start+1, end-1 {
				p1, p2 := ij2i(i, start), ij2i(i, end)
				b[p1], b[p2] = b[p2], b[p1]
			}
		}
	}

	return b
}
