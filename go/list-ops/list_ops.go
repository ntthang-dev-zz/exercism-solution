/*
Package listops Implement a series of basic list operations, without using existing functions.

The precise number and names of the operations to be implemented will be
track dependent to avoid conflicts with existing names, but the general
operations you will implement include:

func (l IntList) Append(m IntList) IntList

func (l IntList) Concat(a []IntList) IntList

func (l IntList) Filter(p predFunc) IntList

func (l IntList) Length() int

func (l IntList) Map(u unaryFunc) IntList

func (l IntList) Foldl(f binFunc, acc int) int

func (l IntList) Foldr(f binFunc, acc int) int

func (l IntList) Reverse() IntList

*/
package listops

// IntList is a slice of integers.
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Append given two lists, add all items in the second list to the end of the first list
func (l IntList) Append(m IntList) IntList {
	f := make(IntList, 0, len(l)+len(m))
	f = append(f, l...)
	f = append(f, m...)
	return f
}

// Concat given a series of lists, combine all items in all lists into one flattened list
func (l IntList) Concat(a []IntList) IntList {
	f := make(IntList, 0, len(l)*len(a))
	f = append(f, l...)
	for _, m := range a {
		f = append(f, m...)
	}
	return f
}

// Filter given a predicate and a list, return the list of all items for which predicate(item) is True
func (l IntList) Filter(p predFunc) IntList {
	f := make(IntList, 0, len(l))
	for _, v := range l {
		if p(v) {
			f = append(f, v)
		}
	}
	return f
}

// Length given a list, return the total number of items within it
func (l IntList) Length() int {
	return len(l)
}

// Map given a function and a list, return the list of the results of applying function(item) on all items
func (l IntList) Map(u unaryFunc) IntList {
	f := make(IntList, len(l))
	for i, v := range l {
		f[i] = u(v)
	}
	return f
}

// Foldl given a function, a list, and initial accumulator, fold (reduce) each item into the
// accumulator from the left using function(accumulator, item)
func (l IntList) Foldl(f binFunc, acc int) int {
	for _, i := range l {
		acc = f(acc, i)
	}
	return acc
}

// Foldr given a function, a list, and an initial accumulator, fold (reduce) each item into the
// accumulator from the right using function(item, accumulator)
func (l IntList) Foldr(f binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = f(l[i], acc)
	}
	return acc
}

// Reverse given a list, return a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {
	f := make(IntList, len(l))
	for i, v := range l {
		f[len(l)-1-i] = v
	}
	return f
}
