package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	//english := []rune("abcdefghijklmnopqrstuvwxyz")
	//swedish := []rune("abcdefghijklmnopqrstuvwxyzåäö")
	ascii := asciiAplhabet()

	text := "ejc\\i d_no\\h"

	// fmt.Println(shift(text, 13, english))
	// fmt.Println(shift(text, -13, english))
	// fmt.Println(shift(text, 13, swedish))
	// fmt.Println(shift(text, -13, swedish))
	fmt.Println(shift(text, -27, ascii))

}

func shift(in string, steps int, alphabet []rune) string {
	alphaLen := float64(len(alphabet))
	if alphaLen < 30 {
		in = strings.ToLower(in)
	}

	out := ""
	for _, r := range in {
		if string(r) == " " {
			out += " "
			continue
		}

		p1 := indexOf(alphabet, r)
		// p2 := p1+steps
		// if(p2 < alphaLenInt){
		// 	out += string(alphabet[int(p2)])
		// 	continue
		// }
		p2 := math.Mod(float64(p1+steps), alphaLen)
		if p2 < 0 {
			p2 = alphaLen + p2
		}
		out += string(alphabet[int(p2)])
	}

	return out
}

func indexOf(haystack []rune, needle rune) int {
	for i, r := range haystack {
		if r == needle {
			return i
		}
	}
	return -1
}

func asciiAplhabet() []rune {
	ret := []rune{}
	for i := 33; i <= 126; i++ {
		ret = append(ret, rune(i))
	}
	return ret
}
