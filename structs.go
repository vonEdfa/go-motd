package main

type flagOpts struct {
	Version  []bool `short:"v" long:"version"   description:"Print the current version number and exit"`
	IgnoreOS []bool	`short:"i" long:"ignore-os" description:"Ignore OS checks at startup"`
}

type termWindow struct {
	Width	int
	Height	int
}