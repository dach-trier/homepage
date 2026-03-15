package main

import (
	"fmt"
	"os"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

func main() {
	res := esbuild.Build(
		esbuild.BuildOptions{
			EntryPoints: []string{},
			Bundle:      true,
			Write:       true,
			Outdir:      "web/js",
		},
	)

	if len(res.Warnings) > 0 {
		report(res.Warnings)
	}

	if len(res.Errors) > 0 {
		report(res.Errors)
		os.Exit(1)
	}
}

func report(messages []esbuild.Message) {
	options := esbuild.FormatMessagesOptions{Color: true}
	formatted := esbuild.FormatMessages(messages, options)
	fmt.Printf("%s", strings.Join(formatted, "\n"))
}
