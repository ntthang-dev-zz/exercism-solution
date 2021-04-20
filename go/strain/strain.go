// Package strain filters a slice of ints, slices, strings
package strain

// Ints defines a collection of int values - slice of ints
type Ints []int

// List defines a collection of arrays of ints - slice of int slices
type Lists [][]int

// Strings  defines a collection of strings - slice of strings
type Strings []string

// Keep filters a collection of ints to only contain the members where the provided function returns true.
func (in Ints) Keep(strainer func(int) bool) (out Ints) {
	for _, v := range in {
		if strainer(v) {
			out = append(out, v)
		}
	}
	return
}

// Discard filters a collection to only contain the members where the provided function returns false.
func (in Ints) Discard(strainer func(int) bool) Ints {
	return in.Keep(func(n int) bool { return !(strainer(n)) })
}

// Keep filters a collection of lists to only contain the members where the provided function returns true.
func (list Lists) Keep(strainer func([]int) bool) (out Lists) {
	for _, v := range list {
		if strainer(v) {
			out = append(out, v)
		}
	}
	return
}

// Keep filters a collection of strings to only contain the members where the provided function returns true.
func (s Strings) Keep(strainer func(string) bool) (out Strings) {
	for _, v := range s {
		if strainer(v) {
			out = append(out, v)
		}
	}
	return
}
