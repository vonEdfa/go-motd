package main

import (
	"fmt"
	"strings"
	"os"
	"runtime"
)

const appName = "GoMOTD"
const appVersion = "v1.0.0"
const appAuthor = "vonEdfa"

var currentOs string
var currentEnv []string
var opts flagOpts
var term termWindow

func init() {
	// Check os type and get some env info
	currentOs = runtime.GOOS 
	currentEnv = os.Environ()

	opts.parseAppFlags()
}

func main() {
	if !(strings.Contains(strings.ToLower(currentOs), "linux")) && opts.IgnoreOS == nil {
		fmt.Printf("Current OS seems to be %s. This script is intended for linux systems.\nUse the flag `-i` or `--ignore-os` to bypass this check in case of a false positive.\nExiting...\n", currentOs)
		os.Exit(0)
	}

	// Save term size
	term.getSize()

	drawBox("")
}