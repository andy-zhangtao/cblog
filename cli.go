package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

var configPath *string
var wholeDir *bool
var rebuild *bool
var markdownFile *string
var preview *bool
var port *int

func cli() {
	fs := flag.NewFlagSet("cblog", flag.ExitOnError)

	configPath = fs.String("config", "cblog.toml", "The configure file path of cblog")
	wholeDir = fs.Bool("whole-dir", false, "Build whole dir markdown file")
	rebuild = fs.Bool("rebuild", false, "Force rebuild current dir")
	markdownFile = fs.String("markdown", "", "The markdown file wants to build")
	preview = fs.Bool("preview", false, "If true, u can preview blog in localhost")
	port = fs.Int("port", 80, "Only valid in preview mode")

	fs.Usage = usageFor(fs)
	fs.Parse(os.Args[1:])
}

func usageFor(fs *flag.FlagSet) func() {
	short := os.Args[0] + " [flags]"
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
