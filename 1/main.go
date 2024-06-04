package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h Human) GetInfo() {
	fmt.Println(fmt.Sprintf("Имя: %s, Возраст: %d, Пол: %s \n", h.Name, h.Age, h.Gender))
}

type Action struct {
	Human // Встраивание (как аналог наслежования в ооп языках) сокращает код и делает его гибким. Human parent по отношению к Action
}

func main() {
	var a Action
	a.Name = "Ivan"
	a.Age = 20
	a.Gender = "муж"
	
	a.GetInfo()
}
