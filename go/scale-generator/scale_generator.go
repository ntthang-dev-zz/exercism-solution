//Package scale generates musical scales. See problem description for less than complete explanation.
package scale

import "strings"

type chromaticScale [12]string

var (
	chrSharp   = chromaticScale{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	chrFlat    = chromaticScale{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	chrLen     = len(chrSharp)
	flatTonics = "Fdgcf"
	intervals  = map[rune]int{
		'm': 1,
		'M': 2,
		'A': 3,
	}
)

// Scale returns the scale for the given tonic and interval
func Scale(tonic string, interval string) []string {
	if interval == "" {
		return buildChromaticScale(tonic, chrScale(tonic))
	}
	return buildScale(tonic, interval, chrScale(tonic))
}

func chrScale(tonic string) chromaticScale {
	if strings.Index(flatTonics, tonic) > -1 {
		return chrFlat
	}

	if len(tonic) == 1 {
		return chrSharp
	}

	if tonic[1] == 'b' {
		return chrFlat
	}
	return chrSharp
}

func findStart(tonic string, chr chromaticScale) int {
	u := strings.ToUpper(tonic[:1]) + tonic[1:]
	for i, v := range chr {
		if v == u {
			return i
		}
	}
	return -1
}

func buildChromaticScale(tonic string, chr chromaticScale) []string {
	c := findStart(tonic, chr)
	s := chr[c:]
	return append(s, chr[:c]...)
}

func buildScale(tonic string, interval string, chr chromaticScale) []string {
	s := make([]string, len(interval))
	c := findStart(tonic, chr)
	for i, r := range interval {
		s[i] = chr[c]
		c = (c + intervals[r]) % chrLen
	}
	return s
}
