package Reflect

import (
	"reflect"
)

// CreateInstance create new instance by type
func CreateInstance(objectType reflect.Type) interface{} {
	var ins reflect.Value

	ins = reflect.New(objectType)

	if objectType.Kind() == reflect.Struct {
		ins = ins.Elem()
	}

	return ins.Interface()
}

// GetCtorFuncOutTypeName get ctor function return type's name.
func GetCtorFuncOutTypeName(ctorFunc interface{}) (string, reflect.Type) {
	typeInfo, err := GetTypeInfo(ctorFunc)
	if err != nil {
		panic(err.Error())
	}
	return typeInfo.Name, typeInfo.Type
}

// getMethodInfo get method info
func getMethodInfo(method reflect.Method, methodValue reflect.Value) MethodInfo {
	methodInfo := MethodInfo{}
	methodInfo.MethodInfoVal = methodValue
	methodInfo.MethodInfoType = methodValue.Type()
	methodInfo.Name = method.Name
	paramsCount := methodInfo.MethodInfoType.NumIn()
	methodInfo.Parameters = make([]ParameterInfo, paramsCount)

	for idx := 0; idx < paramsCount; idx++ {
		methodInfo.Parameters[idx].ParameterType = methodInfo.MethodInfoType.In(idx)
		methodInfo.Parameters[idx].Name = methodInfo.Parameters[idx].ParameterType.Name()
		if methodInfo.MethodInfoType.NumMethod() > 0 {
			methodInfo.OutType = methodInfo.MethodInfoType.Out(0)
		}
	}

	return methodInfo
}

func GetObjectMethodInfoList(object interface{}) []MethodInfo {
	objectType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	return GetObjectMethodInfoListWithValueType(objectType, objValue)
}

func GetObjectMethodInfoListWithValueType(objectType reflect.Type, objValue reflect.Value) []MethodInfo {
	methodCount := objValue.NumMethod()
	methodInfos := make([]MethodInfo, methodCount)
	for idx := 0; idx < methodCount; idx++ {
		methodInfo := getMethodInfo(objectType.Method(idx), objValue.Method(idx))
		methodInfos[idx] = methodInfo
	}
	return methodInfos
}

func GetObjectMethodInfoByName(object interface{}, methodName string) (MethodInfo, bool) {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	return GetObjectMethodInfoByNameWithType(objType, objValue, methodName)
}

func GetObjectMethodInfoByNameWithType(objectType reflect.Type, objValue reflect.Value, methodName string) (MethodInfo, bool) {
	var methodInfo MethodInfo
	methodType, rbl := objectType.MethodByName(methodName)
	if rbl {
		methodInfo = getMethodInfo(methodType, objValue.MethodByName(methodName))
	}
	return methodInfo, rbl
}
