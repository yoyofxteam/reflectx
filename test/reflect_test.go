package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofxteam/yoyo-reflect"
	"strings"
	"testing"
)

func Test_MethodCallerCall2(t *testing.T) {
	utype := &UserInfo{}

	methodInfo, _ := Reflect.GetObjectMethodInfoByName(utype, "Hello")
	results := methodInfo.Invoke(utype, "yoyogo", "hello world! ")

	fmt.Println()
	fmt.Printf("Result: %s", results)
	fmt.Println()

	assert.Equal(t, results[0].(string), "hello world! yoyogo")
}

func Test_RecCreateStruct(t *testing.T) {
	//yourtype := reflect.TypeOf(Mvc.RequestBody{})
	//dd := Reflect.CreateInstance(yourtype)
	//_ = dd
	typeInfo, _ := Reflect.GetTypeInfo(RequestBody{})
	ins := typeInfo.CreateInstance()
	assert.Equal(t, ins != nil, true)
}

func Test_GetCtorFuncTypeName(t *testing.T) {
	ctorFunc := NewUserController
	name, _ := Reflect.GetCtorFuncOutTypeName(ctorFunc)
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

	ptype, _ := Reflect.GetTypeInfo(p)
	pf1 := ptype.GetFieldByName("Name")
	assert.Equal(t, pf1.GetValue(p), "Json")
	pf2 := ptype.GetFieldByName("Student")
	assert.Equal(t, pf2.GetValue(p), student)

	typeInfo, _ := pf2.AsTypeInfo()
	typeInfo.GetFieldByName("Grade").SetValue(student, 11)
	assert.Equal(t, student.Grade, 11)
	assert.Equal(t, typeInfo.HasMethods(), true)
	sayRet := typeInfo.GetMethodByName("Say").Invoke(student, "World!")[0].(string)
	assert.Equal(t, sayRet, "Hello World!")

}
