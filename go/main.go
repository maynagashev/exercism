/*
a-z 97-122
97-122
98-121
99-120
*/
package main

import "fmt"

func main() {
	Atbash("yesyes")
}

func Atbash(s string) (res string) {
	b:= make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shift:= s[i]-'a'
		b[i] ='z'-shift
		fmt.Printf("%#v %d %T %+v %d %d %s\n", s[i], s[i], s[i], i, shift, b, string(b))
	}
	fmt.Println(b, string(b))
	return string(b)
}
