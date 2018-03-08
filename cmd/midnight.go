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

// midnightCmd represents the midnight command
var midnightCmd = &cobra.Command{
	Use:   "midnight",
	Short: "List upcoming midnights",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("\nNext midnight will occur at:")

		t := time.Now().UTC()

		var midnights []string

		for _, loc := range utcOffsetLocations {
			midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)

			if midnight.After(t) {
				midnights = append(midnights, fmt.Sprint(midnight.In(time.Local), "  |  ", loc.String()))
			}
		}

		sort.Strings(midnights)

		fmt.Println("\n          Local Time           | UTC Offset")
		fmt.Println("-------------------------------------------")

		for _, midnight := range midnights {
			fmt.Println(midnight)
		}

	},
}

func init() {
	RootCmd.AddCommand(midnightCmd)
}
