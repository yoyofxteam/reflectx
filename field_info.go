package reflectx

import (
	"errors"
	"reflect"
)

// FieldInfo : field info
type FieldInfo struct {
	Name string
	Type reflect.Type
	Kind reflect.Kind
	Tags reflect.StructTag
}

// SetValue : set value to field, field must be kind of reflect.Ptr
func (field FieldInfo) SetValue(instance interface{}, v interface{}) {
	val := reflect.ValueOf(instance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	vField := val.FieldByName(field.Name)
	if vField.CanSet() {
		vField.Set(reflect.ValueOf(v))
	}
}

// GetValue : get value of field
func (field FieldInfo) GetValue(instance interface{}) interface{} {
	val := reflect.ValueOf(instance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	vField := val.FieldByName(field.Name)
	//return field.Value.Interface()
	return vField.Interface()
}

// AsTypeInfo : convert field to TypeInfo
func (field FieldInfo) AsTypeInfo() (TypeInfo, error) {
	if field.Kind == reflect.Struct || field.Kind == reflect.Ptr {
		return GetTypeInfoWithValueType(field.Type)
	}
	return TypeInfo{}, errors.New("must be struct")
}
