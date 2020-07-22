package Reflect

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
	return hi
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
	return "hello"
}

func (typeInfo Student) Say(hi string) string {
	return "Hello " + hi
}

func Test_MethodCallerCall2(t *testing.T) {
	utype := &UserInfo{}

	methodInfo, _ := GetObjectMethodInfoByName(utype, "Hello")
	results := methodInfo.Invoke("", "hello world!")

	fmt.Println()
	fmt.Printf("Result: %s", results)
	fmt.Println()

	assert.Equal(t, results[0].(string), "hello world!")
}

func Test_RecCreateStruct(t *testing.T) {
	//yourtype := reflect.TypeOf(Mvc.RequestBody{})
	//dd := Reflect.CreateInstance(yourtype)
	//_ = dd
	typeInfo, _ := GetTypeInfo(RequestBody{})
	ins := typeInfo.CreateInstance()
	assert.Equal(t, ins != nil, true)
}

func Test_GetCtorFuncTypeName(t *testing.T) {
	ctorFunc := NewUserController
	name, _ := GetCtorFuncOutTypeName(ctorFunc)
	name = strings.ToLower(name)
	assert.Equal(t, name, "usercontroller")
}

func Test_ReflectStructFields(t *testing.T) {
	student := &Student{
		Name:  "json",
		Age:   18,
		Grade: 9,
	}
	p := Person{
		Name:    "Json",
		Student: student,
	}

	ptype, _ := GetTypeInfo(p)
	pf1 := ptype.GetFieldByName("Name")
	assert.Equal(t, pf1.GetValue(), "Json")
	pf2 := ptype.GetFieldByName("Student")
	assert.Equal(t, pf2.GetValue(), student)
	typeInfo, _ := pf2.AsTypeInfo()

	typeInfo.GetFieldByName("Grade").SetValue(11)
	assert.Equal(t, student.Grade, 11)
	assert.Equal(t, typeInfo.HasMethods(), true)
	sayRet := typeInfo.GetMethodByName("Say").Invoke("World!")[0].(string)
	assert.Equal(t, sayRet, "Hello World!")

}
