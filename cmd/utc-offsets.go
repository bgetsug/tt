package cmd

import (
	"time"

	"github.com/jinzhu/now"
)

var (
	// Reference: https://en.wikipedia.org/wiki/List_of_UTC_time_offsets
	utcOffsets = []utcOffset{
		{"UTC-12:00", "Y", -12 * time.Hour},                  // 0
		{"UTC-11:00", "X", -11 * time.Hour},                  // 1
		{"UTC-10:00", "W", -10 * time.Hour},                  // 2
		{"UTC-09:30", "V†", -(9*time.Hour + 30*time.Minute)}, // 3
		{"UTC-09:00", "V", -9 * time.Hour},                   // 4
		{"UTC-08:00", "U", -8 * time.Hour},                   // 5
		{"UTC-07:00", "T", -7 * time.Hour},                   // 6
		{"UTC-06:00", "S", -6 * time.Hour},                   // 7
		{"UTC-05:00", "R", -5 * time.Hour},                   // 8
		{"UTC-04:00", "Q", -4 * time.Hour},                   // 9
		{"UTC-03:30", "P†", -(3*time.Hour + 30*time.Minute)}, // 10
		{"UTC-03:00", "P", -3 * time.Hour},                   // 11
		{"UTC-02:00", "O", -2 * time.Hour},                   // 12
		{"UTC-01:00", "N", -1 * time.Hour},                   // 13
		{"UTC±00:00", "Z", 0},                                // 14
		{"UTC+01:00", "A", 1 * time.Hour},                    // 15
		{"UTC+02:00", "B", 2 * time.Hour},                    // 16
		{"UTC+03:00", "C", 3 * time.Hour},                    // 17
		{"UTC+03:30", "C†", 3*time.Hour + 30*time.Minute},    // 18
		{"UTC+04:00", "D", 4 * time.Hour},                    // 19
		{"UTC+04:30", "D†", 4*time.Hour + 30*time.Minute},    // 20
		{"UTC+05:00", "E", 5 * time.Hour},                    // 21
		{"UTC+05:30", "E†", 5*time.Hour + 30*time.Minute},    // 22
		{"UTC+05:45", "E*", 5*time.Hour + 45*time.Minute},    // 23
		{"UTC+06:00", "F", 6 * time.Hour},                    // 24
		{"UTC+06:30", "F†", 6*time.Hour + 30*time.Minute},    // 25
		{"UTC+07:00", "G", 7 * time.Hour},                    // 26
		{"UTC+08:00", "H", 8 * time.Hour},                    // 27
		{"UTC+08:30", "H†", 8*time.Hour + 30*time.Minute},    // 28
		{"UTC+08:45", "H*", 8*time.Hour + 45*time.Minute},    // 29
		{"UTC+09:00", "I", 9 * time.Hour},                    // 30
		{"UTC+09:45", "I†", 9*time.Hour + 30*time.Minute},    // 31
		{"UTC+10:00", "K", 10 * time.Hour},                   // 32
		{"UTC+10:30", "K†", 10*time.Hour + 30*time.Minute},   // 33
		{"UTC+11:00", "L", 11 * time.Hour},                   // 34
		{"UTC+12:00", "M", 12 * time.Hour},                   // 35
		{"UTC+12:45", "M*", 12*time.Hour + 45*time.Minute},   // 36
		{"UTC+13:00", "M†", 13 * time.Hour},                  // 37
		{"UTC+14:00", "M†", 14 * time.Hour},                  // 38
	}

	utcOffsetLocations []*time.Location
)

func init() {
	for _, utcOffset := range utcOffsets {
		utcOffsetLocations = append(utcOffsetLocations, time.FixedZone(utcOffset.name, int(utcOffset.offset.Seconds())))
	}
}

type utcOffset struct {
	name         string
	nauticalName string
	offset       time.Duration
}

func (u utcOffset) Name() string {
	return u.name
}

func (u utcOffset) NauticalName() string {
	return u.nauticalName
}

func (u utcOffset) Offset() time.Duration {
	return u.offset
}

func currentMidnights(t time.Time) []time.Time {
	var matchingMidnights []time.Time

	for _, loc := range utcOffsetLocations {
		if midnight := now.New(t.In(loc)).BeginningOfDay(); midnight.Equal(t) {
			matchingMidnights = append(matchingMidnights, midnight)
		}
	}

	return matchingMidnights
}
