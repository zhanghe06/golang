package main

import (
	"fmt"
)

type PeopleSpeak interface {
	Speak(string) string
}

type StudentSpeak struct{}

func (stu *StudentSpeak) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	think := "bitch"

	// 声明实例化
	var sp StudentSpeak
	fmt.Println(sp.Speak(think))

	// new 函数实例化
	var stu PeopleSpeak = new(StudentSpeak)
	fmt.Println(stu.Speak(think))

	// 取址实例化
	var peo PeopleSpeak = &StudentSpeak{}
	fmt.Println(peo.Speak(think))
}
