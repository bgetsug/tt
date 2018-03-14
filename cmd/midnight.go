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
	midnightCount int
)

// midnightCmd represents the midnight command
var midnightCmd = &cobra.Command{
	Use:   "midnight",
	Short: "List upcoming midnights for each UTC offset",
	Run:   midnight,
}

func init() {
	RootCmd.AddCommand(midnightCmd)

	flags := midnightCmd.PersistentFlags()

	flags.IntVarP(&midnightCount, "count", "n", 1, "Number of midnights to generate")
}

func midnight(cmd *cobra.Command, args []string) {
	t := time.Now().UTC()

	fmt.Print("\nNext midnight(s) will occur at:\n\n")
	fmt.Println("          Local Time           |               Zone Time               ")
	fmt.Println("-----------------------------------------------------------------------")

	midnights := make(chan time.Time)
	sentMidnights := 0

	go midnightsAfter(t, t, midnights, &sentMidnights)

	for midnight := range midnights {
		fmt.Println(fmt.Sprint(
			midnight.In(time.Local), "  |  ", midnight.Format("2006-01-02 15:04:05 -0700 (MST)"),
		))
	}
}

func midnightsAfter(initialTime, t time.Time, midnights chan time.Time, sentMidnights *int) {

	var mns []time.Time

	for _, loc := range utcOffsetLocations {
		midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)

		if midnight.After(initialTime) {
			mns = append(mns, midnight)
		}
	}

	sort.Slice(mns, func(i, j int) bool { return mns[i].Before(mns[j]) })

	for _, midnight := range mns {
		if *sentMidnights == midnightCount {
			break
		}

		midnights <- midnight
		*sentMidnights++
	}

	if *sentMidnights < midnightCount {
		go midnightsAfter(initialTime, t.AddDate(0, 0, 1), midnights, sentMidnights)
		return
	}

	close(midnights)
}
