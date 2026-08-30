package main

import (
	"fmt"
	"math"
)

type Board struct {
	rows, cols uint
	curr, prev [][]float32
	damping    float32
}

// NewBoard initalizes a zero valued board
func NewBoard(rows, cols uint, damping float32) (Board, error) {
	if rows == 0 || cols == 0 {
		return Board{}, fmt.Errorf("invalid board dimensions: (%d, %d)", rows, cols)
	}

	if damping < 0 || damping > 1 {
		return Board{}, fmt.Errorf("invalid range for damping [0, 1]: %f", damping)
	}

	var curr, prev [][]float32
	for range rows {
		curr = append(curr, make([]float32, cols))
		prev = append(prev, make([]float32, cols))
	}

	return Board{
		rows: rows, cols: cols,
		curr: curr, prev: prev,
		damping: damping,
	}, nil
}

// Step executes one iteration of the water rippling algorithm.
func (b *Board) Step() [][]float32 {
	// Swap the buffers
	b.curr, b.prev = b.prev, b.curr

	// Each step executes on the boundary cells
	for x := uint(1); x < b.rows-1; x++ {
		for y := uint(1); y < b.cols-1; y++ {
			neighbours := b.prev[x-1][y] + b.prev[x+1][y] + b.prev[x][y+1] + b.prev[x][y-1]
			b.curr[x][y] = (neighbours / 2) - b.curr[x][y]
			b.curr[x][y] *= float32(b.damping)
		}
	}

	// print out the 'curr' state
	return b.curr
}

// Resize alters the board dimensions while retaining the board contents
func (b *Board) Resize(rows, cols uint) error {
	if rows == 0 || cols == 0 {
		return fmt.Errorf("invalid board dimensions: (%d, %d)", rows, cols)
	}

	// Init board with new dimensions (value: 0)
	var curr, prev [][]float32
	for range rows {
		curr = append(curr, make([]float32, cols))
		prev = append(prev, make([]float32, cols))
	}

	// Copy whatever value we can from original board
	for x := range min(rows, b.rows) {
		for y := range min(cols, b.cols) {
			prev[x][y] = b.prev[x][y]
			curr[x][y] = b.curr[x][y]
		}
	}

	b.rows, b.cols = rows, cols
	b.curr, b.prev = curr, prev

	return nil
}

func (b *Board) Click(x, y uint) bool {
	if x >= b.rows || y >= b.cols {
		return false
	}

	const radius = 2

	xStart := uint(0)
	if x > radius {
		xStart = x - radius
	}
	yStart := uint(0)
	if y > radius {
		yStart = y - radius
	}

	for xi := xStart; xi < min(b.rows, x+radius+1); xi++ {
		for yi := yStart; yi < min(b.cols, y+radius+1); yi++ {
			dx := float64(xi) - float64(x)
			dy := float64(yi) - float64(y)
			distance := math.Sqrt(dx*dx + dy*dy)
			if distance <= radius {
				factor := 1 / (1 + distance)
				b.curr[xi][yi] = float32(255 * factor)
			}
		}
	}

	return true
}
