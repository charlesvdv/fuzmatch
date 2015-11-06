package fuzmatch

import (
	"unicode/utf8"
)

//LevenshteinDistance calculate the levenshtein distance between two strings.
//I use an algorithm from Sten Hjelmqvist. http://www.codeproject.com/Articles/13525/Fast-memory-efficient-Levenshtein-algorithm
//You can also see the algorithm on wikipedia : https://en.wikipedia.org/wiki/Levenshtein_distance
//It's the last code.
func LevenshteinDistance(s1, s2 string) int {
	//length of the string s1, s2
	s1Len, s2Len := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2)

	//if the two strings equals
	if s1 == s2 {
		return 0
	}
	//if a string is length 0
	if s1Len == 0 {
		return s2Len
	}
	if s2Len == 0 {
		return s1Len
	}

	v0 := make([]int, s2Len+1)
	v1 := make([]int, s2Len+1)

	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < s1Len; i++ {

		v1[0] = i + 1

		for j := 0; j < s2Len; j++ {
			cost := 1
			if s1[i] == s2[j] {
				cost = 0
			}
			v1[j+1] = minimum(v1[j]+1, v0[j+1]+1, v0[j]+cost)
		}

		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}

	}
	return v1[s2Len]
}

//return the minimum from 3 ints
func minimum(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= c {
		return b
	}
	return c
}
