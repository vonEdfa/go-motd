package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"

	pretty "github.com/vonEdfa/go-pretty-strings"
)

const appName = "GoMOTD"
const appVersion = "v1.0.0"
const appAuthor = "Malin \"vonEdfa\" von Matern"

// Feel free to add your name here when contributing :3
var appContributors = []string{""}

var currentOs string
var currentEnv []string
var opts flagOpts
var term termWindow

func init() {
	// Check os type and get some env info
	currentOs = runtime.GOOS
	currentEnv = os.Environ()

	opts.ParseAppFlags()
}

func main() {
	w := pretty.New()
	w.GetTermSize()
	fmt.Printf("%#v %#v\n", w.Terminal.Height, w.Terminal.Width)
	if !(strings.Contains(strings.ToLower(currentOs), "linux")) && opts.IgnoreOS == nil {
		fmt.Printf("Current OS seems to be %s. This script is intended for linux systems.\nUse the flag `-i` or `--ignore-os` to bypass this check in case of a false positive.\n\nExiting...\n", currentOs)
		os.Exit(0)
	}

	// Save term size
	term.GetSize()

	term.DrawBox("i", "gris")

	fmt.Printf("%#v\n", pretty.Padding("Apan ola\nVar en apa\n", "Sitta och snora", "Lorem ipsum"))
	fmt.Printf("%v\n", pretty.Center("Apan ola\nVar en apa\nSitta och snora\nLorem ipsum"))
	fmt.Printf("%#v %v %v %v\n", 9%2, float64(9)/2, math.Floor(float64(9)/2), math.Ceil(float64(9)/2))
}
