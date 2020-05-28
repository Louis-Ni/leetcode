package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string {"abc", "tta"}
	res := groupAnagrams(str)
	fmt.Println(res)
}
func groupAnagrams(strs []string) [][]string {
	fm := make(map[string][]string)
	result := [][]string{}
	for _, word := range(strs){
		charSet := [26]int{0}
		for _, ch :=range(word){
			index := ch - 97
			charSet[index] += 1
		}
		key := strings.Replace(strings.Trim(fmt.Sprint(charSet), "[]"), " ", "", -1)
		fm[key] = append(fm[key], word)
	}

	for _, item := range(fm){
		result = append(result, item)
	}
	return result
}