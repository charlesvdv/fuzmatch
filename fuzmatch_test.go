package fuzmatch_test

import (
	"testing"

	"./fuzmatch"
)

var testLevenshtein = []struct {
	n      int
	s1, s2 string
}{
	{2, "book", "back"},
	{4, "booker", "back"},
	{0, "test", "test"},
	{3, "this is a test", "this isn't a test"},
}

var testRa = []struct {
	n      float32
	s1, s2 string
}{
	{0, "", ""},
	{1, "test", "test"},
	{1, "!- ' () test!!", "test"},
	{0, "45 [] # ~", ""},
}

//test the LevenshteinDistance from a set of pre encoded string
func testLevenshteinDistance(t *testing.T) {
	for _, v := range testLevenshtein {
		if res := fuzmatch.LevenshteinDistance(v.s1, v.s2); res != v.n {
			t.Error("LevenshteinDistance Error :", res, "found as distance result, expected :", v.n)
		}
	}
}

//default test for all the ratio just all the equals and the nulls test
func testDefaultAllRatio(t *testing.T) {
	for _, v := range testRa {
		if res := fuzmatch.Ratio(v.s1, v.s2); res != v.n {
			t.Error("Ratio Error : ", res, "found as ratio result, expected :", v.n)
		}
		if res := fuzmatch.PartialRatio(v.s1, v.s2); res != v.n {
			t.Error("PartialRatio Error : ", res, "found as partial ratio result, expected :", v.n)
		}
		if res := fuzmatch.TokenSortRatio(v.s1, v.s2); res != v.n {
			t.Error("TokenSortRatio Error : ", res, "found as token sort ratio result, expected:", v.n)
		}
		if res := fuzmatch.TokenSetRatio(v.s1, v.s2); res != v.n {
			t.Error("TokenSetRatio Error : ", res, "found as token set ratio resutl, expected :", v.n)
		}
	}
}

func testRatio(t *testing.T) {
	// test the symetricity of the ratio function
	s1 := "test"
	s2 := "text"
	if fuzmatch.Ratio(s1, s2) != fuzmatch.Ratio(s2, s1) {
		t.Error("Ratio Error : function not symetric")
	}
}

func testPartialRatio(t *testing.T) {
	//test the basic use of PartialRatio
	s1 := "world"
	s2 := "hello world!"
	if fuzmatch.PartialRatio(s1, s2) != 1 {
		t.Error("PartialRatio Error : should be equals")
	}
}

func testTokenSortRatio(t *testing.T) {
	//test the basic use of TokenSortRatio
	s1 := "golang vs rust"
	s2 := "rust vs golang"
	if fuzmatch.TokenSortRatio(s1, s2) != 1 {
		t.Error("TokenSortRatio : should be equals")
	}
}

func testTokenSetRatio(t *testing.T) {
	//test the basic use of TokenSetRatio
	s1 := "golang from google vs rust from mozilla"
	s2 := "rust vs golang"
	if fuzmatch.TokenSetRatio(s1, s2) != 1 {
		t.Error("TokenSetRatio Error : should be equals")
	}
}
