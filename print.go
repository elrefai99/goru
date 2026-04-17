package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintResponse(body JSONData) {
	format := "%v\t%v\t%v\t\n" // match 3 columns

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Repository", "Stars", "Description")
	fmt.Fprintf(tw, format, "----------", "-----", "----------")

	for _, i := range body.Items {
		desc := i.Description
		if len(desc) > 50 {
			desc = string(desc[:50]) + "..."
		}

		fmt.Fprintf(tw, format, i.FullName, i.StargazersCount, desc)
	}
	tw.Flush()
}
