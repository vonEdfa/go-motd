package main

type flagOpts struct {
	Version       []bool `short:"v" long:"version"        description:"Print the current version number and exit"`
	IgnoreOS      []bool `short:"i" long:"ignore-os"      description:"Ignore OS checks at startup"`
	SquareBorders []bool `short:"q" long:"square-borders" description:"Use square edges instead of rounded for the borders. (Good workaround if your terminal font lacks the rounded corner symbols)"`
}

type termWindow struct {
	Width  int
	Height int
}
