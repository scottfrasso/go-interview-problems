package main

import (
	"strconv"
	"unicode"
)

// From https://leetcode.com/problems/decode-string/

// Given an encoded string, return its decoded string.
//
// The encoding rule is: k[encoded_string], where the
// encoded_string inside the square brackets is being
// repeated exactly k times. Note that k is guaranteed
// to be a positive integer.
//
// You may assume that the input string is always valid;
// No extra white spaces, square brackets are well-formed, etc.
//
// Furthermore, you may assume that the original data
// does not contain any digits and that digits are only
// for those repeat numbers, k. For example,
// there won't be input like 3a or 2[4].

func encodedString(stringToDecode string) string {
	var counts []int
	var results []string

	currentResult := ""

	values := []rune(stringToDecode)

	i := 0
	for i < len(values) {
		if unicode.IsDigit(values[i]) {
			// get the digits off the start
			count := 0
			for unicode.IsDigit(values[i]) {
				countValue, _ := strconv.Atoi(string(values[i]))
				count = (count * 10) +  countValue
				i++
			}

			counts = append(counts, count)
		} else if values[i] == '[' {
			// pop onto the stack
			results = append(results, currentResult)
			currentResult = ""
			i++
		} else if values[i] == ']' {
			// pop off the stack
			temp := results[len(results)-1]
			results = results[:len(results)-1]

			currentCount := counts[len(counts)-1]
			counts = counts[:len(counts)-1]

			for j := 0; j < currentCount; j++ {
				temp += currentResult
			}
			currentResult = temp
			i++
		} else {
			currentResult += string(values[i])
			i++
		}
	}

	return currentResult
}
