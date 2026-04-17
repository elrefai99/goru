package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintResponse(items []Item) {
	format := "%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Repository", "Stars", "Description")
	fmt.Fprintf(tw, format, "----------", "-----", "-----------")

	for _, i := range items {
		desc := i.Description
		if len(desc) > 50 {
			desc = desc[:50] + "..."
		}
		fmt.Fprintf(tw, format, i.FullName, i.StargazersCount, desc)
	}
	tw.Flush()
}
