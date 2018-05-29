package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/jessevdk/go-flags"
)

func (opts *flagOpts) ParseAppFlags() {
	if _, err := flags.Parse(opts); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	} else {
		if opts.Version != nil {
			contribs := sortAndJoinStrings(appContributors)

			fmt.Printf("%s version %s\nWritten by %s\n", appName, appVersion, appAuthor)

			if contribs != "" {
				fmt.Printf("\nCONTRIBUTORS:\n%s ♥\n", contribs)
			}

			os.Exit(0)
		}
	}
}

/**
 * -------------------
 * STRING MANIPULATORS
 * -------------------
 */

func sortAndJoinStrings(list []string) string {
	var output string
	var cleanList []string

	for _, s := range list {
		if s != "" {
			cleanList = append(cleanList, s)
		}
	}

	sort.Strings(cleanList)
	output = strings.Join(cleanList, ", ")

	return output
}

func getLongestStringLength(list []string) int {
	var longest int

	for _, s := range list {
		if len(s) > longest {
			longest = len(s)
		}
	}

	return longest
}

func stringPadding(s string, paddingLeft int, paddingRight int) string {
	var lp []string
	var rp []string

	for i := paddingLeft; i > 0; i-- {
		lp = append(lp, " ")
	}

	for i := paddingRight; i > 0; i-- {
		rp = append(rp, " ")
	}

	return strings.Join(lp, "") + s + strings.Join(rp, "")
}

/**
 * --------------
 * MISC FUNCTIONS
 * --------------
 */

func runCommand(cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Stdin = os.Stdin
	out, err := exe.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func horizontalBorder(maxWidth int, symbol string, padding int, corners bool) string {
	var border []string
	var corner = 0

	if corners {
		corner = 2
	}

	width := maxWidth - padding*2 - corner

	//debug
	fmt.Printf("My width is: %v\nMy symbol is: %s\n", width, symbol)

	for i := width; i > 0; i-- {
		border = append(border, symbol)
	}

	return strings.Join(border, "")
}

func renderBanner() {
	// TODO: Create a script to render the banner
}
