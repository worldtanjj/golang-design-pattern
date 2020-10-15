package main

import (
	"fmt"
)

func main() {
	s := make([]string, 0)
	s = append(s, "aa")
	s = append(s, "bb")
	s = append(s, "cc")

	for i, item := range s {
		fmt.Println(i, item, len(s))
		s = append([]string{fmt.Sprintf("item%v", i)}, s...)
		//s[i] = item + "item"
	}

	fmt.Println(s)
}
