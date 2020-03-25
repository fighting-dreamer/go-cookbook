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
	res := parse.MatchString("nqrnqnflqnflqNipunqjqfqNipun") // true case, btw this only gives yes/no (true/false)
	fmt.Println(res)
	index := parse.FindIndex([]byte("nqrnqnflqnflqNipunqjqfqNipun"))
	fmt.Println(index)
	indexAll := parse.FindAllIndex([]byte("nqrnqnflqnflqNipunqjqfqNipun"), 2) // only find till next 2 times
	fmt.Println(indexAll)
	indexAll = parse.FindAllIndex([]byte("nqrnqnflqnflqNipunqjqfqNipun"), 3) // suppose to find till next 2 times, what if I ask for 3 times : only 2 times are printed!!!
	fmt.Println(indexAll)
	//replace using parse

	x := parse.ReplaceAll([]byte("nqrnqnflqnflqNipunqjqfq"), []byte(" NIPUN "))
	fmt.Println(string(x))
}
