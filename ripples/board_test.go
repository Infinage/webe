package main

import (
	"testing"
)

func Test_NewBoard(t *testing.T) {
	t.Run("Invalid inputs", func(t *testing.T) {
		tests := []struct {
			name       string
			rows, cols uint
			damping    float32
		}{
			{name: "Zero valued rows/cols", rows: 0, cols: 0, damping: 0},
			{name: "Damping greater than 1", rows: 1, cols: 1, damping: 2},
			{name: "Damping lesser than 0", rows: 1, cols: 1, damping: -1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if _, err := NewBoard(tt.rows, tt.cols, tt.damping); err == nil {
					t.Errorf("Expected failure, but passed: (rows: %d, cols: %d, damping: %f)",
						tt.rows, tt.cols, tt.damping)
				}
			})
		}
	})

	t.Run("Valid inputs", func(t *testing.T) {
		tests := []struct {
			name       string
			rows, cols uint
			damping    float32
		}{
			{name: "Smallest board", rows: 1, cols: 1, damping: 0},
			{name: "Random size (rows == cols)", rows: 10, cols: 10, damping: 1},
			{name: "Random size (rows != cols)", rows: 10, cols: 20, damping: 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				board, err := NewBoard(tt.rows, tt.cols, tt.damping)
				if err != nil {
					t.Fatalf("Unexpected error with valid inputs: %v", err)
				}

				crows, ccols := len(board.curr), len(board.curr[0])
				prows, pcols := len(board.prev), len(board.prev[0])

				if board.rows != tt.rows || board.cols != tt.cols ||
					crows != int(tt.rows) || ccols != int(tt.cols) ||
					prows != int(tt.rows) || pcols != int(tt.cols) {
					t.Fatalf("Board dims mismatch:"+
						"\nExpected: (%d, %d)"+
						"\nMeta:     (%d, %d)"+
						"\nCurr:     (%d, %d)"+
						"\nPrev:     (%d, %d)",
						tt.rows, tt.cols, board.rows, board.cols, crows, ccols, prows, pcols)
				}
				for x := range tt.rows {
					for y := range tt.cols {
						if board.curr[x][y] != 0 || board.prev[x][y] != 0 {
							t.Errorf("Board is not empty at (%d, %d)", x, y)
						}
					}
				}
			})
		}

	})
}

func Test_Click(t *testing.T) {
	b, _ := NewBoard(5, 5, 0.9)

	t.Run("Valid coordinate", func(t *testing.T) {
		if !b.Click(2, 2) {
			t.Errorf("Expected true for valid click")
		} else if b.curr[2][2] != 255 {
			t.Errorf("Expected clicked cell to be 255, got %f", b.curr[2][2])
		}
	})

	t.Run("Out of bounds", func(t *testing.T) {
		if b.Click(5, 5) {
			t.Errorf("Expected false for out of bounds click")
		}
	})

	t.Run("Corner click does not underflow", func(t *testing.T) {
		if !b.Click(0, 0) {
			t.Errorf("Expected true for corner click")
		} else if b.curr[0][0] != 255 {
			t.Errorf("Expected corner cell to be 255, got %f", b.curr[0][0])
		}
	})

	t.Run("Edge click paints neighbours within radius", func(t *testing.T) {
		b, _ := NewBoard(5, 5, 0.9)
		if !b.Click(0, 2) {
			t.Errorf("Expected true for edge click")
		}
		if b.curr[0][2] != 255 {
			t.Errorf("Expected center cell to be 255, got %f", b.curr[0][2])
		}
		if b.curr[1][2] == 0 {
			t.Errorf("Expected neighbouring cell within radius to be painted, got 0")
		}
	})
}

func Test_Resize(t *testing.T) {
	t.Run("Invalid dimensions", func(t *testing.T) {
		b, _ := NewBoard(5, 5, 0.9)
		if err := b.Resize(0, 5); err == nil {
			t.Errorf("Expected error for 0 rows")
		}
	})

	t.Run("Scale up", func(t *testing.T) {
		b, _ := NewBoard(3, 3, 0.9)
		b.curr[1][1] = 100

		err := b.Resize(5, 5)
		if err != nil {
			t.Errorf("Unexpected error scaling up: %v", err)
		} else if b.rows != 5 || b.cols != 5 {
			t.Errorf("Expected dimensions to be 5x5, got %dx%d", b.rows, b.cols)
		} else if b.curr[1][1] != 100 {
			t.Errorf("Expected data to be preserved, got %f", b.curr[1][1])
		}
	})

	t.Run("Scale down", func(t *testing.T) {
		b, _ := NewBoard(5, 5, 0.9)
		b.curr[1][1] = 100
		b.curr[4][4] = 200

		err := b.Resize(3, 3)
		if err != nil {
			t.Errorf("Unexpected error scaling down: %v", err)
		} else if b.rows != 3 || b.cols != 3 {
			t.Errorf("Expected dimensions to be 3x3, got %dx%d", b.rows, b.cols)
		} else if b.curr[1][1] != 100 {
			t.Errorf("Expected inner data to be preserved, got %f", b.curr[1][1])
		}
	})
}

func Test_Step(t *testing.T) {
	b, _ := NewBoard(3, 3, 1.0)
	b.curr[0][1] = 100
	res := b.Step()

	if b.prev[0][1] != 100 {
		t.Errorf("Expected prev buffer to hold the old curr state")
	}
	if res[1][1] != 50 {
		t.Errorf("Expected center cell to be 50, got %f", res[1][1])
	}
}
