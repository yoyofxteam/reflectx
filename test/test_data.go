package test

import "fmt"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

type RequestBody struct {
}

type UserInfo struct {
	Name string `json:"name" w1:"12"`
	Age  int
}

func (user *UserInfo) Say(hi string) {
	fmt.Print(hi)
}

func (user *UserInfo) Hello(a string, hi string) string {
	return hi + a
}

type Person struct {
	Name    string
	Student *Student
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func (typeInfo Student) Hello() string {
	fmt.Println("hello ")
	return "hello"
}

func (typeInfo Student) Say(hi string) string {
	fmt.Println("Hello " + hi)
	return "Hello " + hi
}
