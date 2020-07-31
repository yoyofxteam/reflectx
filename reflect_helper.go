package reflectx

import (
	"reflect"
)

// Create new instance(interface{}) by type
func CreateInstance(objectType reflect.Type) interface{} {
	var ins reflect.Value

	ins = reflect.New(objectType)

	if objectType.Kind() == reflect.Struct {
		ins = ins.Elem()
	}

	return ins.Interface()
}

// Create new instance(Ptr) by type
func CreateInstancePtr(objectType reflect.Type) interface{} {
	var ins reflect.Value
	ins = reflect.New(objectType)
	return ins
}

// Get ctor function return type's name and reflect.Type.
func GetCtorFuncOutTypeName(ctorFunc interface{}) (string, reflect.Type) {
	typeInfo, err := GetTypeInfo(ctorFunc)
	if err != nil {
		panic(err.Error())
	}
	return typeInfo.Name, typeInfo.Type
}

// Get method info
func getMethodInfo(method reflect.Method) MethodInfo {
	methodInfo := MethodInfo{}
	methodInfo.Name = method.Name
	methodInfo.MethodInfoType = method.Type
	paramsCount := method.Type.NumIn()
	methodInfo.Parameters = make([]MethodParameterInfo, paramsCount)

	for idx := 0; idx < paramsCount; idx++ {
		methodInfo.Parameters[idx].ParameterType = methodInfo.MethodInfoType.In(idx)
		parameterType := methodInfo.Parameters[idx].ParameterType
		if parameterType.Kind() == reflect.Ptr {
			parameterType = parameterType.Elem()
		}
		methodInfo.Parameters[idx].Name = parameterType.Name()
	}
	if methodInfo.MethodInfoType.NumOut() > 0 {
		methodInfo.OutType = methodInfo.MethodInfoType.Out(0)
	}

	return methodInfo
}

// Get instance's method list
func GetObjectMethodInfoList(object interface{}) []MethodInfo {
	objectType := reflect.TypeOf(object)
	return GetObjectMethodInfoListWithValueType(objectType)
}

// Get method list by reflect.Type
func GetObjectMethodInfoListWithValueType(objectType reflect.Type) []MethodInfo {
	methodCount := objectType.NumMethod()
	methodInfos := make([]MethodInfo, methodCount)
	for idx := 0; idx < methodCount; idx++ {
		methodInfo := getMethodInfo(objectType.Method(idx))
		methodInfos[idx] = methodInfo
	}
	return methodInfos
}

// Get Instance's MethodInfo By Method Name and reflect.Type
func GetObjectMethodInfoByName(object interface{}, methodName string) (MethodInfo, bool) {
	objType := reflect.TypeOf(object)
	return GetObjectMethodInfoByNameWithType(objType, methodName)
}

// Get Instance's MethodInfo By Method Name and reflect.Type
func GetObjectMethodInfoByNameWithType(objectType reflect.Type, methodName string) (MethodInfo, bool) {
	var methodInfo MethodInfo
	methodType, rbl := objectType.MethodByName(methodName)
	if rbl {
		methodInfo = getMethodInfo(methodType)
	}
	return methodInfo, rbl
}

// Get final type, means the type is Ptr elem or else.
func GetFinalType(outType reflect.Type) (reflect.Type, bool) {
	var cType reflect.Type
	isPtr := false
	if outType.Kind() != reflect.Ptr {
		cType = outType
	} else {
		isPtr = true
		cType = outType.Elem()
	}
	return cType, isPtr
}

// Get final value, means the type is Ptr elem or else.
func GetFinalValue(outValue reflect.Value) (reflect.Value, bool) {
	var cValue reflect.Value
	isPtr := false
	if outValue.Kind() != reflect.Ptr {
		cValue = outValue
	} else {
		isPtr = true
		cValue = outValue.Elem()
	}
	return cValue, isPtr
}
