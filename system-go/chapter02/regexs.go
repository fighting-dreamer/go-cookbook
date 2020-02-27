package chapter02

import (
	"fmt"
	"regexp"
)

func Start1() {
	match, _ := regexp.MatchString("nipun", "NIPUN nipun") // will try to match first arg with 2nd arg [string]
	fmt.Println(match)                                     // true !!!

	match, _ = regexp.MatchString("Nipun", "NIPUN nipun") // will try to match first arg with 2nd arg [string]
	fmt.Println(match)                                    // false !!!

	//parse
	parse, _ := regexp.Compile("Nipun")
	res := parse.MatchString("nqrnqnflqnflqNipunqjqfq") // true case, btw this only gives yes/no (true/false)
	fmt.Println(res)
	//replace using parse

	x := parse.ReplaceAll([]byte("nqrnqnflqnflqNipunqjqfq"), []byte(" NIPUN "))
	fmt.Println(string(x))
}
