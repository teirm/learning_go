// determine if two strings are anagrams
// of each other
package main

import (
    "fmt"
)

// convert a string to a map of characters
func mapString(s string) map[rune]int {
    
    stringMap := make(map[rune]int) 
    
    for _, c := range s {
        stringMap[c]++
    }
    return stringMap
}

// check if two strings are anagrams of
// each other
func areAnagrams(s1, s2 string) bool {
    // if the strings are the same they are not 
    // anagrams of each other
    if s1 == s2 {
        return false
    }

    if len(s1) != len(s2) {
        return false
    }
   
    map1 := mapString(s1)
    map2 := mapString(s2)
    
    for k, _ := range map1 {
        if map1[k] != map2[k] {
            return false
        }
    }

    return true 
}

func main() {
    s1 := "dogs"
    s2 := "gods"
    fmt.Println(areAnagrams(s1, s2))
    
    s1 = "abcd"
    s2 = "aaad"
    fmt.Println(areAnagrams(s1, s2))
}
