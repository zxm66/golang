package main

import (
	"fmt"
)

func main() {
	var str = 1
	if str == 2 {
		say("hello world")
	} else {
		eat()
	}
}
func say(str string) {
	fmt.Println(str)
}

func eat() {
	fmt.Println("this is eat method")

}
