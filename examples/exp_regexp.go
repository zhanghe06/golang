package main

import (
	"fmt"
	"regexp"
)

func main()  {
	//alphaRegexString          = "^[a-zA-Z]+$"
	//alphaNumericRegexString   = "^[a-zA-Z0-9]+$"
	const (
		codeRegexString = `^[\w\-.@]+$` // \w== [0-9A-Za-z_]
	)

	reg := regexp.MustCompile(codeRegexString)
	fmt.Println(reg.MatchString("2你df"))
	fmt.Println(reg.MatchString("Ww"))
	fmt.Println(reg.MatchString("a"))
	fmt.Println(reg.MatchString("a.2"))
	fmt.Println(reg.MatchString("a--2"))
	fmt.Println(reg.MatchString("a_2"))
}
