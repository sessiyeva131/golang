// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, ШЇЩ†ЩЉШ§!")
	fmt.Println("EX#1")
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	fmt.Println()
	fmt.Println("EX#2")
	fmt.Println(Anagram("listen", "silent"))
	fmt.Println(Anagram("rat", "car"))

	fmt.Println()
	fmt.Println("EX#3")
	list := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(FindEvens(list[:]))

	fmt.Println()
	fmt.Println("EX#4")
	fmt.Println(SliceProduct(list[:]))

	fmt.Println()
	fmt.Println("EX#5")
	var m = make(map[string]int)
	m["T"] = 1
	m["O"] = 2
	m["M"] = 3
	fmt.Println(InvertMap(m))

	fmt.Println()
	fmt.Println("EX#6")
	fmt.Println(TopCharacters("aasssasasww", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	// TODO
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, " ", "")
	// fmt.Println(phone)

	var res string = fmt.Sprintf("(%s) %s-%s", phone[0:3], phone[3:6], phone[6:len(phone)])
	return res
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func Anagram(s1, s2 string) bool {
	// TODO
	var rune1 ByRune = StringToRuneSlice(s1)
	var rune2 ByRune = StringToRuneSlice(s2)

	sort.Sort(rune1)
	sort.Sort(rune2)

	return string(rune1) == string(rune2)
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	// TODO
	var newlist []int
	for i, n := range e {
		if n%2 == 0 {
			newlist = append(newlist, n)
		}
		i = i
	}

	return newlist
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	// TODO
	var sum int = 0
	for i := 0; i < len(e); i++ {
		sum += e[i]
	}
	return sum
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	// TODO
	return nil
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	// TODO
	m := make(map[int]string)

	i := 0
	for k, v := range kv {
		m[v] = k
		i++
	}

	return m
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	// TODO
	m := make(map[rune]int)
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		var count int = strings.Count(s, char)
		if k < count {
			m[chars[i]] = count
		}
	}
	return m
}
