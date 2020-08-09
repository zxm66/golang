package main

import (
	"fmt"
)

func main() {
	say("hello world")
	eat()
}
func say(str string) {
	fmt.Println(str)
}

func eat() {
	fmt.Println("this is eat method")

}
