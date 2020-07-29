package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofxteam/reflectx"
	"strings"
	"testing"
)

func Test_MethodCallerCall2(t *testing.T) {
	utype := &UserInfo{}

	methodInfo, _ := reflectx.GetObjectMethodInfoByName(utype, "Hello")
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
	typeInfo, _ := reflectx.GetTypeInfo(RequestBody{})
	ins := typeInfo.CreateInstance()
	assert.Equal(t, ins != nil, true)
}

func Test_GetCtorFuncTypeName(t *testing.T) {
	ctorFunc := NewUserController
	name, _ := reflectx.GetCtorFuncOutTypeName(ctorFunc)
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

	//Get instance 's TypeInfo and get field value
	personTypeInfo, _ := reflectx.GetTypeInfo(p)
	pfName := personTypeInfo.GetFieldByName("Name")
	assert.Equal(t, pfName.GetValue(p), "Json")

	//Get field by name and get it value
	pfStudent := personTypeInfo.GetFieldByName("Student")
	assert.Equal(t, pfStudent.GetValue(p), student)

	// as TypeInfo and SetValue
	pfStudentTypeInfo, _ := pfStudent.AsTypeInfo()
	pfStudentTypeInfo.
		GetFieldByName("Grade").
		SetValue(student, 11)

	assert.Equal(t, student.Grade, 11)
	//---------------------------------------------------------------------------------------------
	assert.Equal(t, pfStudentTypeInfo.HasMethods(), true)
	sayRet := pfStudentTypeInfo.GetMethodByName("Say").Invoke(student, "World!")[0].(string)
	assert.Equal(t, sayRet, "Hello World!")

}

func Test_GetStructMethodList(t *testing.T) {
	userInfo := &UserInfo{}
	userMethodList := reflectx.GetObjectMethodInfoList(userInfo)
	assert.Equal(t, len(userMethodList), 2)
	assert.Equal(t, getMehtodInfoByName(userMethodList, "Hello").Invoke(userInfo, "World!", "UserInfo Func Call:Hello,")[0],
		"UserInfo Func Call:Hello,World!")
	//---------------------------------------------------------------------------------------------
	student := Student{}
	studentMethodList := reflectx.GetObjectMethodInfoList(student)
	assert.Equal(t, len(studentMethodList), 2)
	assert.Equal(t, studentMethodList[0].Name, "Hello")
	assert.Equal(t, studentMethodList[1].Name, "Say")

	assert.Equal(t, getMehtodInfoByName(studentMethodList, "Hello").Invoke(student)[0], "hello")
	assert.Equal(t, getMehtodInfoByName(studentMethodList, "Say").Invoke(student, "Say: Student")[0], "Hello Say: Student")
}

func getMehtodInfoByName(infos []reflectx.MethodInfo, name string) reflectx.MethodInfo {
	for _, m := range infos {
		if m.Name == name {
			return m
		}
	}
	return infos[0]
}
