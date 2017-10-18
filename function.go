package fuzmatch

import (
	"sort"
	"strings"
	"unicode"
)

//------- function for fuzmatch.go ---------

//test if a string is contains in an array
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

//return the maximum from 3 floats
func maximum(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c {
		return b
	}
	return c
}

//return the smallest string and then the bigger string
func minMax(s1, s2 string) (string, string) {
	if len(s1) <= len(s2) {
		return s1, s2
	}
	return s2, s1
}

//process the string to change to lowercase and remove
//things that are not letters...
func processString(s string) string {
	space := false
	var str []rune
	for _, v := range strings.ToLower(s) {
		if unicode.IsLetter(v) || unicode.IsNumber(v){
			space = false
			str = append(str, v)
		} else if !space {
			space = true
			str = append(str, ' ')
		}
	}

	return strings.TrimSpace(string(str))
}

//algorithm to sort word in t0, t1, t2
//t0 list of strings common in s1 and s2
//t1 list of strings not in t0 from s1
//t2 same as t1 but for s2
func sortListTokenSetRatio(s1, s2 []string) ([]string, []string, []string) {
	t0, t1, t2 := make([]string, 0), make([]string, 0), make([]string, 0)
	//fill the slice t0 and t1
	for i := 0; i < len(s1); i++ {
		found := false
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				t0 = append(t0, s1[i])
				j = len(s2) //finish the inner loop for
				found = true
			}
		}
		if !found {
			t1 = append(t1, s1[i])
		}
	}
	//fill the slice t2
	for _, v := range s2 {
		if !contains(t0, v) {
			t2 = append(t2, v)
		}
	}
	return t0, t1, t2
}

//return a alphabetically sorted string from a list of strings
func sortString(slice []string) string {
	sl := slice
	sort.Strings(sl)
	return strings.Join(sl, " ")
}

// sum two slices to give one slices with slice1 then slice2
func sumSlices(slice1, slice2 []string) []string {
	for _, v := range slice2 {
		slice1 = append(slice1, v)
	}
	return slice1
}
