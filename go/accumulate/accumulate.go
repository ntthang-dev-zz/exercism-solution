// Package accumulate iven a collection and an operation to perform on
// each element of the collection, returns a new collection containing
// the result of applying that operation to each element of the input collection.
package accumulate

//Accumulate make operation to perform on each element of the collection, returns a new collection containing
func Accumulate(given []string, converter func(string) string) []string {
	var output []string
	for _, item := range given {
		output = append(output, converter(item))
	}
	return output
}
