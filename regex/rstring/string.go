package rstring

import (
	"regexp"
)

// CountStringOccurrenceInString Count the number of times a particular string occurs in another string
func CountStringOccurrenceInString(subString string, str string) int {
	re := regexp.MustCompile(subString)
	results := re.FindAllString(str, -1)
	//fmt.Println(len(results))
	return len(results)
}

// split a string into a slice,  Split slices s into substrings separated by the expression and
// returns a slice of the substrings between those expression matches
// strs take 2 arguments,  seperator string and string  which need to be split,
// int n , is max split count
func SplitString(n int, strs ...string) []string {
	str := ""
	sep := ""
	if len(strs) == 1 {
		str = strs[0]
		sep = " "

	} else if len(strs) == 2 {
		str = strs[0]
		sep = strs[1]
	}
	re := regexp.MustCompile(sep)
	split := re.Split(str, n)
	return split

}

// func main() {
// 	countStringOccurrenceInString(2, "a is biy isisisis")
// }
