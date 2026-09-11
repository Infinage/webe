package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// renderBoard returns a html markup of the board with the provided animations.
func renderBoard(b Board, animations [16]string) string {
	var buf bytes.Buffer
	buf.WriteString(`
		<div id="board" 
			class="grid grid-cols-4 grid-rows-4 gap-3 p-3 bg-[#9C8B7C] w-full 
				max-w-md aspect-square rounded-xl select-none touch-none"
			data-on:touchstart="$_tCoords = [evt.changedTouches[0].clientX, evt.changedTouches[0].clientY]"
			data-on:touchmove__window="$_touching = true"
			data-on:touchend__window="
			  if (!$_touching) return;

			  const minSwipeDistance = 10;
			  const deltaX = evt.changedTouches[0].clientX - $_tCoords[0]; 
			  const deltaY = evt.changedTouches[0].clientY - $_tCoords[1];

			  // Horizontal swipe
			  let direction = '';
			  if (Math.abs(deltaX) > Math.abs(deltaY)) {
			    if (Math.abs(deltaX) > minSwipeDistance) {
				  direction = deltaX > 0 ? 'ArrowRight': 'ArrowLeft';
				}
			  } 

			  // Vertical swipe
			  else {
			    if (Math.abs(deltaY) > minSwipeDistance) {
				  direction = deltaY > 0 ? 'ArrowDown': 'ArrowUp';
				}
			  }

			  // Send the slide request
			  if (direction) {
				  @post('/api/slide', { payload: {
					 key: direction, 
					 boards: $boards, 
					 scores: $scores, 
					 hscore: $_hscore
				  }});
			  }

			  $_touching = false;
  			"
		>
	`)

	colors := map[uint16]string{
		0:    "bg-[#BDAC97] text-transparent",
		2:    "bg-[#EEE4DA] text-[#776E65]",
		4:    "bg-[#EDE0C8] text-[#776E65]",
		8:    "bg-[#F2B179] text-white",
		16:   "bg-[#F59563] text-white",
		32:   "bg-[#F67C5F] text-white",
		64:   "bg-[#F65E3B] text-white",
		128:  "bg-[#EDCF72] text-white",
		256:  "bg-[#EDCC61] text-white",
		512:  "bg-[#EDC850] text-white",
		1024: "bg-[#EDC53F] text-white",
		2048: "bg-[#EDC22E] text-white",
	}

	for idx, cell := range b {
		color, ok := colors[cell]
		if !ok {
			color = "text-white bg-[#3C3A32]"
		}

		text := strconv.FormatUint(uint64(cell), 10)

		// Make larger numbers slightly smaller
		fontSize := "text-4xl sm:text-5xl"
		if cell >= 10000 {
			fontSize = "text-2xl sm:text-3xl"
		} else if cell >= 1000 {
			fontSize = "text-3xl sm:text-4xl"
		}

		cellHtml := fmt.Sprintf(`
			<div class="rounded-xl flex items-center justify-center 
				font-bold select-none %s %s %s">%s</div>
		`, fontSize, color, animations[idx], text)

		buf.WriteString(cellHtml)
	}

	buf.WriteString("</div>")
	return buf.String()
}
