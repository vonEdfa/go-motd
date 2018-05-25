package main

import (
	"strconv"
	"strings"
	"log"
	"fmt"
	"os"
	"os/exec"

	"github.com/jessevdk/go-flags"
)

func (opts *flagOpts) parseAppFlags() {
	if _, err := flags.Parse(opts); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	} else {
		if opts.Version != nil {
			fmt.Printf("%s version %s\nWritten by %s\n", appName, appVersion, appAuthor)
			os.Exit(0)
		}
	}
}

func runCommand(cmd string, args []string) string {
	exe := exec.Command(cmd, args...)
	exe.Stdin = os.Stdin
	out, err := exe.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func (window *termWindow) getSize() {
	size := runCommand("stty", []string{"size"})
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

func drawBox(content ...string) {
	borders := map[string]string{
		"topL": "╭",
		"topR": "╮",
		"ver": "─",
		"hor": "│",
		"botL": "╰",
		"botR": "╯",
	}
	fmt.Printf("%s%s%s\n", "\u2570", borders["ver"], borders["topR"])
}

func renderBanner() {
	// TODO: Create a script to render the banner
}