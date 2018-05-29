package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (window *termWindow) GetSize() {
	size := runCommand("stty", "size")
	size = strings.Trim(size, "\n")
	sizes := strings.Split(string(size), " ")
	h, err := strconv.ParseInt(sizes[0], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	w, err := strconv.ParseInt(sizes[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	window.Height = int(h)
	window.Width = int(w)
}

func (window *termWindow) DrawBox(content ...string) {
	var width int
	var margin int
	var padding int
	var borders map[string]string

	borders = map[string]string{
		"topL": "╭",
		"topR": "╮",
		"hor":  "─",
		"ver":  "│",
		"botL": "╰",
		"botR": "╯",
	}
	if opts.SquareBorders != nil {
		borders["topL"] = "┌"
		borders["topR"] = "┐"
		borders["botL"] = "└"
		borders["botR"] = "┘"
	}

	// Update term size
	window.GetSize()
	maxWidth := window.Width

	// Set margin and padding
	margin = 2
	padding = 2

	// Calculate width of content and use that if it's shorter than maxWidth
	contentWidth := getLongestStringLength(content) + padding*2 + margin*2 + 2

	if contentWidth < maxWidth {
		width = contentWidth
	} else {
		width = maxWidth
		contentWidth = maxWidth - padding*2 - margin*2 - 2
	}

	// Create our vertical borders
	verL := stringPadding(borders["ver"], margin, padding)
	verR := stringPadding(borders["ver"], padding, margin)

	// Create our horizontal border and print the top row
	horizontal := horizontalBorder(width, borders["hor"], margin, true)
	fmt.Println(stringPadding(borders["topL"]+horizontal+borders["topR"], margin, margin))

	// Print content surrounded by our vertical borders
	for _, s := range content {
		if len(s) <= getLongestStringLength(content) {
			fill := getLongestStringLength(content) - len(s)
			paddedString := stringPadding(s, 0, fill)
			fmt.Printf("%s%s%s\n", verL, paddedString, verR)
		} else {
			// TODO: Use word wrap! I downloaded https://github.com/bbrks/wrap for this.
			//wrap.Wrap(content, width)
		}
	}

	// Print bottom row
	fmt.Println(stringPadding(borders["botL"]+horizontal+borders["botR"], margin, margin))
}
