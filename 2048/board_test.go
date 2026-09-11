package main

import (
	"testing"
)

// counter returns the frequency of numbers on the input board
func counter(t *testing.T, b Board) map[uint16]int {
	t.Helper()
	freq := make(map[uint16]int, 16)
	for _, v := range b {
		freq[v]++
	}
	return freq
}

func Test_NewBoard(t *testing.T) {
	b := NewBoard()
	if freq := counter(t, b); freq[0] != 14 || freq[2]+freq[4] != 2 {
		t.Errorf("Expected 2 cells to be filled: %v", freq)
	}
}

func Test_Spawn(t *testing.T) {
	var b Board
	initialFreq := counter(t, b)
	if len(initialFreq) != 1 || initialFreq[0] != 16 {
		t.Errorf("Expected to init empty board, got %v", b)
	}

	for i := range 16 {
		b = b.Spawn()
		t.Log(b)
		freq := counter(t, b)
		if freq[0] != 16-(i+1) || freq[2]+freq[4] != i+1 {
			t.Errorf("Expected iteration #%d to have %d/16 cells filled", i+1, freq[2]+freq[4])
		}
	}
}

func Test_Slide(t *testing.T) {
	type TC struct {
		before Board
		dir    Direction
		merged Board
		deltas Board
		score  int
		ok     bool
	}

	t.Run("Zero score change", func(t *testing.T) {
		tests := []TC{
			{
				before: Board{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
				merged: Board{0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0},
				deltas: Board{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0},
				dir:    DirectionLeft,
			},
			{
				before: Board{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
				merged: Board{0, 2, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				deltas: Board{0, 1, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				dir:    DirectionUp,
			},
			{
				before: Board{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
				merged: Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 2},
				deltas: Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
				dir:    DirectionDown,
			},
			{
				before: Board{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
				merged: Board{0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
				deltas: Board{0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
				dir:    DirectionRight,
			},
		}

		for _, tt := range tests {
			merged, deltas, score, ok := tt.before.Slide(tt.dir)
			if !ok || score != 0 || tt.merged != merged || tt.deltas != deltas {
				t.Errorf("\nInit: (%v, Dir: %v)\nWant: (%v, %v, 0, true)\n"+
					"Got : (%v, %v, %d, %t)", tt.before, tt.dir, tt.merged,
					tt.deltas, merged, deltas, score, ok)
			}
		}
	})

	t.Run("Fold near wall", func(t *testing.T) {
		tests := []TC{
			{
				before: Board{2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 2, 0, 0, 2, 2, 2},
				merged: Board{4, 2, 0, 0, 4, 2, 0, 0, 4, 2, 0, 0, 4, 2, 0, 0},
				deltas: Board{1, 1, 0, 0, 2, 2, 0, 0, 1, 1, 0, 0, 2, 2, 0, 0},
				dir:    DirectionLeft, score: 16,
			},
			{
				before: Board{2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 2, 0, 0, 2, 2, 2},
				merged: Board{0, 0, 2, 4, 0, 0, 2, 4, 0, 0, 2, 4, 0, 0, 2, 4},
				deltas: Board{0, 0, 2, 2, 0, 0, 1, 1, 0, 0, 2, 2, 0, 0, 1, 1},
				dir:    DirectionRight, score: 16,
			},
			{
				before: Board{2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 2, 0, 0, 2, 2, 2},
				merged: Board{4, 4, 4, 4, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				deltas: Board{2, 1, 1, 3, 0, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				dir:    DirectionUp, score: 24,
			},
			{
				before: Board{2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 2, 0, 0, 2, 2, 2},
				merged: Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 0, 4, 4, 4, 4},
				deltas: Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 3, 1, 1, 2},
				dir:    DirectionDown, score: 24,
			},
		}

		for _, tt := range tests {
			merged, deltas, score, ok := tt.before.Slide(tt.dir)
			if !ok || score != tt.score || merged != tt.merged || deltas != tt.deltas {
				t.Errorf("\nInit: (%v, Dir: %v)\nWant: (%v, %v, 0, true)\n"+
					"Got : (%v, %v, %d, %t)", tt.before, tt.dir, tt.merged,
					tt.deltas, merged, deltas, score, ok)
			}
		}
	})

	t.Run("Doesn't merge twice", func(t *testing.T) {
		tests := []TC{
			{
				before: Board{4, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				merged: Board{4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				deltas: Board{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				dir:    DirectionLeft, score: 4,
			},
			{
				before: Board{4, 0, 0, 0, 4, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0},
				merged: Board{0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 8, 0, 0, 0},
				deltas: Board{0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0},
				dir:    DirectionDown, score: 8,
			},
		}

		for _, tt := range tests {
			merged, deltas, score, ok := tt.before.Slide(tt.dir)
			if !ok || score != tt.score || merged != tt.merged || deltas != tt.deltas {
				t.Errorf("\nInit: (%v, Dir: %v)\nWant: (%v, %v, 0, true)\n"+
					"Got : (%v, %v, %d, %t)", tt.before, tt.dir, tt.merged,
					tt.deltas, merged, deltas, score, ok)
			}
		}
	})

	t.Run("Insolvable", func(t *testing.T) {
		for _, b := range []Board{
			{2, 4, 2, 4, 4, 2, 4, 2, 2, 4, 2, 4, 4, 2, 4, 2},
			{2, 8, 4, 2, 8, 2, 16, 8, 16, 32, 8, 2, 4, 8, 32, 8},
			{2, 4, 8, 4, 4, 2, 4, 16, 8, 4, 8, 64, 2, 16, 32, 8},
			{4, 8, 2, 4, 2, 4, 8, 16, 4, 8, 4, 2, 8, 2, 8, 4},
		} {
			for _, dir := range Directions {
				merged, _, score, ok := b.Slide(dir)
				if score != 0 || ok {
					t.Errorf("Expected slide [%v] to fail, got %v", dir, merged)
				}
			}
		}
	})
}

func Test_MergeLine(t *testing.T) {
	tests := []struct {
		input, mergedLine, deltaLine [4]uint16
	}{
		{
			input:      [4]uint16{0, 2, 0, 2},
			mergedLine: [4]uint16{4, 0, 0, 0},
			deltaLine:  [4]uint16{3, 0, 0, 0},
		},
		{
			input:      [4]uint16{0, 0, 0, 2},
			mergedLine: [4]uint16{2, 0, 0, 0},
			deltaLine:  [4]uint16{3, 0, 0, 0},
		},
		{
			input:      [4]uint16{0, 2, 2, 2},
			mergedLine: [4]uint16{4, 2, 0, 0},
			deltaLine:  [4]uint16{2, 2, 0, 0},
		},
		{
			input:      [4]uint16{2, 0, 0, 2},
			mergedLine: [4]uint16{4, 0, 0, 0},
			deltaLine:  [4]uint16{3, 0, 0, 0},
		},
		{
			input:      [4]uint16{2, 2, 2, 0},
			mergedLine: [4]uint16{4, 2, 0, 0},
			deltaLine:  [4]uint16{1, 1, 0, 0},
		},
		{
			input:      [4]uint16{8, 0, 2, 4},
			mergedLine: [4]uint16{8, 2, 4, 0},
			deltaLine:  [4]uint16{0, 1, 1, 0},
		},
		{
			input:      [4]uint16{2, 4, 2, 4},
			mergedLine: [4]uint16{2, 4, 2, 4},
			deltaLine:  [4]uint16{0, 0, 0, 0},
		},
		{
			input:      [4]uint16{2, 4, 4, 2},
			mergedLine: [4]uint16{2, 8, 2, 0},
			deltaLine:  [4]uint16{0, 1, 1, 0},
		},
		{
			input:      [4]uint16{2, 2, 2, 2},
			mergedLine: [4]uint16{4, 4, 0, 0},
			deltaLine:  [4]uint16{1, 2, 0, 0},
		},
	}

	for _, tt := range tests {
		if mLine, dLine := mergeLine(tt.input); mLine != tt.mergedLine {
			t.Errorf("mergeLine(%v) => want (%v, %v), got (%v, %v)",
				tt.input, tt.mergedLine, tt.deltaLine, mLine, dLine)
		}
	}
}

func Test_RotateCW(t *testing.T) {
	type TC struct {
		before Board
		after  Board
		count  int
	}

	tests := []TC{
		{
			before: Board{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			after:  Board{13, 9, 5, 1, 14, 10, 6, 2, 15, 11, 7, 3, 16, 12, 8, 4},
			count:  1,
		},
		{
			before: Board{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			after:  Board{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			count:  2,
		},
		{
			before: Board{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			after:  Board{4, 8, 12, 16, 3, 7, 11, 15, 2, 6, 10, 14, 1, 5, 9, 13},
			count:  3,
		},
	}

	for _, tt := range tests {
		for i := range 5 {
			count := tt.count + 4*i
			if got := rotateCW(tt.before, count); got != tt.after {
				t.Errorf("Rot(%d)\nInit: %v\nwant: %v\ngot : %v", count,
					tt.before, tt.after, got)
			}
		}
	}
}

func Test_Status(t *testing.T) {
	type TC struct {
		input  Board
		status GameStatus
	}

	tests := []TC{
		{input: Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, status: StatusLoss},
		{input: Board{2, 4, 2, 4, 4, 2, 4, 2, 2, 4, 2, 4, 4, 2, 4, 2}, status: StatusLoss},
		{input: Board{4, 8, 2, 4, 2, 4, 8, 16, 4, 8, 4, 2, 8, 2, 8, 4}, status: StatusLoss},
		{input: Board{0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}, status: StatusInProgress},
		{input: Board{0, 0, 2, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}, status: StatusInProgress},
		{input: Board{0, 0, 512, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 512, 0}, status: StatusInProgress},
		{input: Board{0, 512, 512, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2048, 0}, status: StatusWin},
		{input: Board{0, 2048, 512, 2, 0, 0, 2, 0, 2048, 0, 0, 0, 0, 0, 2048, 0}, status: StatusWin},
	}

	for _, tt := range tests {
		if st := tt.input.Status(); st != tt.status {
			t.Errorf("%v: want %q, got %q", tt.input, tt.status, st)
		}
	}
}
