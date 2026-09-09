package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func renderBoard(b Board) string {
	var buf bytes.Buffer
	buf.WriteString(`<div id="board" class="grid grid-cols-4 grid-rows-4 gap-3 p-3 
		bg-[#9C8B7C] h-1/2 aspect-square rounded-xl"
	>`)

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

	for _, cell := range b {
		color, ok := colors[cell]
		if !ok {
			color = "text-white bg-[#3C3A32]"
		}

		text := strconv.FormatUint(uint64(cell), 10)

		// Make larger numbers slightly smaller
		fontSize := "text-4xl"
		if cell >= 1000 {
			fontSize = "text-3xl"
		}
		if cell >= 10000 {
			fontSize = "text-2xl"
		}

		cellHtml := fmt.Sprintf(`
			<div class="rounded-xl flex items-center justify-center 
				font-bold select-none %s %s">%s</div>
		`, fontSize, color, text)

		buf.WriteString(cellHtml)
	}

	buf.WriteString("</div>")
	return buf.String()
}
