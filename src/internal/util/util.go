package util

import (
	"log"
	"reflect"
)

// GetMapValue retrieves a field value from a struct or map by name.
func GetMapValue(data interface{}, name string) interface{} {
	v := reflect.ValueOf(data)

	// Dereference pointer if necessary
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		field := v.FieldByName(name)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	case reflect.Map:
		key := reflect.ValueOf(name)
		val := v.MapIndex(key)
		if !val.IsValid() {
			return nil
		}
		return val.Interface()
	default:
		return nil
	}
}

func LogError(prefix string, err interface{}) {
	log.Fatalf("%v, Error: %v", prefix, err)
}
