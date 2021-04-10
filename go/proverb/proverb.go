// Package proverb given a list of inputs, generate the relevant proverb.
// For want of a horseshoe nail, a kingdom was lost, or so the saying goes.
// 	For example, given the list `["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]`,
//	you will output the full text of this proverbial rhyme:
// 	For want of a nail the shoe was lost.
// 	For want of a shoe the horse was lost.
//	For want of a horse the rider was lost.
//	For want of a rider the message was lost.
//	For want of a message the battle was lost.
//	For want of a battle the kingdom was lost.
//	And all for the want of a nail.
package proverb

import "fmt"

// Proverb generate the relevant proverb for a given list.
func Proverb(rhyme []string) []string {
	// Create proverbial rhymes with rhyme length given
	proverbialRhyme := make([]string, len(rhyme))

	for index := range rhyme {
		if index+1 == len(rhyme) {
			proverbialRhyme[index] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			proverbialRhyme[index] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1])
		}
	}
	return proverbialRhyme
}
