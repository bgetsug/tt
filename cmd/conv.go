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
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	format   string
	timezone string

	tNow = time.Now()
	yr   = tNow.Year()
	mo   = tNow.Month()
	d    = tNow.Day()
	hr   = tNow.Hour()
	min  = tNow.Minute()
	sec  = tNow.Second()
	ns   = tNow.Nanosecond()
	loc  = time.Local
)

// convCmd represents the conv command
var convCmd = &cobra.Command{
	Use:   "conv [time | year] [month] [day] [hour] [minute] [second] [nanosecond] [location]",
	Short: "Convert a time",
	Run:   conv,
}

func init() {
	RootCmd.AddCommand(convCmd)

	flags := convCmd.PersistentFlags()

	flags.StringVarP(&format, "format", "f", "2006-01-02 15:04:05 -0700 MST", "Time format (layout)")
	flags.StringVarP(&timezone, "timezone", "z", time.Local.String(), "The output timezone")
}

func conv(cmd *cobra.Command, args []string) {
	loc := location(timezone)

	in := parseTimeFromArgs(args)

	fmt.Println(in.Format(format), "occurs at", in.In(loc).Format(format))
}

func location(name string) *time.Location {
	var err error

	loc := time.Local

	if len(name) > 0 {
		if strings.HasPrefix(name, "UTC") {
			for _, offset := range utcOffsetLocations {
				if name == offset.String() {
					return offset
				}
			}
		}

		switch name {
		case "CDT":
			name = "CST6CDT"
		case "EDT":
			name = "EST5EDT"
		case "MDT":
			name = "MST7MDT"
		case "PDT":
			name = "PST8PDT"
		}

		if loc, err = time.LoadLocation(name); err != nil {
			panic(err)
		}
	}

	return loc
}
