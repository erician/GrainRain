package main

import "fmt"

func main() {
	var a = '\u0061' //存放的是码值
	fmt.Println(a)

	// var literal []rune
	literal := make([]rune, 0, 10)
	literal = append(literal, 'a')
	fmt.Println(string(literal))
	switch 's' {
	case 's':
		fmt.Println('\v')
		fmt.Println('\f')
	case 'a':
		fmt.Println('a')
	}
	var b rune
	b = '\u0000'
	fmt.Println(string(b))

	var c string
	c = "\001"
	fmt.Println(c)

}
