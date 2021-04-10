// Package twofer given a name, return a string with the message:
// One for X, one for me. Where X is the given name.
// However, if the name is missing, return the string: One for you, one for me.
package twofer

// ShareWith tells you how to share with the given name.
func ShareWith(name string) string {
	// If the name is missing, name replaces the value with "you"
	// else return a string with the message have given name.
	if name == "" {
		name = "you"
	}
	// return message
	return "One for " + name + ", one for me."

}
