// Package leap given a year, report if it is a leap year
package leap

// IsLeapYear returns true if it is a leap year otherwise it returns false
// if a year multiple of 400, then it is a leap year
// else if a year multiple of 100, then it is not a leap year
// else if a year multiple of 4, then it is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
