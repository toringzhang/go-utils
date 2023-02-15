package basics

import (
	"fmt"
	"reflect"
)

// TravelStruct 遍历复杂struct结构
func TravelStruct(field reflect.Value, prefix string) {
	switch field.Kind() {
	case reflect.String:
		fmt.Printf("%s: %s\n", prefix, field.String())
	case reflect.Int, reflect.Int64:
		fmt.Printf("%s: %d\n", prefix, field.Int())
	case reflect.Struct:
		for i := 0; i < field.NumField(); i++ {
			f := field.Field(i)
			tag := field.Type().Field(i).Name
			if jsr, ok := field.Type().Field(i).Tag.Lookup("json"); ok {
				tag = jsr
			}
			TravelStruct(f, prefix+"."+tag)
		}
	case reflect.Ptr:
		if field.IsNil() {
			fmt.Printf("%s: nil\n", prefix)
			break
		}
		TravelStruct(field.Elem(), prefix)
	default:
		fmt.Printf("name: %s, type: %s\n", prefix, field.Kind().String())
	}
}
