package reflectx

import (
	"reflect"
)

type MethodParameterInfo struct {
	Name          string
	ParameterType reflect.Type
}
