package others

import (
	"kep/algorithms-go/others"
	"testing"
)

type Anagram struct {
	First    string
	Second   string
	Expected bool
}

func TestIsAnagramWithSuccess(t *testing.T) {
	anagrams := []Anagram{
		Anagram{First: "rat", Second: "tar"},
		Anagram{First: "cider", Second: "cried"},
		Anagram{First: "cider", Second: "criED"},
		Anagram{First: "Dormitory", Second: "Dirty room"}}

	for _, an := range anagrams {
		isAnagram := others.IsAnagram(an.First, an.Second)

		if !isAnagram {
			t.Fatalf("Expected true but got false [%s, %s]", an.First, an.Second)
		}
	}
}

func TestIsAnagramWithError(t *testing.T) {
	anagrams := []Anagram{
		Anagram{First: "dog", Second: "cat"},
		Anagram{First: "dog", Second: "gos"},
		Anagram{First: "big dog", Second: "small dog"}}

	for _, an := range anagrams {
		isAnagram := others.IsAnagram(an.First, an.Second)

		if isAnagram {
			t.Fatalf("Expected false but got true [%s, %s]", an.First, an.Second)
		}
	}
}
