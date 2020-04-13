package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

var configPath *string
var wholeDir *bool
var markdownFile *string

func cli() {
	fs := flag.NewFlagSet("cblog", flag.ExitOnError)

	configPath = fs.String("config", "cblog.toml", "The configure file path of cblog")
	wholeDir = fs.Bool("whole-dir", false, "Build whole dir markdown file")
	markdownFile = fs.String("markdown", "", "The markdown file wants to build")

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