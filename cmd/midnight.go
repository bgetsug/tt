// Copyright © 2018 Benjamin Getsug
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

var (
	days int
)

// midnightCmd represents the midnight command
var midnightCmd = &cobra.Command{
	Use:   "midnight",
	Short: "List upcoming midnights for each UTC offset",
	Run: midnight,
}

func init() {
	RootCmd.AddCommand(midnightCmd)

	flags := midnightCmd.PersistentFlags()

	flags.IntVarP(&days, "days", "d", 1, "Number of days to generate")
}

func midnight(cmd *cobra.Command, args []string) {
	var midnights []string

	for i := 0; i < days; i++ {
		t := time.Now().UTC()

		for _, loc := range utcOffsetLocations {
			midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, i)

			if midnight.After(t) {
				midnights = append(midnights, fmt.Sprint(
					midnight.In(time.Local), "  |  ", midnight.Format("2006-01-02 15:04:05 -0700 (MST)"),
				))
			}
		}
	}

	sort.Strings(midnights)

	fmt.Print("\nNext midnight will occur at:\n\n")
	fmt.Println("          Local Time           |               Zone Time               ")
	fmt.Println("-----------------------------------------------------------------------")

	for _, midnight := range midnights {
		fmt.Println(midnight)
	}
}