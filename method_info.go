package reflectx

import (
	"reflect"
)

// Method Info
type MethodInfo struct {
	Name           string                //Method Name
	MethodInfoType reflect.Type          //method type
	Parameters     []MethodParameterInfo //method's Parameters
	OutType        reflect.Type          //function's return type.
}

// IsValid : method is valid
func (method MethodInfo) IsValid() bool {
	return true
}

// Invoke : invoke the method with interface params.
func (method MethodInfo) Invoke(instance interface{}, params ...interface{}) []interface{} {
	paramsCount := len(params)
	//paramsValues := make([]reflect.Value, paramsCount)
	var paramsValues []reflect.Value
	for idx := 0; idx < paramsCount; idx++ {
		paramsValues = append(paramsValues, reflect.ValueOf(params[idx]))

	}
	return method.InvokeWithValue(reflect.ValueOf(instance), paramsValues...)
}

// InvokeWithValue: invoke the method with value params.
func (method MethodInfo) InvokeWithValue(instance reflect.Value, paramsValues ...reflect.Value) []interface{} {
	methodInfoVal := instance.MethodByName(method.Name)
	returns := methodInfoVal.Call(paramsValues)
	outNum := method.MethodInfoType.NumOut()
	results := make([]interface{}, outNum)
	if len(returns) > 0 {
		for i, res := range returns {
			results[i] = res.Interface()
		}
	}
	return results
}

// AsTypeInfo : convert method to TypeInfo
func (method MethodInfo) AsTypeInfo() (TypeInfo, error) {
	return GetTypeInfoWithValueType(method.MethodInfoType)
}
